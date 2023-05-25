package userrankingmodel

// implement interface Model

func (UserRanking) CollectionName() string {
	return "UserRanking"
}

// GetID method return model's id
func (f *UserRanking) GetID() interface{} {
	return f.ID
}

// SetID set id value of model's id field.
func (f *UserRanking) SetID(id interface{}) {
	f.ID = id
}

// --------------- //

type UserRanking struct {
	ID                      interface{}                            `json:"id" bson:"_id,omitempty"`
	Trophy                  int                                    `json:"trophy" bson:"trophy"`
	TrophyHighest           int                                    `json:"trophyHighest" bson:"trophyHighest"`
	SeasonTrophyHighest     int                                    `json:"seasonTrophyHighest" bson:"seasonTrophyHighest"`
	RankingLevel            int                                    `json:"rankingLevel" bson:"rankingLevel"`
	RankingLevelHighest     int                                    `json:"rankingLevelHighest" bson:"rankingLevelHighest"`
	TrophyCurRank           int                                    `json:"trophyCurRank" bson:"trophyCurRank"`
	SeasonId                int32                                  `json:"seasonId" bson:"seasonId"`
	Rewards                 map[int]map[int]*UserRankingRewardInfo `json:"-" bson:"rewards,omitempty"`
	ListRewards             []int                                  `json:"rewards,omitempty" bson:"-"`
	ListSeasonRewardClaimed []SeasonRewardClaimed                  `json:"listSeasonRewardClaimed,omitempty" bson:"listSeasonRewardClaimed,omitempty"`
	//Division       int                   `json:"division" bson:"division"` //parse from TrophyCurRank + RankingLevel + config
	//				rankingLevl-rewardId
}

type UserRankingRewardInfo struct {
	IsReceived   bool  `json:"isReceived" bson:"isReceived"`
	TimeReceived int64 `json:"timeReceived" bson:"timeReceived"`
}

type SeasonRewardClaimed struct {
	IsReceived   bool  `json:"isReceived" bson:"isReceived"`
	TimeReceived int64 `json:"timeReceived" bson:"timeReceived"`
}
