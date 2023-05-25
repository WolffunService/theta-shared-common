package lbrewardbiz

import (
	"context"

	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardapi"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardstorage"
)

//var logger = thetalog.NewBizLogger("leaderboard_reward_business")

func Init(ctx context.Context, lbStorageConf lbrewardstorage.Config, internalAddress string) error {
	if err := lbrewardapi.Init(internalAddress); err != nil {
		return err
	}

	lbrewardstorage.Init(lbStorageConf)
	lbrewardstorage.CreateIndexScoreModelLB(ctx)

	return nil
}
