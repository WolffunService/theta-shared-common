package watcher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/WolffunService/thetan-shared-common/database/mredis/thetanlock"
	"github.com/WolffunService/thetan-shared-common/lockkey"
	"github.com/WolffunService/thetan-shared-common/models/mongopubsubmodel"

	"github.com/go-co-op/gocron"
	"github.com/go-redsync/redsync/v4"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	retryOption   = redsync.WithTries(1)
	timeoutOption = redsync.WithExpiry(time.Second)
)

const (
	defaultLimit int64 = 100
)

type EventDocument struct {
	//OperationType string               `bson:"operationType"`
	FullDocument mongopubsubmodel.MongoPSubModel `bson:"fullDocument"`
}

type MongoWatcher struct {
	Service string
	Topic   string
	Fn      func(sessionCtx context.Context, val interface{}) error
	Coll    *mgm.Collection
	client  *mongo.Client
	Options *Options
}

type Options struct {
	Limit *int64
}

func mergeOptions(opts ...*Options) *Options {
	merged := &Options{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Limit != nil {
			merged.Limit = opt.Limit
		}
	}
	return merged
}

func newDefaultOption() *Options {
	limit := defaultLimit
	return &Options{
		Limit: &limit,
	}
}

func NewMongoWatcher(db *mongo.Database, service string, topic string, fn func(sessionCtx context.Context, val interface{}) error, options ...*Options) (*MongoWatcher, error) {
	coll := mgm.NewCollection(db, mongopubsubmodel.MongoPSubModel{}.CollectionName())
	if coll == nil {
		err := errors.New("watcher: cannot connect to mongodb")
		return nil, err
	}

	client := db.Client()
	if client == nil {
		err := errors.New("watcher: client is nil")
		return nil, err
	}

	m := &MongoWatcher{
		Service: service,
		Topic:   topic,
		Fn:      fn,
		Coll:    coll,
		client:  db.Client(),
	}

	if len(options) == 0 || options == nil {
		m.Options = newDefaultOption()
	} else {
		m.Options = mergeOptions(options...)
	}

	return m, nil
}

func (m MongoWatcher) Start() {
	m.createIndexes()

	go m.startWatching()

	m.startJob()

	return
}

func (m MongoWatcher) createIndexes() error {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{"status", 1},
			},
		},
	}

	ctx := context.Background()
	if _, err := m.Coll.Indexes().CreateMany(ctx, indexes); err != nil {
		return err
	}

	return nil
}

func (m MongoWatcher) startWatching() {
	matchPipeline := bson.D{
		{
			"$match", bson.D{
				{"operationType", "insert"}, //only stream watch type insert new document
				{"fullDocument.service", bson.D{
					{"$eq", m.Service},
				}},
				{"fullDocument.topic", bson.D{
					{"$eq", m.Topic},
				}},
			},
		},
	}

	stream, err := m.Coll.Watch(context.Background(), mongo.Pipeline{matchPipeline})
	if err != nil {
		log.Printf("watcher: error when calling `mongo watch`: %v", err)
		return
	}
	defer stream.Close(context.Background())

	for stream.Next(context.TODO()) {
		var someRecord EventDocument
		if err := stream.Decode(&someRecord); err != nil {
			log.Printf("watcher: error while watching stream: decode error: %s", err.Error())
			continue
		}

		document := someRecord.FullDocument

		err = m.callbackAndUpdateStatus(context.Background(), document)

		if err != nil {
			log.Printf("watcher: error when doing callbackAndUpdateStatus: %s. document _id: %v", err.Error(), document.ID)
			continue
		}
	}
}

func (m MongoWatcher) startJob() {
	cronjob := gocron.NewScheduler(time.UTC)
	// todo co the dua interval vao config
	cronjob.Every(10 + rand.Intn(5)).Second().Do(m.crawlAndProcess)
	cronjob.StartAsync()
	return
}

func (m MongoWatcher) crawlAndProcess() {
	// query unprocessed documents
	option := options.FindOptions{
		Limit: m.Options.Limit,
	}
	cursor, err := m.Coll.Find(
		context.Background(),
		bson.M{"status": bson.M{operator.Eq: false}},
		&option,
	)
	if err != nil {
		return
	}

	// process them
	for cursor.Next(context.Background()) {
		var result mongopubsubmodel.MongoPSubModel
		if err := cursor.Decode(&result); err != nil {
			log.Printf("watcher: job decode error: %s", err.Error())
			continue
		}

		err = m.callbackAndUpdateStatus(context.Background(), result)

		if err != nil {
			log.Printf("watcher: job: callbackAndUpdateStatus error: %s. document _id: %v", err.Error(), result.ID)
			continue
		}
	}
}

func (m MongoWatcher) callbackAndUpdateStatus(ctx context.Context, doc mongopubsubmodel.MongoPSubModel) error {
	// Note:
	// - timeout cần lớn hơn timeout của mgm.TransactionWithClient. Tạm để 30s
	// - try 1 lần thôi bởi goroutine khác đang xử lý thì để nó xử lý thôi
	lKey := lockkey.GetKey(lockkey.LockDocumentToProcess, doc.ID.String())
	lock, err := thetanlock.LockCustom(lKey, retryOption, timeoutOption)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	return mgm.TransactionWithClient(ctx, m.client, func(session mongo.Session, sc mongo.SessionContext) error {
		// Ack message
		filter := bson.D{
			{"_id", doc.ID},
			{"status", false},
		}
		update := bson.D{
			{"$set", bson.D{
				{"status", true},
			}},
		}
		if result, err := m.Coll.UpdateOne(sc, filter, update); err != nil {
			return err
		} else {
			if result.ModifiedCount == 0 {
				return fmt.Errorf("document is not updated")
			}
		}

		// Trigger callback
		if err := m.Fn(sc, doc.Data); err != nil {
			return err
		}
		return session.CommitTransaction(sc)
	})
}
