package receiptdatamodel

import (
	"encoding/json"
	"github.com/WolffunService/theta-shared-common/iapvalidator"

	platformtype "github.com/WolffunService/theta-shared-common/enums/platformtypeenum"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReceiptShopInStore struct {
	ID            primitive.ObjectID        `json:"id" bson:"_id,omitempty"`
	PlatformType  platformtype.PlatformType `bson:"platformType"`
	TransactionId string                    `bson:"transactionId"`
	UserId        string                    `bson:"userId"`
	PackId        int                       `bson:"packId"`
	Data          any                       `bson:"data"`
	GameName      string                    `bson:"gameName"`
}

func (model ReceiptShopInStore) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method returns a model's ID
func (model ReceiptShopInStore) GetID() interface{} {
	return model.ID
}

// SetID sets the value of a model's ID field.
func (model ReceiptShopInStore) SetID(id interface{}) {
	model.ID = id.(primitive.ObjectID)
}

func (model ReceiptShopInStore) CollectionName() string {
	return "ReceiptShopInStore"
}

func NewReceiptShopInStore(userId string, buyRequest iapvalidator.BuyItemReq, gameName string) *ReceiptShopInStore {
	receipt := &ReceiptShopInStore{}
	receipt.PlatformType = buyRequest.PlatformType
	receipt.TransactionId, _ = buyRequest.GetTransactionId()
	receipt.UserId = userId
	receipt.PackId = buyRequest.PackId
	var data any
	json.Unmarshal([]byte(buyRequest.ReceiptJson), &data)
	receipt.Data = data
	receipt.GameName = gameName
	return receipt
}
