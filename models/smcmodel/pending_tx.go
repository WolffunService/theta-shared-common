package smcmodel

import "github.com/WolffunService/thetan-shared-common/enums/smcenum"

type PendingTx struct {
	ID        uint64                 `bson:"_id,omitempty"`
	EventName smcenum.EventName      `bson:"eventName,omitempty"`
	RefID     string                 `bson:"refId,omitempty"`
	RawData   map[string]interface{} `bson:"rawData"`
	ExpiredAt int64                  `bson:"expiredAt"`
}

func (PendingTx) CollName() string {
	return "BlockchainPendingTransactions"
}
