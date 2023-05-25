package lbrewardstorage

import (
	"context"
	"time"

	"github.com/WolffunService/theta-shared-common/database/mongodb/utils"
	"github.com/WolffunService/theta-shared-common/lbreward/lbrewardmodel"
	"github.com/WolffunService/theta-shared-common/models/leaderboardmodel"

	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/moper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TrackScore(ctx context.Context, userID string, lbKey *leaderboardmodel.LBKeyModel, score int64) (*lbrewardmodel.LeaderboardScoreCacheDB, error) {
	filter := mongoquery.Init(
		mongoquery.Equal("userID", userID),
		mongoquery.Equal("lb_key", lbKey.GetKey()),
		mongoquery.Equal("lb_id", lbKey.GetID()),
		mongoquery.Equal("claimed", false),
	)
	update := mongoquery.Init(
		mongoquery.IncInt(
			mongoquery.PairSetterInt{FieldName: "score", Value: int(score)}))

	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)

	res, err := FindAndUpdateScore(ctx, filter, update, opt)
	if err != nil {
		//if err == mongo.ErrNoDocuments {
		//	return res, insertNewScore(ctx, userID, lbKey, score)
		//}
		return nil, err
	}
	return res, nil
}

//TODO @Tinh update later
//func UnTrackScore(ctx context.Context, userId string, lbKey *leaderboardmodel.LBKeyModel) (*lbrewardmodel.ScoreModel, error) {
//	key := fmt.Sprintf("scores.%s", lbKey)
//	filter := bson.M{"_id": util.ObjectIDFromHex(userId),
//		operator.Exists: key}
//
//	update := mongoquery.Init(
//		bson.E{Key: operator.Unset, Value: key})
//
//	opt := options.FindOneAndUpdate().SetReturnDocument(options.Before)
//
//	res, err := FindAndUpdateScore(ctx, filter, update, opt)
//	if err != nil {
//		return nil, err
//	}
//	return res.GetScoreModel(lbKey), nil
//}

func FindAndUpdateScore(ctx context.Context, filter interface{}, update bson.D, opts ...*options.FindOneAndUpdateOptions) (*lbrewardmodel.LeaderboardScoreCacheDB, error) {
	opts = append(opts, options.FindOneAndUpdate().SetUpsert(true))
	model := &lbrewardmodel.LeaderboardScoreCacheDB{}
	update = utils.BsonSet(update, "lastModified", time.Now())
	err := coll.FindOneAndUpdate(ctx, filter, update, opts...).Decode(model)
	return model, err
}

func Claimed(ctx context.Context, userID string, lbKey *leaderboardmodel.LBKeyModel, reward interface{}) error {
	filter := mongoquery.Init(
		mongoquery.Equal("userID", userID),
		mongoquery.Equal("lb_key", lbKey.GetKey()),
		mongoquery.Equal("lb_id", lbKey.GetID()),
		mongoquery.Equal("claimed", false),
	)

	update := mongoquery.Init(
		mongoquery.Set(
			mongoquery.PairSetter{FieldName: "rewards", Value: reward},
			mongoquery.PairSetter{FieldName: "claimed", Value: true},
		))
	return updateOne(ctx, filter, update)
}

type ClaimCommand struct {
	UserID  string
	LBKey   *leaderboardmodel.LBKeyModel
	Rewards interface{}
}

func BulkClaim(ctx context.Context, cmds []ClaimCommand) (*mongo.BulkWriteResult, error) {
	models := make([]mongo.WriteModel, 0, len(cmds))
	for _, v := range cmds {
		filter := moper.NewD().
			Equal("userID", v.UserID).
			Equal("lb_key", v.LBKey.GetKey()).
			Equal("lb_id", v.LBKey.GetID()).
			Equal("claimed", false)
		update := moper.NewD().Set(moper.P{K: "rewards", V: v.Rewards}, moper.P{K: "claimed", V: true})

		models = append(models, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update))
	}

	return coll.BulkWrite(ctx, models)
}

func updateOne(ctx context.Context, filter interface{}, update bson.D, opts ...*options.UpdateOptions) error {
	update = utils.BsonSet(update, "lastModified", time.Now())
	opts = append(opts, options.Update().SetUpsert(true))
	result, err := coll.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
