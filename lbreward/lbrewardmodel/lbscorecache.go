package lbrewardmodel

import (
	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/enums/leaderboardenum"
)

const (
	LeaderboardScoreCacheDBName = "LeaderboardScoreCache"
)

//create index

func (LeaderboardScoreCacheDB) CollectionName() string {
	return LeaderboardScoreCacheDBName
}

type LeaderboardScoreCacheDB struct {
	mongodb.DefaultModel `json:"-" bson:",inline"`
	mongodb.DateFields   `json:"-" bson:",inline"`
	UserID               string                      `json:"userID" bson:"userID"` //pk
	LBKey                leaderboardenum.LBKeyPrefix `json:"lb_key" bson:"lb_key"` //pk
	LBID                 int                         `json:"lb_id" bson:"lb_id"`   //pk
	Score                int64                       `json:"score" bson:"score"`
	Claimed              bool                        `json:"claimed" bson:"claimed"`
	Priority             int                         `json:"priority" bson:"priority,omitempty"` //TODO @Tinh increase when claimaction has error occur
	Rewards              interface{}                 `json:"rewards" bson:"rewards"`
}
