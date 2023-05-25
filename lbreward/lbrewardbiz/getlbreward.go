package lbrewardbiz

import (
	"context"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"
	"github.com/WolffunService/thetan-shared-common/leaderboard"

	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardapi"
	"github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"
	"github.com/WolffunService/thetan-shared-common/models/usermodel"
	podium "github.com/WolffunService/thetan-shared-common/proto/podium"
)

//TODO profiler data - optimize funcz

// GetLeaderboardReward
// -lbKey- leaderboard key
// -limit - skip option find
// -RewardFunc- function update reward
func GetLeaderboardReward[T any](ctx context.Context, req lbrewardmodel.RequestGetLBReward, handler LBRewardHandler[T]) (*lbrewardmodel.LeaderBroadResponse[T], error) {
	//const op = "GetLeaderboardReward"

	reqGetLBData := &leaderboardmodel.RequestGetLBFullData{
		UserID:     req.UserID,
		LBKey:      req.LBKey,
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
	}

	lbres, err := lbrewardapi.GetLeaderboardFullData(reqGetLBData)
	if err != nil {
		//logger.Error().Err(err).Op(op).Msg("get leaderboard full data")
		return nil, err
	}

	response := &lbrewardmodel.LeaderBroadResponse[T]{}

	// Loop for each entry to update reward
	numEntries := len(lbres.Entries)
	entries := make([]lbrewardmodel.LeaderboardDataModel[T], 0, numEntries)
	if lbres.Entries != nil && numEntries > 0 {
		for _, e := range lbres.Entries {
			entry := parseLBData[T](e)
			reward, err := handler.ExcReward(ctx, e)
			if err != nil {
				//logger.Error().Err(err).Op(op).Msg("entity exc reward")
				return nil, err
			}
			entry.Reward = reward
			entries = append(entries, *entry)
		}
	}

	// Also update reward for me
	mine := lbres.MineLeaderboard
	if mine != nil {
		response.MineLeaderboard = parseLBData[T](mine)
		reward, err := handler.ExcReward(ctx, mine)
		if err != nil {
			//logger.Error().Err(err).Op(op).Msg("mine exc reward")
			return nil, err
		}
		response.MineLeaderboard.Reward = reward
		//go checkMineRank(req.UserID, req.LBKey, mine.Rank)
	}

	//closest rank reward
	mineRankTemp := lbres.TotalEntries + 1
	mineScore := float64(0)
	if mine != nil {
		mineScore = mine.Score
		if mine.Rank > 1 {
			mineRankTemp = mine.Rank
		}
	}
	if mineRankTemp != 0 {
		if hook, ok := handler.(LBClosestRankHandler[T]); ok {
			if reward, rank, err := hook.ClosestRank(ctx, mineRankTemp); err == nil {
				req := &podium.GetTopMembersRequest{
					LeaderboardId: req.LBKey,
					PageNumber:    rank,
					PageSize:      1,
				}
				if resLb, err := leaderboard.GetTopLeaderboard(ctx, req); err == nil {
					members := resLb.GetMembers()
					if len(members) != 0 && members[0] != nil {
						response.Encourage = &leaderboardmodel.Encourage[T]{
							Reward:       reward,
							ClosestRank:  rank,
							RequireScore: members[0].Score - mineScore,
						}
						if response.Encourage.RequireScore < 1 {
							response.Encourage.RequireScore = 1
						}
					}
				}
			}
		}
	}
	response.Entries = entries
	response.TotalEntries = lbres.TotalEntries
	return response, nil
}

func parseLBData[T any](model *leaderboardmodel.LeaderboardDataModelFullData) *lbrewardmodel.LeaderboardDataModel[T] {
	result := &lbrewardmodel.LeaderboardDataModel[T]{}
	if model == nil {
		return result
	}
	result.ProfileInfo = getUserProfileInfo(&model.UserModel)
	if &model.UserRankingModel != nil {
		result.ProfileInfo.Trophy = model.UserRankingModel.Trophy
		result.ProfileInfo.TrophyLevel = model.UserRankingModel.RankingLevel
	}
	result.Rank = model.Rank
	result.Score = model.Score
	return result
}

func getUserProfileInfo(user *usermodel.User) lbrewardmodel.PlayerInfo {
	playerInfo := lbrewardmodel.PlayerInfo{}
	if user == nil {
		return playerInfo
	}
	playerInfo.UserId = user.ID.(string)
	playerInfo.Username = user.UserName
	playerInfo.AvatarId = user.AvatarId
	playerInfo.Country = user.Country

	playerInfo.GuildName = ""
	playerInfo.FrameId = user.FrameId

	return playerInfo
}
