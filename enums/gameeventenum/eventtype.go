package gameeventenum

import "strings"

type EventType int32

const (
	NONE EventType = iota // for none - ranked
	TOURNAMENT_EVENT
	end_enum
)

func (e EventType) IsValid() bool {
	return e >= 0 && e < end_enum
}

type EventTypeString string

const (
	TOURNAMENT_EVENT_STRING       EventTypeString = "TE_{0}"
	DAILY_TOURNAMENT_EVENT_STRING EventTypeString = "TE_{0}_{1}"
)

func (t EventTypeString) EventName() string {
	return strings.Split(string(t), "_")[0]
}

func IsEventType(leaderboardName string) bool {
	split := strings.Split(leaderboardName, "_")
	switch split[0] {
	case TOURNAMENT_EVENT_STRING.EventName():
		return true

	}
	return false
}
