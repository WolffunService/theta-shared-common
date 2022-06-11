package currency

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-common/common/thetaerror"
	"github.com/WolffunGame/theta-shared-common/thetalog"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"time"
)

type TransferId = gocql.UUID
type TransferState int32
type Row = map[string]interface{}

const (
	TransferStateNew       TransferState = 2
	TransferStateLocked    TransferState = 1
	TransferStateCompleted TransferState = 0
)

var nilUuid gocql.UUID

func NewTransferId() TransferId {
	return gocql.TimeUUID()
}

type UserBalance struct {
	UserId          string
	CurrencyType    int
	Balance         int64
	PendingTransfer gocql.UUID
	PendingAmount   int64
	UpdatedAt       time.Time
}

type Transfer struct {
	TransferId   TransferId
	CurrencyType int
	SourceID     string
	DestID       string
	Amount       int64
	State        TransferState
	ClientId     gocql.UUID
	CreatedAt    time.Time
}

//Stats Storing stat của hệ thống currency mới nhằm mục đích monitoring
type Stats struct {
	errors            uint64
	noSuchAccount     uint64
	insufficientFunds uint64
	retries           uint64
	recoveries        uint64
}

type Client struct {
	clientId             gocql.UUID // For locking
	session              gocqlx.Session
	rawSession           *gocql.Session
	stats                *Stats
	logger               thetalog.Logger
	insertTransfer       *gocql.Query
	setTransferClient    *gocql.Query
	clearTransferClient  *gocql.Query
	setTransferState     *gocql.Query
	deleteTransfer       *gocql.Query
	fetchTransfer        *gocql.Query
	fetchTransferClient  *gocql.Query
	lockAccount          *gocql.Query
	unlockAccount        *gocql.Query
	fetchAccountBalance  *gocql.Query
	updateAccountBalance *gocql.Query
}

func (c *Client) Init(session gocqlx.Session, stats *Stats) {
	c.clientId = gocql.TimeUUID()
	c.logger.Info().Op("CurrencyClientInit").Var("client_id", c.clientId).Send()

	c.session = session
	c.rawSession = session.Session
	c.stats = stats
	c.insertTransfer = session.Session.Query(INSERT_TRANSFER)
	c.setTransferClient = session.Session.Query(SET_TRANSFER_CLIENT)
	c.setTransferState = session.Session.Query(SET_TRANSFER_STATE)
	c.clearTransferClient = session.Session.Query(CLEAR_TRANSFER_CLIENT)
	c.deleteTransfer = session.Session.Query(DELETE_TRANSFER)
	c.fetchTransfer = session.Session.Query(FETCH_TRANSFER)
	c.fetchTransfer.SerialConsistency(gocql.Serial)
	c.fetchTransferClient = session.Session.Query(FETCH_TRANSFER_CLIENT)
	c.fetchTransferClient.SerialConsistency(gocql.Serial)
	c.lockAccount = session.Session.Query(LOCK_ACCOUNT)
	c.unlockAccount = session.Session.Query(UNLOCK_ACCOUNT)
	c.updateAccountBalance = session.Session.Query(UPDATE_BALANCE)
	c.fetchAccountBalance = session.Session.Query(FETCH_BALANCE)
	c.fetchAccountBalance.SerialConsistency(gocql.Serial)
}

func (c *Client) RegisterTransfer(t *Transfer) error {
	c.logger.Info().Op("RegisterTransfer").Var("client_id", c.clientId).Send()

	q := qb.Insert(transferMetadata.Name).Unique().Query(c.session).Consistency(gocql.Quorum).BindStruct(t)
	var prev Transfer

	if applied, err := q.GetCASRelease(&prev); err != nil || !applied {
		if err == nil && !applied {
			// Should never happen, transfer id is globally unique
			c.logger.Fatal().Op("RegisterTransfer").Var("client_id", c.clientId).Var("transfer_id", t.TransferId).Send()
		}

		return &thetaerror.Error{
			Code:    thetaerror.ErrorInternal,
			Message: "RegisterTransfer Error",
			Op:      "RegisterTransfer",
			Err:     err,
		}
	}
	return c.SetTransferClient(t.TransferId)
}

func (c *Client) SetTransferClient(transferId TransferId) error {
	c.logger.Info().Op("SetTransferClient").Var("client_id", c.clientId).Var("transfer_id", transferId).Send()

	type TransferQuery struct {
		TransferID        TransferId
		Amount            interface{}
		ClientID          gocql.UUID
		ConditionClientID interface{}
	}

	q := qb.Update(transferMetadata.Name).
		Set("client_id").
		Where(qb.Eq("transfer_id")).
		If(qb.Ne("amount")).
		If(qb.EqNamed("client_id", "condition_client_id")).Query(c.session).Consistency(gocql.Quorum).BindStruct(&TransferQuery{
		TransferID:        transferId,
		Amount:            nil,
		ClientID:          c.clientId,
		ConditionClientID: nil,
	})
	var prev Transfer

	if applied, err := q.GetCASRelease(prev); err != nil || !applied {
		if err != nil {
			return &thetaerror.Error{
				Code: thetaerror.ErrorInternal,
				Err:  err,
			}
		}

		if prev.ClientId == nilUuid {
			c.logger.Trace().Op("SetTransferClient").Var("transfer_id", transferId).Msg("Failed to set client: no such transfer")
			return &thetaerror.Error{
				Code: thetaerror.ErrorInternal,
				Err:  gocql.ErrNotFound,
			}
		}

		if c.clientId != prev.ClientId {
			// The transfer is already worked on.
			return &thetaerror.Error{
				Message: fmt.Sprintf("our id %v, previous id %v",
					c.clientId, prev.ClientId),
				Code: thetaerror.ErrorInternal,
				Err:  gocql.ErrNotFound,
			}
		}

		// c.clientId == rowClientId
	}

	return nil
}

func (c *Client) SetTransferState(t *Transfer, state TransferState) error {
	c.logger.Info().Op("SetTransferState").Var("transfer_id", t.TransferId).Send()

	cql := c.setTransferState
	cql.Bind(state, t.TransferId, c.clientId)
	row := Row{}
	if applied, err := cql.MapScanCAS(row); err != nil || !applied {
		if err != nil {
			return &thetaerror.Error{
				Code: thetaerror.ErrorInternal,
				Err:  err,
			}
		}

		rowClientId, exists := row["client_id"]
		if !exists || rowClientId == nilUuid {
			c.logger.Trace().Op("SetTransferState").Var("client_id", c.clientId).Msg("Failed to set state: no such transfer")
			return &thetaerror.Error{
				Code: thetaerror.ErrorInternal,
				Err:  gocql.ErrNotFound,
			}
		}
		return &thetaerror.Error{
			Message: fmt.Sprintf("our id %v, previous id %v",
				c.clientId, rowClientId),
			Code: thetaerror.ErrorInternal,
			Err:  gocql.ErrNotFound,
		}
	}
	t.State = state
	return nil
}
