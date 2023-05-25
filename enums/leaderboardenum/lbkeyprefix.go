package leaderboardenum

import "fmt"

type LBKeyPrefix string

//Thetan Rivals (NOTE not use `_` in LBKeyPrefix)
const (
	TOURNAMENT        LBKeyPrefix = "TE"
	HERO_FUSION       LBKeyPrefix = "hero_fusion"
	TRADING_EVENT     LBKeyPrefix = "trading_event"
	MARKETPLACE_QUEST LBKeyPrefix = "MARKETPLACE_QUEST"

	TRTrophy        LBKeyPrefix = "TRTrophy"
	TRTrophyCountry LBKeyPrefix = "TRTrophy-%s" //TRTrophy-VN_1

	TRMinion        LBKeyPrefix = "TRMinion"
	TRMinionCountry LBKeyPrefix = "TRMinion-%s" //TRMinion-VN_1

	TRSpecialEventCE LBKeyPrefix = "TRSECE-%s"
)

func (lbKey LBKeyPrefix) GetCountry(country string) LBKeyPrefix {
	return LBKeyPrefix(fmt.Sprintf(string(lbKey), country))
}
