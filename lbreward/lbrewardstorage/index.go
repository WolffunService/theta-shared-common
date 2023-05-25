package lbrewardstorage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexScoreModelLB(ctx context.Context) {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				bson.E{Key: "userID", Value: 1},
				bson.E{Key: "lb_key", Value: 1},
			},
		},
		{
			Keys: bson.D{
				bson.E{Key: "userID", Value: 1},
				bson.E{Key: "lb_key", Value: 1},
				bson.E{Key: "lb_id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	if _, err := coll.Indexes().CreateMany(ctx, indexes); err != nil {
		fmt.Println(err.Error())
		return
	}
}
