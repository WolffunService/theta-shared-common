package lbrewardstorage

import (
	"context"
	"time"

	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/enums/leaderboardenum"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"
	"github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"
)

func insertNewScore(ctx context.Context, userID string, lbKey *leaderboardmodel.LBKeyModel, score int64) error {
	model := &lbrewardmodel.LeaderboardScoreCacheDB{
		UserID:  userID,
		LBID:    lbKey.GetID(),
		LBKey:   leaderboardenum.LBKeyPrefix(lbKey.GetKey()),
		Score:   score,
		Claimed: false,
	}
	now := time.Now().UTC()
	model.CreatedAt = now
	model.UpdatedAt = now
	col := mongodb.Coll(model)

	return col.CreateWithCtx(ctx, model)
}
