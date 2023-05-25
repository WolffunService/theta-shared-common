package auditmodel

import (
	"github.com/WolffunService/thetan-shared-common/enums/currencyenum"
	"github.com/WolffunService/thetan-shared-common/models/currencymodel"
	"github.com/WolffunService/thetan-shared-common/proto/auditproto"
)

type BalanceCreator struct{}

type BalanceEvent struct {
	ID             interface{}                  `json:"id" bson:"_id,omitempty"`
	UserId         string                       `json:"userId" bson:"userId"`
	Timestamp      int64                        `json:"timestamp" bson:"timestamp"`
	PreValue       int64                        `json:"-" bson:"preValue"`
	NewValue       int64                        `json:"-" bson:"newValue"`
	Source         int32                        `json:"source" bson:"source"`
	ChangedBalance currencymodel.SystemCurrency `json:"changedBalance" bson:"changedbalance"`
}

func (BalanceCreator) NewBalanceEvent(res *auditproto.CurrencyChangeResult) (*BalanceEvent, bool) {
	if res.ErrorCode != -1 || len(res.ErrorMsg) > 0 {
		return nil, false
	}

	return &BalanceEvent{
		UserId:    res.UserId,
		Timestamp: res.Timestamp,
		ChangedBalance: currencymodel.SystemCurrency{
			Type:     currencyenum.Currency(res.CurrencyType),
			Value:    res.ChangeValue,
			Decimals: int(res.Decimals),
		},
		PreValue: res.PreValue,
		NewValue: res.NextValue,
		Source:   res.Source,
	}, true
}
