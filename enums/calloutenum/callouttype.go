package calloutenum

import "github.com/WolffunService/thetan-shared-common/enums/mkptabenum"

type CallOutType int

const (
	_Start CallOutType = iota
	FusionLeaderboard
	TournamentLeaderboard
	HeroTradingLeaderboard
	QuestCompetition
	NO_NAME
	SpecialBox
	FusionLive
	TradingEventLive
	TournamentLive
	MKPSingleQuest
	QuestStageComplete
	DailyLoginReward
	SuccessSellItem
	SuccessRentoutItem
	TradingEventChallenge
	_End
)

func (c CallOutType) IsValid() bool {
	return c > _Start && c < _End
}

var _priority = map[CallOutType]int{
	FusionLeaderboard:      11,
	TournamentLeaderboard:  12,
	HeroTradingLeaderboard: 10,
	QuestCompetition:       13,
	NO_NAME:                -1,
	SpecialBox:             9,
	FusionLive:             8,
	TradingEventLive:       7,
	TournamentLive:         6,
	MKPSingleQuest:         5,
	QuestStageComplete:     3,
	DailyLoginReward:       2,
	SuccessSellItem:        15,
	SuccessRentoutItem:     14,
	TradingEventChallenge:  4,
}

func (c CallOutType) Priority() int {
	return _priority[c]
}

var _mess = map[CallOutType]string{
	FusionLeaderboard:      "Leaderboard rank up!",
	TournamentLeaderboard:  "Leaderboard rank up!",
	HeroTradingLeaderboard: "Leaderboard rank up!",
	QuestCompetition:       "Leaderboard rank up!",
	NO_NAME:                "",
	SpecialBox:             "Special box available",
	FusionLive:             "Fusion event started!",
	TradingEventLive:       "Hero trading event started!",
	TournamentLive:         "Thetan Tournament started!",
	MKPSingleQuest:         "Quest completed!",
	QuestStageComplete:     "Quest stage completed!",
	DailyLoginReward:       "Check in now and get {0} quest points!",
	SuccessSellItem:        "Item sold successfully!",
	SuccessRentoutItem:     "Hero rented out successfully!",
	TradingEventChallenge:  "Challenge reward ready!",
}

func (c CallOutType) Message() string {
	return _mess[c]
}

func (c CallOutType) TabEnum() mkptabenum.MKPTabEnum {
	switch c {
	case SpecialBox:
		return mkptabenum.ThetanBox
	case FusionLeaderboard, TournamentLeaderboard, HeroTradingLeaderboard,
		FusionLive, TradingEventLive, TournamentLive, TradingEventChallenge:
		return mkptabenum.SpecialEvent
	case QuestCompetition, MKPSingleQuest, QuestStageComplete, DailyLoginReward:
		return mkptabenum.MKPQuest
	case SuccessSellItem, SuccessRentoutItem:
		return mkptabenum.Profile
	}
	return mkptabenum.Dashboard
}
