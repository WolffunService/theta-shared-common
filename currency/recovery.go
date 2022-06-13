package currency

import (
	"github.com/scylladb/gocqlx/v2"
	"sync"
)

func recoveryWorker(session gocqlx.Session, stats *Stats,
	wg *sync.WaitGroup) {

	defer wg.Done()

	var c = Client{}
	c.Init(session, stats)

loop:
	for {
		transferID, more := <-q.queue
		if !more {
			break loop
		}
		c.RecoverTransfer(transferID)
	}
}

type RecoveryQueue struct {
	queue   chan TransferId
	wg      sync.WaitGroup
	session gocqlx.Session
	stats   *Stats
}

func (q *RecoveryQueue) Init(session gocqlx.Session, stats *Stats) {

	q.session = session
	q.stats = stats
	// Recovery is recursive, create the channels first
	q.queue = make(chan TransferId, 4096000)
}

func (q *RecoveryQueue) StartRecoveryWorker() {
	q.wg.Add(1)
	go recoveryWorker(q.session, q.stats, &q.wg)
}

func (q *RecoveryQueue) Stop() {
	close(q.queue)
	q.wg.Wait()
}

var q RecoveryQueue

func RecoverTransfer(transferId TransferId) {
	q.queue <- transferId
}

func Recover() {
	var c = Client{}
	c.Init(q.session, q.stats)

	c.logger.Info().Op("TransferRecovery").Msg("Fetching dead transfer")

	for {
		iter := q.session.Query(FetchDeadTransfers, []string{}).Iter()
		if iter.NumRows() == 0 {
			break
		}
		// Ignore possible errors
		c.logger.Info().Op("TransferRecovery").
			Var("transfer_count", iter.NumRows()).
			Msg("Found outstanding transfer to recover")
		var transferId TransferId
		for iter.Scan(&transferId) {
			c.RecoverTransfer(transferId)
		}

		if err := iter.Close(); err != nil {
			c.logger.Error().Op("TransferRecovery").Err(err).Msg("Failed to fetch dead transfers")
		}
	}
}

func RecoveryStart(session gocqlx.Session, stat *Stats) {
	q.Init(session, stat)

	// Start background fiber working on the queue to
	// make sure we purge it even during the initial recovery
	for i := 0; i < 8; i++ {
		q.StartRecoveryWorker()
	}

	Recover()
}

func RecoveryStop() {
	Recover()
	q.Stop()
}
