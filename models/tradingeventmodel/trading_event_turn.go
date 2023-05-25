package tradingeventmodel

import (
	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ---------- INPUT
type TradingEventTurn struct {
	mongodb.DefaultModel `bson:",inline"`
	mongodb.DateFields   `json:"-" bson:",inline"`

	UserID   string `json:"-" bson:"userId"`
	SeasonID int64  `json:"seasonId" bson:"seasonId"`

	// reward
	CashbackMilestones map[int]bool   `json:"cbMilestones" bson:"cbMilestones"` // true là có thể claim, vì false là giá trị default cho những milestone không thể claim
	DailyChallenge     DailyChallenge `json:"daily" bson:"daily"`
	ClaimedAll         bool           `json:"-" bson:"claimedAll"`
}

func (h TradingEventTurn) GetStrID() string {
	return h.ID.(primitive.ObjectID).Hex()
}

func (TradingEventTurn) CollectionName() string {
	return "TradingEventTurns"
}
