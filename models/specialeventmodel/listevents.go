package specialeventmodel

import (
	"time"

	"github.com/WolffunService/theta-shared-common/enums/specialeventenum"
)

type ListEventsProfile []EventProfile

type EventProfile struct {
	EventType specialeventenum.EventType `json:"eventType"`
	StartsIn  int64                      `json:"startsIn"`
	Duration  int64                      `json:"duration"`
	Priority  int                        `json:"priority"`
	WarmUpIn  int64                      `json:"warmUpIn"`
}

func (e EventProfile) EventStatus() specialeventenum.EventStatus {
	now := time.Now().UTC().Unix()
	if e.IsEnded(now) {
		return specialeventenum.ENDED
	}
	if e.IsWarmUpTime(now) {
		return specialeventenum.WARM_UP
	}
	return specialeventenum.LIVE
}

func (e EventProfile) IsEnded(now int64) bool {
	return (e.StartsIn + e.Duration) <= now
}

func (e EventProfile) IsWarmUpTime(now int64) bool {
	return e.WarmUpIn <= now && now < e.StartsIn
}

func (e EventProfile) IsShowEvent(now int64) bool {
	if e.WarmUpIn > 0 {
		return e.WarmUpIn < now
	}
	return true
}

type EventProfileClient struct {
	EventType          specialeventenum.EventType `json:"eventType"`
	StartsIn           int64                      `json:"startsIn"`
	EndsIn             int64                      `json:"endsIn"`
	EventName          string                     `json:"eventName"`
	EventDescription   string                     `json:"eventDescription"`
	EventURLBackground string                     `json:"eventURLBackground"`
}

func ToEventProfileClient(events ListEventsProfile, fn func(events ListEventsProfile) []EventProfileClient) []EventProfileClient {
	return fn(events)
}
