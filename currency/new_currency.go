package currency

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-common/common/thetaerror"
	"github.com/WolffunGame/theta-shared-common/thetalog"
	"github.com/ansel1/merry/v2"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gopkg.in/inf.v0"
	"math/rand"
	"sync/atomic"
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
	Found           bool
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

	Participation []UserBalance
	LockOrder     []*UserBalance
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
	c.rawSession.SetConsistency(gocql.Quorum)
	c.stats = stats
	c.insertTransfer = session.Session.Query(InsertTransfer)

	c.setTransferClient = session.Session.Query(SetTransferClient)
	c.setTransferState = session.Session.Query(SET_TRANSFER_STATE)
	c.clearTransferClient = session.Session.Query(ClearTransferClient)
	c.deleteTransfer = session.Session.Query(DeleteTransfer)
	c.fetchTransfer = session.Session.Query(FetchTransfer)
	c.fetchTransfer.SerialConsistency(gocql.Serial)
	c.fetchTransferClient = session.Session.Query(FetchTransferClient)
	c.fetchTransferClient.SerialConsistency(gocql.Serial)
	c.lockAccount = session.Session.Query(LockAccount)
	c.unlockAccount = session.Session.Query(UnlockAccount)
	c.updateAccountBalance = session.Session.Query(UpdateBalance)
	c.fetchAccountBalance = session.Session.Query(FetchBalance)

	func(queries ...*gocql.Query) {
		for _, query := range queries {
			query.SetConsistency(gocql.Quorum)
		}
	}(c.setTransferClient, c.setTransferState, c.clearTransferClient,
		c.deleteTransfer, c.lockAccount, c.unlockAccount, c.updateAccountBalance)

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

		return merry.Wrap(err)
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
			return merry.Wrap(gocql.ErrNotFound)
		}

		if c.clientId != prev.ClientId {
			// The transfer is already worked on.
			return merry.New(fmt.Sprintf("our id %v, previous id %v",
				c.clientId, prev.ClientId), merry.WithCause(gocql.ErrNotFound))
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
			return merry.Wrap(err)
		}

		rowClientId, exists := row["client_id"]
		if !exists || rowClientId == nilUuid {
			c.logger.Trace().Op("SetTransferState").Var("client_id", c.clientId).Msg("Failed to set state: no such transfer")
			return merry.Wrap(gocql.ErrNotFound)
		}

		return merry.Wrap(gocql.ErrNotFound, merry.WithCause(merry.New(fmt.Sprintf("our id %v, previous id %v",
			c.clientId, rowClientId))))
	}
	t.State = state
	return nil
}

func (c *Client) ClearTransferClient(transferId TransferId) {
	c.logger.Info().Op("ClearTransferClient").Var("transfer_id", transferId).Send()

	cql := c.clearTransferClient
	cql.Bind(transferId, c.clientId)
	row := Row{}
	if applied, err := cql.MapScanCAS(row); err != nil || !applied {
		if err != nil {
			c.logger.Error().Err(err).Op("ClearTransferClient").Var("transfer_id", transferId).Msg("Failed to clear transfer client")
		} else if !applied {
			rowClientId, exists := row["client_id"]
			if !exists || rowClientId == nilUuid {
				// The transfer is gone, do not complain
			} else {
				c.logger.Error().Err(merry.New(fmt.Sprintf("Client id mismatch: %v != %v",
					c.clientId, row["client_id"]))).Op("ClearTransferClient").
					Var("transfer_id", transferId).
					Var("client_id", c.clientId).
					Var("transfer_client_id", rowClientId).
					Msg("Failed to clear transfer client")
			}
		}
	}
}

func (c *Client) FetchAccountBalance(bal *UserBalance) error {
	cql := c.fetchAccountBalance
	cql.Bind(bal.UserId, bal.CurrencyType)
	if err := cql.Scan(&bal.Balance, &bal.PendingAmount); err != nil {
		return err
	}
	bal.Found = true
	return nil
}

