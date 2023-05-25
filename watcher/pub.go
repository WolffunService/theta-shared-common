package watcher

import (
	"context"

	"github.com/WolffunService/thetan-shared-common/models/mongopubsubmodel"

	"github.com/kamva/mgm/v3"
)

func Publish(ctx context.Context, service string, topic string, data interface{}) error {
	// init record's fields
	record := mongopubsubmodel.MongoPSubModel{
		Service: service,
		Topic:   topic,
		Data:    data,
		Status:  false,
	}

	// insert to coll
	col := mgm.Coll(&mongopubsubmodel.MongoPSubModel{})
	_, err := col.InsertOne(ctx, record)

	return err
}
