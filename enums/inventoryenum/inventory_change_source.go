package inventoryenum

import (
	"fmt"
	"github.com/WolffunService/theta-shared-common/enums/rivalitemenum"
)

type ChangeSource int

const (
	CSNonTracking   ChangeSource = 1000
	CSOnBoarding    ChangeSource = 1001
	CSEndBattle     ChangeSource = 1002
	CSRankingReward ChangeSource = 1003
	CSBoxBought     ChangeSource = 1004 //@Tinh comment- khong dung source nay - TODO remove
	CSBoxOpen       ChangeSource = 1005
	CSSkinEvolve    ChangeSource = 1006

	CSSeasonPassUnlockPage ChangeSource = 1007
	CSSeasonPassFree       ChangeSource = 1008
	CSSeasonPassPremium    ChangeSource = 1009

	CSInAppPurchase             ChangeSource = 1010
	CSInAppPurchaseDailyDeal    ChangeSource = 1011
	CSSpecialEventCEMilestone   ChangeSource = 1012
	CSSpecialEventCELeaderboard ChangeSource = 1013
	CSSpecialEventCEPremium     ChangeSource = 1014
	CSSpecialEventPE            ChangeSource = 1015
	CSFreeBox                   ChangeSource = 1016
	CSPremiumBox                ChangeSource = 1017

	CSSSeasonPassEndSeason ChangeSource = 1018

	CSDailyQuestClaimReward ChangeSource = 1020

	CSLobbyClaimMineReward     ChangeSource = 1030
	CSLobbyClaimActivitiReward ChangeSource = 1031

	CSReferralFriend ChangeSource = 1040
	// Admin change source 1900 -> 1999
	CSAdminSend ChangeSource = 1900 // Manual
	CSSystem    ChangeSource = 1901 // Automation
)

var sourceName = map[ChangeSource]string{
	CSOnBoarding:    "On-boarding",
	CSEndBattle:     "Battle end",
	CSRankingReward: "Ranking reward",
	//CSBoxBought:                 "Bought box",
	CSBoxOpen:                   "Open box",
	CSSkinEvolve:                "Evolve skin",
	CSSeasonPassUnlockPage:      "Season Pass Unlock Page",
	CSSeasonPassFree:            "Season Pass Free",
	CSSeasonPassPremium:         "Season Pass Premium",
	CSInAppPurchase:             "In-App purchase",
	CSInAppPurchaseDailyDeal:    "Daily deal",
	CSSpecialEventCEMilestone:   "Special Event CE Milestone",
	CSSpecialEventCELeaderboard: "Special Event CE Leaderboard",
	CSSpecialEventCEPremium:     "Special Event CE buy premium",
	CSSpecialEventPE:            "Special Event PE",
	CSFreeBox:                   "Free Box",
	CSPremiumBox:                "Premium Box",
	CSReferralFriend:            "Referral friend",

	CSAdminSend: "Admin send",
	CSSystem:    "Sytem",
}

func (source ChangeSource) String() string {
	if name, found := sourceName[source]; found {
		return name
	}

	return fmt.Sprintf("Unknown (%d)", source)
}

func (source ChangeSource) IsFree(itemType rivalitemenum.ItemType) bool {
	//special case -- hardcode
	if source == CSSpecialEventCELeaderboard && itemType == rivalitemenum.ITEnhancer {
		return false
	}

	if !valid(itemType) {
		return false
	}

	switch source {
	case CSFreeBox,
		CSInAppPurchaseDailyDeal,
		CSSpecialEventCEMilestone,
		CSSpecialEventCELeaderboard,
		CSSeasonPassFree,
		CSRankingReward,
		CSOnBoarding:
		return true
	default:
		return false
	}
}

// hardcode truoc, tinh sau
func valid(itemType rivalitemenum.ItemType) bool {
	//Minion	Profile		HalfEvolve	FullEvolve	Gold Enhancer

	if itemType == rivalitemenum.ITSkin {
		return true
	}

	if itemType.IsProfile() {
		return true
	}

	if itemType.IsAddIn() {
		return true
	}

	return itemType == rivalitemenum.ITEnhancer || itemType == rivalitemenum.ITGold

}
