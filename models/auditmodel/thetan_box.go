package auditmodel

import (
	"log"
	"strconv"

	"github.com/WolffunService/thetan-shared-common/enums/currencyenum"
	"github.com/WolffunService/thetan-shared-common/enums/thetanboxenum"
	"github.com/WolffunService/thetan-shared-common/models/auditmodel/auditconst"
	"github.com/WolffunService/thetan-shared-common/models/currencymodel"
	"github.com/WolffunService/thetan-shared-common/proto/auditproto"
)

type ThetanBoxAuditCreator struct{}

type BoxEvent struct {
	ID               interface{}                  `json:"id" bson:"_id,omitempty"`
	UserId           string                       `json:"userId" bson:"userId"`
	BoxId            int64                        `json:"boxId" bson:"boxId"`
	BoxType          int32                        `json:"boxType" bson:"boxType"`
	Source           thetanboxenum.BoxEventSource `json:"source" bson:"source"`
	Timestamp        int64                        `json:"timestamp" bson:"timestamp"`
	BoxPurchaseEvent *BoxPurchaseEvent            `json:"boxPurchaseEvent,omitempty" bson:"boxPurchaseEvent,omitempty"`
}

type BoxPurchaseEvent struct {
	Price         currencymodel.SystemCurrency `json:"price" bson:"price"`
	Amount        int32                        `json:"amount" bson:"amount"`
	TransactionId string                       `json:"-" bson:"transactionId"`
}

// NewOpenEvent return BoxEvent after convert and a bool determine that you need write event to database
func (ThetanBoxAuditCreator) NewOpenEvent(res *auditproto.BoxOpenResult) (*BoxEvent, bool) {
	if res.ErrorCode != -1 || len(res.ErrorMsg) > 0 {
		return nil, false
	}

	return &BoxEvent{
		UserId:    res.UserId,
		BoxType:   res.BoxType,
		Timestamp: res.Timestamp,
		Source:    thetanboxenum.BES_OPEN,
		BoxId:     res.BoxId,
	}, true
}

func (ThetanBoxAuditCreator) NewBoxEvent(e *auditproto.SimpleEvent) (*BoxEvent, bool) {
	auditMap := map[string]string{}
	if e.Event != nil {
		for _, v := range e.Event.EventParams {
			auditMap[v.Key] = v.Value
		}
	}

	// get boxType
	boxType, err := strconv.Atoi(auditMap["boxType"])
	if err != nil {
		log.Println("[error][audit][box-event] cannot get boxType", err)
		boxType = -1
	}

	// get box event source
	var source thetanboxenum.BoxEventSource
	var boxPurchaseEvent *BoxPurchaseEvent = nil

	switch e.Event.EventName {
	case auditconst.EventName.EN_BOX_PURCHASE:
		source = thetanboxenum.BES_PURCHASE

		// get currencyType
		currencyType, err := stringToInt(auditMap["currencyType"], -1)
		if err != nil {
			log.Println("[error][audit][box-event] cannot get currencyType", err)
		}

		// get price
		price, err := stringToInt(auditMap["price"], -1)
		if err != nil {
			log.Println("[error][audit][box-event] cannot get price", err)
		}

		// get decimal
		decimal, err := stringToInt(auditMap["decimal"], -1)
		if err != nil {
			log.Println("[error][audit][box-event] cannot get decimal", err)
		}

		// get amount
		amount, err := stringToInt(auditMap["boxNumber"], -1)
		if err != nil {
			log.Println("[error][audit][box-event] cannot get amount", err)
		}

		boxPurchaseEvent = &BoxPurchaseEvent{
			Price: currencymodel.SystemCurrency{
				Type:     currencyenum.Currency(currencyType),
				Value:    int64(price),
				Decimals: decimal,
			},
			Amount:        int32(amount),
			TransactionId: auditMap["transactionHash"],
		}
	case auditconst.EventName.EN_BOX_OPEN:
		source = thetanboxenum.BES_OPEN
	default:
		return nil, false
	}

	// get boxId
	boxId, err := strconv.Atoi(auditMap["boxId"])
	if err != nil {
		log.Println("[error][audit][box-event] cannot get boxId", err)
		boxId = -1
	}

	return &BoxEvent{
		UserId:           auditMap["userId"],
		BoxType:          int32(boxType),
		Timestamp:        e.Event.Timestamp,
		Source:           source,
		BoxId:            int64(boxId),
		BoxPurchaseEvent: boxPurchaseEvent,
	}, true
}

// NewPurchaseEvent return BoxEvent after convert and a bool determine that you need write event to database
func (ThetanBoxAuditCreator) NewPurchaseEvent(res *auditproto.BoxPurchaseResult) (*BoxEvent, bool) {
	if res.ErrorCode != -1 || len(res.ErrorMsg) > 0 {

		return nil, false
	}

	return &BoxEvent{
		UserId:    res.UserId,
		BoxType:   res.BoxType,
		Timestamp: res.Timestamp,
		Source:    thetanboxenum.BES_PURCHASE,
		BoxId:     res.BoxId,
		BoxPurchaseEvent: &BoxPurchaseEvent{
			Price: currencymodel.SystemCurrency{
				Value:    res.Price,
				Type:     currencyenum.Currency(res.CurrencyType),
				Decimals: int(res.Decimals),
			},
			Amount:        res.BoxNumber,
			TransactionId: res.TransactionHash,
		},
	}, true
}

func (ThetanBoxAuditCreator) NewNormalEvent(res *auditproto.BoxEvent) (*BoxEvent, bool) {
	return &BoxEvent{
		UserId:    res.UserId,
		BoxType:   res.BoxType,
		Source:    thetanboxenum.BoxEventSource(res.Source),
		Timestamp: res.Timestamp,
		BoxId:     res.BoxId,
	}, true
}
