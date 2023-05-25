package specialeventmodel

import (
	"fmt"

	"github.com/WolffunService/theta-shared-common/database/mongodb"
	"github.com/WolffunService/theta-shared-common/enums/gamemodeenum"
)

func (UserEventData) CollectionName() string {
	return "UserSpecialEventData"
}

type UserEventData struct {
	mongodb.DefaultModel `bson:",inline"`
	mongodb.DateFields   `bson:",inline"`

	TournamentEventsData UserTournamentEventsData `json:"tournamentEventsData" bson:"tournamentEventsData"`
}

type UserTournamentEventsData struct {
	EventId         int              `json:"eventId" bson:"eventId"`
	DailyBattleUser *UserDailyBattle `json:"dailyBattleUser" bson:"dailyBattleUser"`
	TrophylEvent    map[int]int      `json:"trophylEvent" bson:"trophylEvent"`
}

type UserDailyBattle struct {
	DailyEventId     int                       `json:"dailyEventId" bson:"dailyEventId"`
	DailyBattle      DailyBattleConfig         `json:"dailyBattleUser,inline" bson:"dailyBattleUser"`
	DailyTrophyEvent map[int]*DailyTrophyModel `json:"dailyTrophyEvent" bson:"dailyTrophyEvent"`
}

type DailyTrophyModel struct {
	EventId     int `json:"eventId" bson:"eventId"`
	TrophyEvent int `json:"trophyEvent" bson:"trophyEvent"`
}

func (t *UserTournamentEventsData) ResetAndJoinNewDailyEvent(dailyEventId int) {
	//t.DailyBattleUser = &UserDailyBattle{}
	t.DailyBattleUser.DailyBattle = DailyBattleConfig{} //reset daily mode battle
	t.DailyBattleUser.DailyEventId = dailyEventId       //join new event
	//t.DailyBattleUser.DailyBattle = DailyBattleConfig{}
}

func (u *UserDailyBattle) IncreaseBattleByMode(mode gamemodeenum.InGameMode) {
	switch mode {
	case gamemodeenum.TEAM_COLLECT_STAR_4_VS_4:
		u.DailyBattle.SSDailyBattleConfig++
	case gamemodeenum.DEATH_MATCH:
		u.DailyBattle.DMDailyBattleConfig++
	case gamemodeenum.TOWER:
		u.DailyBattle.TWDailyBattleConfig++
	case gamemodeenum.SOLO_SURVIVAL:
		u.DailyBattle.SoloDailyBattleConfig++
	case gamemodeenum.DUAL_SURVIVAL:
		u.DailyBattle.DualDailyBattleConfig++
	}
}

func (t *UserTournamentEventsData) IncreaseTrophyEvent(trophyEvent int) int {
	if len(t.TrophylEvent) == 0 {
		t.TrophylEvent = make(map[int]int) //for nil map
	}
	t.TrophylEvent[t.EventId] += trophyEvent

	if t.TrophylEvent[t.EventId] < 0 {
		t.TrophylEvent[t.EventId] = 0
	}

	return t.TrophylEvent[t.EventId]
}

func (t *UserTournamentEventsData) GetTrophyEvent() int {
	return t.TrophylEvent[t.EventId]
}

func (u *UserDailyBattle) IncreaseDailyEventTrophy(eventId, trophy int) int {
	if len(u.DailyTrophyEvent) == 0 {
		u.DailyTrophyEvent = make(map[int]*DailyTrophyModel) //for nil map
	}

	if u.DailyTrophyEvent[u.DailyEventId] == nil {
		u.DailyTrophyEvent[u.DailyEventId] = &DailyTrophyModel{eventId, 0}
	}

	u.DailyTrophyEvent[u.DailyEventId].TrophyEvent += trophy

	if u.DailyTrophyEvent[u.DailyEventId].TrophyEvent < 0 {
		u.DailyTrophyEvent[u.DailyEventId].TrophyEvent = 0
	}

	return u.DailyTrophyEvent[u.DailyEventId].TrophyEvent
}

// inActive if event deactive : check with current event id
// return eventid has reward
func (t *UserTournamentEventsData) CheckReward(lastEventId int, inActive bool) (int, error) {
	if inActive {
		lastEventId--
	}
	for i, _ := range t.TrophylEvent {
		if i <= lastEventId {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Doesn't own any rewards %d ", lastEventId)
}

// inActive if event deactive : check with current event id
func (t *UserTournamentEventsData) CheckRewardDaily(lastEventId int, inActive bool) (int, error) {
	if inActive {
		lastEventId--
	}

	if t.DailyBattleUser == nil {
		return -1, fmt.Errorf("Doesn't own any rewards %d ", lastEventId)
	}

	for i, _ := range t.DailyBattleUser.DailyTrophyEvent {
		if i <= lastEventId {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Doesn't own any rewards %d ", lastEventId)

}
