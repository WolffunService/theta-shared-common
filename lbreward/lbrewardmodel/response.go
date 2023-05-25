package lbrewardmodel

import "github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"

type (
	PlayerInfo struct {
		UserId      string `json:"userId"`
		Username    string `json:"username"`
		Trophy      int    `json:"trophy"`
		TrophyLevel int    `json:"trophyLevel"`
		AvatarId    int    `json:"avatarId"`
		GuildName   string `json:"guildName"`
		FrameId     int    `json:"frameId"`
		Country     string `json:"country"`
	}

	//LeaderboardDataModel T -> type data reward
	LeaderboardDataModel[T any] struct {
		ProfileInfo PlayerInfo `json:"profileInfo"`
		Score       float64    `json:"score"`
		Rank        int32      `json:"rank"`
		Reward      T          `json:"reward"`
	}
	LeaderBroadResponse[T any] struct {
		Entries         []LeaderboardDataModel[T]      `json:"entries"`
		MineLeaderboard *LeaderboardDataModel[T]       `json:"mineLb"`
		TotalEntries    int32                          `json:"total"`
		Encourage       *leaderboardmodel.Encourage[T] `json:"encourage"`
	}

	LeaderBroadResponseFullDataWrap struct {
		Success bool                                          `json:"success"`
		Data    *leaderboardmodel.LeaderBroadResponseFullData `json:"data"`
	}
	LeaderboardDataModelFullDataWrap struct {
		Success bool                                           `json:"success"`
		Data    *leaderboardmodel.LeaderboardDataModelFullData `json:"data"`
	}

	ResponseCheckReward struct {
		ScoreData *LeaderboardScoreCacheDB                       `json:"scoreData"`
		LBData    *leaderboardmodel.LeaderboardDataModelFullData `json:"lbData"`
		LBKey     *leaderboardmodel.LBKeyModel                   `json:"lbKey"`
	}
)
