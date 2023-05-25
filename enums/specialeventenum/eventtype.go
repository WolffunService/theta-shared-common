package specialeventenum

import "strings"

type EventType int32

const (
	NONE EventType = iota // for none - ranked
	TOURNAMENT_EVENT
	LIST_TOURNAMENT_EVENT
	HERO_FUSION_EVENT
	COINBASE_LOOT_EVENT
	TRADING_EVENT
	HERO_POOL_MKP_QUEST_EVENT
	GIFTMAS_WHEEL_EVENT
	SAMSUNG_EVENT
	end_enum
)

func (e EventType) IsValid() bool {
	return e > NONE && e < end_enum
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
