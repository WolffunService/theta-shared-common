package currency

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-common/thetalog"
	"github.com/spenczar/tdigest"
	"time"
)

type Metrics struct {
	nRequests  int64
	cpuTime    time.Duration
	latencyMin time.Duration
	latencyMax time.Duration
	latencyAvg time.Duration
	tdigest    *tdigest.TDigest
}

func (m *Metrics) Reset() {
	m.nRequests = 0
	m.cpuTime = 0
	m.latencyMax = 0
	m.latencyMin = 0
	m.latencyAvg = 0
	m.tdigest = tdigest.New()
}

func (m *Metrics) Update(elapsed time.Duration) {
	m.nRequests++
	m.cpuTime += elapsed
	if elapsed > m.latencyMax {
		m.latencyMax = elapsed
	}
	if m.latencyMin == 0 || m.latencyMin > elapsed {
		m.latencyMin = elapsed
	}
	m.tdigest.Add(elapsed.Seconds(), 1)
}

type stats struct {
	n_total   int64
	starttime time.Time
	periodic  Metrics
	summary   Metrics
	queue     chan time.Duration
	done      chan bool
}

var s stats

type cookie struct {
	time time.Time
}

func statsWorker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	more := true
loop:
	for {
		var elapsed time.Duration
		select {
		case <-ticker.C:
			if s.summary.nRequests > 0 {
				var progress string
				if s.n_total > 0 {
					progress = fmt.Sprintf("%5s%% done, RPS %d",
						fmt.Sprintf("%.2f", float64(s.summary.nRequests)/float64(s.n_total)*100),
						s.periodic.nRequests)
				} else {
					progress = fmt.Sprintf("Done %10d requests", s.summary.nRequests)
				}

				thetalog.Info().Msgf("%s, Latency min/max/med: %.3fs/%.3fs/%.3fs",
					progress,
					s.periodic.latencyMin.Seconds(),
					s.periodic.latencyMax.Seconds(),
					s.periodic.tdigest.Quantile(0.5),
				)

				s.periodic.Reset()
			}
		case elapsed, more = <-s.queue:
			if !more {
				break loop
			}
			s.periodic.Update(elapsed)
			s.summary.Update(elapsed)
		}
	}
	s.done <- true
}

func StatsSetTotal(n int) {
	s.n_total = int64(n)
}

func StatsInit() {
	s.starttime = time.Now()
	s.periodic.Reset()
	s.summary.Reset()
	s.queue = make(chan time.Duration, 1000)
	s.done = make(chan bool, 1)
	go statsWorker()
}

func StatsRequestStart() cookie {
	return cookie{
		time: time.Now(),
	}
}

func StatsRequestEnd(c cookie) {
	s.queue <- time.Since(c.time)
}

func StatsReportSummary() {
	// Stop background work
	close(s.queue)
	<-s.done

	if s.summary.nRequests == 0 {
		return
	}

	wallclocktime := time.Since(s.starttime).Seconds()

	thetalog.Info().Msgf("Total time: %.3fs, %v t/sec",
		wallclocktime,
		int(float64(s.summary.nRequests)/wallclocktime),
	)

	thetalog.Info().Msgf("Latency min/max/avg: %.3fs/%.3fs/%.3fs",
		s.summary.latencyMin.Seconds(),
		s.summary.latencyMax.Seconds(),
		s.summary.cpuTime.Seconds()/float64(s.summary.nRequests),
	)

	thetalog.Info().Msgf("Latency 95/99/99.9%%: %.3fs/%.3fs/%.3fs",
		s.summary.tdigest.Quantile(0.95),
		s.summary.tdigest.Quantile(0.99),
		s.summary.tdigest.Quantile(0.999),
	)
}
