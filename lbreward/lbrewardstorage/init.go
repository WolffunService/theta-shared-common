package lbrewardstorage

import (
	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

var coll *mongodb.Collection

type Config struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func Init(conf Config) {
	mongodb.SetConfig(conf.DB, conf.Client)
	coll = mongodb.Coll(&lbrewardmodel.LeaderboardScoreCacheDB{})
}