func (c *Client) UnlockAccount(transferId TransferId, bal *UserBalance) error {
	return c.unlockAccount.Bind(bal.UserId, bal.CurrencyType, transferId).Exec()
}

func (c *Client) LockAccounts(t *Transfer, wait bool) error {
	if t.State == TransferStateCompleted {
		return nil
	}
	if t.State == TransferStateLocked {
		// Transfer đã được lock, fetch balance kiểm tra các account tham gia vào quá trình giao dịch có tồn tại không
		for i := 0; i < 2; i++ {
			if t.Participation[i].UserId == "" {
				//system credit
				continue
			}

			if err := c.FetchAccountBalance(&t.Participation[i]); err != nil && err != gocql.ErrNotFound {
				return merry.Wrap(err)
			}
		}

		c.logger.Trace().Op("LockAccounts").Var("transfer", t).Var("transfer_id", t.TransferId).Msg("Fetched locked")
		return nil
	}

	c.logger.Trace().Op("LockAccounts").Var("transfer", t).Var("transfer_id", t.TransferId).Msg("Locking")
	sleepDuration := time.Millisecond*time.Duration(rand.Intn(10)) + time.Millisecond
	maxSleepDuration, _ := time.ParseDuration("10s")

	// Upon failure to take lock on the second account, we should try to rollback
	// lock on the first to avoid deadlocks. We shouldn't, however, accidentally
	// rollback the lock if we haven't taken it - in this case lock0
	// and lock1 both may have been taken, and the transfer have progressed
	// to moving the funds, so rolling back the lock would break isolation.
	var previousAccount *UserBalance

	var i = 0
	for i < 2 {
		account := t.LockOrder[i]
		if account.UserId == "" {
			//system credit
			continue
		}
		cql := c.lockAccount
		cql.Bind(t.TransferId, account.PendingAmount, account.UserId, account.CurrencyType)
		row := Row{}
		// If the update is not applied because we've already locked the
		// transfer, it's a success. This is possible during recovery.
		lockFailed := func(applied bool) bool {
			if applied {
				return false
			}
			// pendingTransfer may be missing from returns (Cassandra)
			pendingTransfer, exists := row["pending_transfer"].(TransferId)
			if exists && pendingTransfer == t.TransferId {
				return false
			}
			return true
		}
		if applied, err := cql.MapScanCAS(row); err != nil || lockFailed(applied) {
			// Remove the pending transfer from the previously
			// locked account, do not wait with locks.
			if i == 1 && previousAccount != nil {
				if err1 := c.UnlockAccount(t.TransferId, previousAccount); err1 != nil {
					return merry.Wrap(err1, merry.WithCause(err))
				}
			}
			// Check for transient errors, such as query timeout, and retry.
			// In case of a non-transient error, return it to the client.
			// No money changed its hands and the transfer can be recovered
			// later
			if err != nil {
				if IsTransientError(err) {
					c.logger.Trace().Op("LockAccounts").Err(err).Var("transfer", t).Var("transfer_id", t.TransferId).Msg("Retrying after error")
				} else {
					return merry.Wrap(err)
				}
			} else {
				// Lock failed because of a conflict or account is missing.
				pendingTransfer, exists := row["pending_transfer"].(TransferId)
				if !exists || pendingTransfer == nilUuid {
					// No such account. We're not holding locks. CompleteTransfer() will delete
					// the trans
					return c.SetTransferState(t, TransferStateLocked)
				}
				// There is a non-empty pending transfer. Check if the
				// transfer we've conflicted with is orphaned and recover
				// it, before waiting

				var clientId gocql.UUID
				c.fetchTransferClient.Bind(pendingTransfer)
				if err := c.fetchTransferClient.Scan(&clientId); err != nil {
					if err != gocql.ErrNotFound {
						return err
					}

					// Transfer không tồn tại, mặc dù ngay trước đó đã aborted lock
					// Trường hợp này có thể xảy ra khi transfer vừa complete
					c.logger.Trace().Op("LockAccounts").
						Err(err).
						Var("pending_transfer_id", pendingTransfer).
						Var("transfer_id", t.TransferId).Msg("Transfer which aborted our lock is now gone")
				} else if clientId == nilUuid {
					// Transfer pending trong user account không được xử lý bởi client nào, recover nó
					c.logger.Trace().Op("LockAccounts").
						Err(err).
						Var("pending_transfer_id", pendingTransfer).
						Var("transfer_id", t.TransferId).Msg("Add transfer to the recovery queue")

					RecoverTransfer(pendingTransfer)
				}
				atomic.AddUint64(&c.stats.retries, 1)

				if !wait {
					return merry.New("Wait aborted")
				}
			}
			// Restart locking
			i = 0

			time.Sleep(sleepDuration)

			llog.Tracef("[%v] [%v] Restarting after sleeping %v",
				c.shortId, t.id, sleepDuration)

			sleepDuration = sleepDuration * 2
			if sleepDuration > maxSleepDuration {
				sleepDuration = maxSleepDuration
			}
			t.acs[0].found = false
			t.acs[1].found = false
			previousAccount = nil
			// Reset client id in case it expired while we were sleeping
			if err := c.SetTransferClient(t.id); err != nil {
				return err
			}
		} else {
			if applied {
				previousAccount = account
			}
			// In Scylla, the previous row returned even if LWT is applied.
			// In Cassandra, make a separate query.
			if account.balance, account.found = row["balance"].(*inf.Dec); !account.found {
				// Support Cassandra which doens't provide balance
				if err = c.FetchAccountBalance(account); err != nil {
					return merry.Wrap(err)
				}
			} else if !applied {
				// Fetch previous pending amount
				account.pending_amount = row["pending_amount"].(*inf.Dec)
			}
			i++
		}
	}
	// Move transfer to 'locked', to not attempt to transfer
	// 	the money twice during recovery
	return c.SetTransferState(t, "locked")
}

