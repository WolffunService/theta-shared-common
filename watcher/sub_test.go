package watcher

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewMongoWatcherNoPass(t *testing.T) {
	db := &mongo.Database{}
	service := ""
	topic := ""
	w, err := NewMongoWatcher(db, service, topic, handler)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(*w.Options.Limit)

	limit := int64(1003)
	option := &Options{Limit: &limit}
	w2, err := NewMongoWatcher(db, service, topic, handler, option)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(*w2.Options.Limit)

	limit3 := int64(1003)
	limit2 := int64(1006)
	option3 := &Options{Limit: &limit3}
	option2 := &Options{Limit: &limit2}
	w3, err := NewMongoWatcher(db, service, topic, handler, option3, option2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(*w3.Options.Limit)
}

func handler(sessionCtx context.Context, val interface{}) error {
	return nil
}
