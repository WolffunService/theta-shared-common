package lbrewardbiz

import (
	"context"

	"github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"
)

// LBRewardHandler is a handler to manage a leaderboard reward function.
// K is type of reward  (SystemCurrency, List System currency, Model reward ,...)
type LBRewardHandler[T any] interface {
	ExcReward(ctx context.Context, data *leaderboardmodel.LeaderboardDataModelFullData) (T, error)
}
type LBClosestRankHandler[T any] interface {
	//Get closest rank have a reward
	ClosestRank(ctx context.Context, myRank int32) (T, int32, error)
}

func callBeforeClaim(ctx context.Context, reward, handler any) error {
	if handler == nil {
		return nil
	}

	if hook, ok := handler.(BeforeClaimHook); ok {
		if err := hook.BeforeClaim(ctx, reward); err != nil {
			return err
		}
	}
	return nil
}

type BeforeClaimHook interface {
	BeforeClaim(ctx context.Context, reward any) error
}

func callAfterClaim(ctx context.Context, reward, i any) error {
	if i == nil {
		return nil
	}
	if hook, ok := i.(AfterClaimHook); ok {
		if err := hook.AfterClaim(ctx, reward); err != nil {
			return err
		}
	}
	return nil
}

type AfterClaimHook interface {
	AfterClaim(ctx context.Context, reward any) error
}

func callBeforeTrack(ctx context.Context, i interface{}) error {
	if i == nil {
		return nil
	}
	if hook, ok := i.(BeforeTrackHook); ok {
		if err := hook.BeforeTrack(ctx); err != nil {
			return err
		}
	}
	return nil
}

type BeforeTrackHook interface {
	BeforeTrack(ctx context.Context) error
}

func callAfterTrack(ctx context.Context, i interface{}) error {
	if i == nil {
		return nil
	}
	if hook, ok := i.(AfterTrackHook); ok {
		if err := hook.AfterTrack(ctx); err != nil {
			return err
		}
	}
	return nil
}

type AfterTrackHook interface {
	AfterTrack(ctx context.Context) error
}
