package mongopubsubmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoPSubModel struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Service string             `json:"service" bson:"service"`
	Topic   string             `json:"topic" bson:"topic"`
	Data    interface{}        `json:"data" bson:"data"`
	Status  bool               `json:"status" bson:"status"`
}

func (m MongoPSubModel) CollectionName() string {
	return "MongoPsubs"
}

// PrepareID method prepares the ID value to be used for filtering
// e.g convert hex-string ID value to bson.ObjectId
func (m *MongoPSubModel) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method returns a model's ID
func (m *MongoPSubModel) GetID() interface{} {
	return m.ID
}

// SetID sets the value of a model's ID field.
func (m *MongoPSubModel) SetID(id interface{}) {
	m.ID = id.(primitive.ObjectID)
}
