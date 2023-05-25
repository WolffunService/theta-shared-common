package thetancontext

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type (
	Track interface {
		TimeTrack(start time.Time, name ...string)
	}

	Tracker struct {
		//more
		enable bool
		logger *zerolog.Logger
	}
)

func NewTracker(enable bool) Tracker {
	return Tracker{enable: enable}
}

func (tr Tracker) TimeTrack(start time.Time, name ...string) {
	if !tr.enable {
		return
	}

	if tr.logger == nil {
		tr.logger = &log.Logger
	}

	elapsed := time.Since(start)
	_, _, fnName := trace()
	tr.logger.Trace().Strs("name", name).Str("fnName", fnName).
		Msgf("took %s ", elapsed)
}
