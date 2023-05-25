package lbrewardbiz

import (
	"context"
	"fmt"
	"time"

	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardstorage"
	"github.com/WolffunService/thetan-shared-common/leaderboard"
	"github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

func TrackLBReward(ctx context.Context, userId string, lbKey *leaderboardmodel.LBKeyModel, score int64, handler interface{}) (*lbrewardmodel.LeaderboardScoreCacheDB, error) {
	var scoreModel *lbrewardmodel.LeaderboardScoreCacheDB
	var err error
	txErr := mongodb.TransactionWithCtx(ctx, func(session mongo.Session, sc mongo.SessionContext) error {
		err := callBeforeTrack(sc, handler)
		if err != nil {
			return err
		}
		//update score db
		scoreModel, err = lbrewardstorage.TrackScore(sc, userId, lbKey, score)
		if err != nil {
			return err
		}

		err = callAfterTrack(sc, handler)
		if err != nil {
			return err
		}
		return session.CommitTransaction(sc)
	})
	if txErr != nil {
		return nil, txErr
	}
	err = leaderboard.TrackLB(ctx, userId, int(scoreModel.Score), lbKey.String())
	if err != nil {
		return nil, err
	}
	//go leaderboard.CheckMineRank(userId, lbKey.String())
	return scoreModel, nil
}

func TrackLBRewardUnsafe(ctx context.Context, userId string, lbKey *leaderboardmodel.LBKeyModel, score int64) (*lbrewardmodel.LeaderboardScoreCacheDB, error) {
	var scoreModel *lbrewardmodel.LeaderboardScoreCacheDB
	var err error
	//update score db
	scoreModel, err = lbrewardstorage.TrackScore(ctx, userId, lbKey, score)
	if err != nil {
		return nil, err
	}

	trackPodiumCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = leaderboard.TrackLB(trackPodiumCtx, userId, int(scoreModel.Score), lbKey.String())
	if err != nil {
		fmt.Println("[error] leaderboard.TrackLB failed", err)
		return scoreModel, nil
	}
	//go leaderboard.CheckMineRank(userId, lbKey.String())
	return scoreModel, nil
}
