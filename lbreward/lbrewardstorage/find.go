package lbrewardstorage

import (
	"context"
	"errors"

	"github.com/WolffunService/thetan-shared-common/enums/leaderboardenum"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"

	"github.com/func25/mongofunc/mongoquery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNoScoreData = errors.New("no data")
)

func FindScore(ctx context.Context, userID string, lbKeyPrefix leaderboardenum.LBKeyPrefix) (*lbrewardmodel.LeaderboardScoreCacheDB, error) {
	var models []lbrewardmodel.LeaderboardScoreCacheDB
	filter := bson.D{
		{Key: "userID", Value: userID},
		{Key: "lb_key", Value: lbKeyPrefix},
		{Key: "claimed", Value: false},
	}
	// opts := options.Find().SetSort(bson.D{{"priority", 1}})
	err := coll.SimpleFindWithCtx(ctx, &models, filter) //, opts)
	if err != nil {
		return nil, err
	}
	for _, e := range models {
		return &e, nil
	}

	return nil, ErrNoScoreData
}

type FindScoresQuery struct {
	Lb_Key   *string
	Lb_Id    *int64
	Claimed  *bool
	LB_ID_In []int64
}

func (q FindScoresQuery) ToFilter(userID string) primitive.D {
	filter := mongoquery.Init(
		mongoquery.Equal("userID", userID),
	)

	if nil != q.Claimed {
		filter = append(filter, mongoquery.Equal("claimed", *q.Claimed))
	}

	if nil != q.Lb_Id {
		filter = append(filter, mongoquery.Equal("lb_id", *q.Lb_Id))
	}

	if nil != q.Lb_Key {
		filter = append(filter, mongoquery.Equal("lb_key", *q.Lb_Key))
	}

	if nil != q.LB_ID_In {
		filter = append(filter, mongoquery.InArray("lb_id", q.LB_ID_In))
	}

	return filter
}

func FindScoreCustom(ctx context.Context, userID string, q FindScoresQuery) (*lbrewardmodel.LeaderboardScoreCacheDB, error) {
	opts := options.Find().SetSort(bson.D{{Key: "priority", Value: 1}})

	var models []lbrewardmodel.LeaderboardScoreCacheDB
	err := coll.SimpleFindWithCtx(ctx, &models, q.ToFilter(userID), opts)
	if err != nil {
		return nil, err
	}

	for _, e := range models {
		return &e, nil
	}

	return nil, ErrNoScoreData
}

func FindManyScores(ctx context.Context, userID string, q FindScoresQuery) ([]lbrewardmodel.LeaderboardScoreCacheDB, error) {
	var models []lbrewardmodel.LeaderboardScoreCacheDB
	err := coll.SimpleFindWithCtx(ctx, &models, q.ToFilter(userID))
	if err != nil {
		return nil, err
	}

	return models, nil
}
