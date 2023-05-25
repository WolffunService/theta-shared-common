package leaderboardmodel

import (
	"github.com/WolffunService/theta-shared-common/models/usermodel"
	"github.com/WolffunService/theta-shared-common/models/userrankingmodel"
)

type (
	LeaderBroadResponse struct {
		Entries         []*LeaderboardDataModel `json:"entries"`
		MineLeaderboard *LeaderboardDataModel   `json:"mineLb"`
		TotalEntries    int32                   `json:"total"`
	}

	LeaderboardDataModelFullData struct {
		UserModel        usermodel.User               `json:"userModel"`
		UserRankingModel userrankingmodel.UserRanking `json:"userRankingModel"`
		IsBot            *bool                        `json:"isBot,omitempty"`
		Score            float64                      `json:"score"`
		Rank             int32                        `json:"rank"`
	}

	LeaderBroadResponseFullData struct {
		Entries         []*LeaderboardDataModelFullData `json:"entries"`
		MineLeaderboard *LeaderboardDataModelFullData   `json:"mineLb"`
		TotalEntries    int32                           `json:"total"`
	}

	Encourage[T any] struct {
		ClosestRank  int32   `json:"closestRank"`
		Reward       T       `json:"reward"`
		RequireScore float64 `json:"requireScore"`
	}
)
