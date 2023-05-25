package specialeventmodel

import "github.com/WolffunService/thetan-shared-common/models/currencymodel"

type RewardResponse struct {
	Rank         int32     `json:"rank"`
	Score        int       `json:"score"`
	EventId      int       `json:"eventId"`
	DailyEventId int       `json:"dailyEventId"`
	Rewards      RewardsSE `json:"rewards"`
	IsDaily      bool      `json:"-"`
}

type RewardsSE struct {
	BaseReward  currencymodel.SystemCurrency `json:"baseReward"`
	BonusReward currencymodel.SystemCurrency `json:"bonusReward"`
}

type RewardsAnalytic struct {
	RankBonusRate  float64
	GTHGRateReward float64
	SharedRate     float64
	RankingLevel   int
}

func (r *RewardResponse) UpdateReward(eventData *TournamentEventData, isDaily bool, trophyLevel int) RewardsAnalytic {
	rewardsAnalytic := RewardsAnalytic{}
	r.Rewards, rewardsAnalytic = eventData.GetGTHGReward(isDaily, int(r.Rank), trophyLevel, r.DailyEventId)
	return rewardsAnalytic
}
