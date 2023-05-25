package lbrewardmodel

import "github.com/WolffunService/theta-shared-common/enums/leaderboardenum"

type (
	ReqClaimLBReward struct {
		UserID       string
		LBRewardKey  leaderboardenum.LBKeyPrefix //is prefix key leaderboard events
		LBRewardID   *int64
		Claimed      *bool
		InLBRewardID []int64 // less then leaderboard reward id
	}

	RequestGetLBReward struct {
		UserID     string `json:"userID"`
		LBKey      string `json:"LBKey"`
		PageSize   int32  `json:"pageSize"`
		PageNumber int32  `json:"pageNumber"`
	}
)