func (c *Client) RecoverTransfer(transferId TransferId) {
	cookie := StatsRequestStart()

	c.logger.Trace().Var("transfer_id", transferId).Msg("Recovering transfer")

	atomic.AddUint64(&c.stats.recoveries, 1)
	if err := c.SetTransferClient(transferId); err != nil {
		if !merry.Cause(err, gocql.ErrNotFound) {
			llog.Errorf("[%v] [%v] Failed to set client on transfer: %v",
				c.shortId, transferId, err)
		}
		return
	}
	t := new(Transfer)
	t.InitEmptyTransfer(transferId)
	cql := c.fetchTransfer
	cql.Bind(transferId)
	// Ignore possible error, we will retry
	if err := cql.Scan(&t.acs[0].bic, &t.acs[0].ban, &t.acs[1].bic,
		&t.acs[1].ban, &t.amount, &t.state); err != nil {

		if err == gocql.ErrNotFound {
			llog.Errorf("[%v] [%v] Transfer not found when fetching for recovery",
				c.shortId, transferId)
		} else {
			llog.Errorf("[%v] [%v] Failed to fetch transfer: %v",
				c.shortId, transferId, err)
		}
		return
	}
	t.InitAccounts()
	if err := c.LockAccounts(t, false); err != nil {
		llog.Errorf("[%v] [%v] Failed to lock accounts: %v",
			c.shortId, t.id, err)
		c.ClearTransferClient(t.id)
		return
	}
	if err := c.CompleteTransfer(t); err != nil {
		llog.Errorf("[%v] [%v] Failed to complete transfer: %v",
			c.shortId, t.id, err)
	} else {
		StatsRequestEnd(cookie)
	}
}

func IsTransientError(err error) bool {
	err = merry.Cause(err)
	reqErr, isRequestErr := err.(gocql.RequestError)
	if isRequestErr && reqErr != nil {
		return true
	} else if err == gocql.ErrTimeoutNoResponse {
		return true
	} else {
		return false
	}
}
