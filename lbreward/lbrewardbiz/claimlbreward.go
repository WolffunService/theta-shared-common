package lbrewardbiz

import (
	"context"
	"errors"

	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardapi"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardstorage"
	"github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrClaimed = errors.New("reward claimed")
)

func ClaimReward[T any](ctx context.Context, req *lbrewardmodel.ReqClaimLBReward, handler LBRewardHandler[T]) (T, error) {
	res, errCheck := CheckReward(ctx, req)
	if errCheck != nil {
		return *new(T), errCheck
	}

	if res.ScoreData.Claimed {
		return *new(T), ErrClaimed
	}
	//if scoreModel.Score != lbRes.Score??

	var reward T
	var err error
	txErr := mongodb.TransactionWithCtx(ctx, func(session mongo.Session, sc mongo.SessionContext) error {
		reward, err = handler.ExcReward(sc, res.LBData)
		if err != nil {
			return err
		}

		err = callBeforeClaim(sc, reward, handler)
		if err != nil {
			return err
		}

		err = lbrewardstorage.Claimed(sc, req.UserID, res.LBKey, reward)
		if err != nil {
			return err
		}

		err = callAfterClaim(sc, reward, handler)
		if err != nil {
			return err
		}
		return session.CommitTransaction(sc)
	})
	if txErr != nil {
		return *new(T), txErr
	}

	return reward, nil
}

func ClaimRewardUnsafe[T any](ctx context.Context, req *lbrewardmodel.ReqClaimLBReward, handler LBRewardHandler[T]) (T, error) {
	res, errCheck := CheckReward(ctx, req)
	if errCheck != nil {
		return *new(T), errCheck
	}

	if res.ScoreData.Claimed {
		return *new(T), ErrClaimed
	}
	//if scoreModel.Score != lbRes.Score??

	var reward T
	var err error
	if handler != nil {
		reward, err = handler.ExcReward(ctx, res.LBData)
		if err != nil {
			return *new(T), err
		}
	}

	err = lbrewardstorage.Claimed(ctx, req.UserID, res.LBKey, reward)
	if err != nil {
		return *new(T), err
	}

	return reward, nil
}

type ClaimManyCommand[T any] struct {
	Handler LBRewardHandler[T]
	Record  *lbrewardmodel.ResponseCheckReward
}

func ClaimManyRewardsUnsafe[T any](ctx context.Context, data []ClaimManyCommand[T]) (int, error) {
	var err error

	cmds := make([]lbrewardstorage.ClaimCommand, 0, len(data))

	for i := range data {
		rewards, err := data[i].Handler.ExcReward(ctx, data[i].Record.LBData)
		if err != nil {
			return 0, err
		}

		cmds = append(cmds, lbrewardstorage.ClaimCommand{
			UserID:  data[i].Record.ScoreData.UserID,
			LBKey:   data[i].Record.LBKey,
			Rewards: rewards,
		})
	}

	result, err := lbrewardstorage.BulkClaim(ctx, cmds)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), err
}

// CheckReward will return ErrNoData if not found any score
func CheckReward(ctx context.Context, req *lbrewardmodel.ReqClaimLBReward) (*lbrewardmodel.ResponseCheckReward, error) {
	x := string(req.LBRewardKey)
	scoreModel, err := lbrewardstorage.FindScoreCustom(ctx, req.UserID, lbrewardstorage.FindScoresQuery{
		Lb_Key:   &x,
		Lb_Id:    req.LBRewardID,
		Claimed:  req.Claimed,
		LB_ID_In: req.InLBRewardID,
	})

	if err != nil {
		return nil, err
	}

	lbKey := leaderboardmodel.NewLBKeyModel(req.LBRewardKey, scoreModel.LBID)
	var reqGetLBData = &leaderboardmodel.RequestGetLBFullData{
		UserID: req.UserID,
		LBKey:  lbKey.String(),
	}

	lbRes, errGetMember := lbrewardapi.GetMemberLeaderboardFullData(reqGetLBData)
	if errGetMember != nil {
		return nil, errGetMember
	}

	return &lbrewardmodel.ResponseCheckReward{
		LBData:    lbRes,
		ScoreData: scoreModel,
		LBKey:     lbKey,
	}, nil
}

// CheckReward will return ErrNoData if not found any score
func CheckManyRewards(ctx context.Context, req lbrewardmodel.ReqClaimLBReward) (rewards []lbrewardmodel.ResponseCheckReward, err error) {
	_lbRewardKey := string(req.LBRewardKey)
	scoreModels, err := lbrewardstorage.FindManyScores(ctx, req.UserID, lbrewardstorage.FindScoresQuery{
		Lb_Key:   &_lbRewardKey,
		Lb_Id:    req.LBRewardID,
		Claimed:  req.Claimed,
		LB_ID_In: req.InLBRewardID,
	})

	if err != nil {
		return nil, err
	}

	for i := range scoreModels {
		lbKey := leaderboardmodel.NewLBKeyModel(req.LBRewardKey, scoreModels[i].LBID)
		var reqGetLBData = &leaderboardmodel.RequestGetLBFullData{
			UserID: req.UserID,
			LBKey:  lbKey.String(),
		}

		lbRes, err := lbrewardapi.GetMemberLeaderboardFullData(reqGetLBData)
		if err != nil {
			return nil, err
		}
		rewards = append(rewards, lbrewardmodel.ResponseCheckReward{
			LBData:    lbRes,
			ScoreData: &scoreModels[i],
			LBKey:     lbKey,
		})
	}

	return
}
