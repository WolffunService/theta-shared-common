package currency

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-common/thetalog"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"time"
)

type TransferId = gocql.UUID
type TransferState int32

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
	UserID          string
	CurrencyType    int
	Balance         int64
	PendingTransfer gocql.UUID
	PendingAmount   int64
	UpdatedAt       time.Time
}

type Transfer struct {
	TransferID   gocql.UUID
	CurrencyType int
	SourceID     string
	DestID       string
	Amount       int64
	State        TransferState
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
	payStats             *Stats
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

func (c *Client) RegisterTransfer(t *Transfer) error {
	c.logger.Info().Op("RegisterTransfer").Var("client_id", c.clientId).Send()

	qb.Insert(transferMetadata.Name).Unique().Query(c.session).Consistency(gocql.Quorum).BindStruct(t)

	// Register a new transfer
	cql := c.insertTransfer
	cql.Bind(t.id, t.acs[0].bic, t.acs[0].ban, t.acs[1].bic, t.acs[1].ban, t.amount)
	row := Row{}
	if applied, err := cql.MapScanCAS(row); err != nil || !applied {
		if err == nil && !applied {
			// Should never happen, transfer id is globally unique
			llog.Fatalf("[%v] [%v] Failed to create: a duplicate transfer exists",
				c.shortId, t.id)
		}
		return merry.Wrap(err)
	}
	return c.SetTransferClient(t.id)
}

func (c *Client) SetTransferClient(transferId TransferId) error {
	c.logger.Info().Op("RegisterTransfer").Var("client_id", c.clientId).Var("transfer_id", transferId).Send()

	cql := c.setTransferClient
	cql.Bind(c.clientId, transferId)
	// Change transfer - set client id

	qb.Update(transferMetadata.Name).
		Set("client_id").
		Where(qb.Eq("transfer_id")).
		If(qb.Ne("amount"))

	row := Row{}
	if applied, err := cql.MapScanCAS(row); err != nil || !applied {
		if err != nil {
			return merry.Wrap(err)
		}
		rowClientId, exists := row["client_id"]
		if !exists || rowClientId == nilUuid {
			llog.Tracef("[%v] [%v] Failed to set client: no such transfer",
				c.shortId, transferId)
			return merry.Wrap(gocql.ErrNotFound)
		}
		if c.clientId != rowClientId {
			// The transfer is already worked on.
			err := merry.New(fmt.Sprintf("our id %v, previous id %v",
				c.clientId, rowClientId))
			return merry.WithCause(gocql.ErrNotFound, err)
		} // c.clientId == rowClientId
	}
	return nil
}
