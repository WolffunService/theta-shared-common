package tradingeventmodel

import (
	"github.com/WolffunService/theta-shared-common/enums/tradingeventenum"
	"github.com/WolffunService/theta-shared-common/models/currencymodel"
)

type DailyChallenge struct {
	DayID int `json:"dayId" bson:"dayId"`

	Challenges     []DailyChallengeData         `json:"challenges" bson:"challenges"`
	PreviousReward currencymodel.SystemCurrency `json:"preReward" bson:"preReward"`
}

type DailyChallengeData struct {
	RefreshTime    int                             `json:"-" bson:"refreshTime"`
	Type           tradingeventenum.DailyQuestType `json:"type" bson:"type"`
	VolumeRequired currencymodel.SystemCurrency    `json:"required" bson:"required"`
	RefreshPrice   currencymodel.SystemCurrency    `json:"refreshPrice" bson:"refreshPrice"`
	Score          currencymodel.SystemCurrency    `json:"score" bson:"score"`

	Reward currencymodel.SystemCurrency `json:"reward" bson:"reward"`
	Status tradingeventenum.ClaimStatus `json:"status" bson:"status"`
}
