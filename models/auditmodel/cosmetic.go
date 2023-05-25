package auditmodel

import (
	"log"
	"strconv"

	"github.com/WolffunService/theta-shared-common/enums/cosmeticenum"
	"github.com/WolffunService/theta-shared-common/enums/currencyenum"
	"github.com/WolffunService/theta-shared-common/models/currencymodel"
	"github.com/WolffunService/theta-shared-common/proto/auditproto"
)

type CosmeticAudits struct {
	ID          interface{}                      `json:"id" bson:"_id,omitempty"`
	UserId      string                           `json:"userId" bson:"userId"`
	UserAddress string                           `json:"userAddress" bson:"userAddress"`
	TypeID      int                              `json:"typeId" bson:"typeId"`
	Source      cosmeticenum.CosmeticEventSource `json:"source" bson:"source"`
	Type        cosmeticenum.CosmeticType        `json:"type" bson:"type"`
	Timestamp   int64                            `json:"timestamp" bson:"timestamp"`
	ItemID      string                           `json:"itemId" bson:"itemId"`
	TradeInfo   *TradeInfo                       `json:"tradeInfo,omitempty" bson:"tradeInfo,omitempty"`
}

const (
	_USER_ADDRESS     = "userAddress"
	_USER_ID          = "userId"
	_SELLER_ADDRESS   = "sellerAddress"
	_SELLER_ID        = "sellerId"
	_DECIMALS         = "decimals"
	_CURRENCY_TYPE    = "currency"
	_PRICE_VALUE      = "price"
	_FEE              = "realFee"
	_TRANSACTION_HASH = "transactionHash"
	_ITEM_TYPE        = "itemType"
	_ITEM_TYPE_ID     = "typeId"
	_ITEM_ID          = "itemId"
	_SOURCE           = "source"
)

type decouple struct {
	s1Addr string
	s1ID   string
	s2Addr string
	s2ID   string
	source cosmeticenum.CosmeticEventSource
}

func (b *CosmeticAudits) ToTradingCosmeticAudits(e *auditproto.SimpleEvent) []interface{} {
	auditMap := map[string]string{}
	if e.Event != nil {
		for _, v := range e.Event.EventParams {
			auditMap[v.Key] = v.Value
		}
	}

	itemType, err := strconv.Atoi(auditMap[_ITEM_TYPE])
	if err != nil {
		log.Println("[error][trading-item] cannot get itemType", err)
	}
	itemTypeId, err := strconv.Atoi(auditMap[_ITEM_TYPE_ID])
	if err != nil {
		log.Println("[error][trading-item] cannot get itemTypeId", err)
	}
	decimals, err := strconv.Atoi(auditMap[_DECIMALS])
	if err != nil {
		log.Println("[error][[trading-item] cannot get decimal", err)
	}
	currencyType, err := strconv.Atoi(auditMap[_CURRENCY_TYPE])
	if err != nil {
		log.Println("[error][[trading-item] cannot get currency type", err)
	}
	priceValue, err := strconv.Atoi(auditMap[_PRICE_VALUE])
	if err != nil {
		log.Println("[error][[trading-item] cannot get priceValue", err)
	}
	fee, err := strconv.ParseFloat(auditMap[_FEE], 64)
	if err != nil {
		log.Println("[error][[trading-item] cannot get fee", err)
	}

	factors := []decouple{
		{
			s1Addr: auditMap[_SELLER_ADDRESS],
			s1ID:   auditMap[_SELLER_ID],
			s2Addr: auditMap[_USER_ADDRESS],
			s2ID:   auditMap[_USER_ID],
			source: cosmeticenum.CES_SOLD,
		},
		{
			s1Addr: auditMap[_USER_ADDRESS],
			s1ID:   auditMap[_USER_ID],
			s2Addr: auditMap[_SELLER_ADDRESS],
			s2ID:   auditMap[_SELLER_ID],
			source: cosmeticenum.CES_PURCHASED,
		},
	}

	var tradingItems []interface{}
	for _, v := range factors {
		tradingItems = append(tradingItems, CosmeticAudits{
			UserId: v.s1ID,
			// UserAddress: v.s1Addr,
			Type:   cosmeticenum.CosmeticType(itemType),
			TypeID: itemTypeId,
			Source: v.source,
			ItemID: auditMap[_ITEM_ID],
			TradeInfo: &TradeInfo{
				OtherUserId:   v.s2ID,
				OtherAddress:  v.s2Addr,
				TransactionId: auditMap[_TRANSACTION_HASH],
				Price: currencymodel.SystemCurrency{
					Type:     currencyenum.Currency(currencyType),
					Value:    int64(priceValue),
					Decimals: decimals,
				},
				Fee: float32(fee),
			},
			Timestamp: e.Event.Timestamp,
		})
	}

	return tradingItems
}

func (b *CosmeticAudits) ToCosmeticEventAudits(e *auditproto.SimpleEvent) *CosmeticAudits {
	itemType, err := strconv.Atoi(e.Metadata[_ITEM_TYPE])
	if err != nil {
		log.Println("[error][trading-item] cannot get itemType", err)
	}
	itemTypeId, err := strconv.Atoi(e.Metadata[_ITEM_TYPE_ID])
	if err != nil {
		log.Println("[error][trading-item] cannot get itemTypeId", err)
	}
	source, err := strconv.Atoi(e.Metadata[_SOURCE])
	if err != nil {
		log.Println("[error][[trading-item] cannot get decimal", err)
	}

	return &CosmeticAudits{
		UserId:    e.Metadata[_USER_ID],
		Type:      cosmeticenum.CosmeticType(itemType),
		TypeID:    itemTypeId,
		Source:    cosmeticenum.CosmeticEventSource(source),
		Timestamp: e.Event.Timestamp,
		ItemID:    e.Metadata[_ITEM_ID],
	}
}

func (b *CosmeticAudits) ToOpenBoxCosmeticAudits(e *auditproto.SimpleEvent) *CosmeticAudits {
	auditMap := map[string]string{}
	if e.Event != nil {
		for _, v := range e.Event.EventParams {
			auditMap[v.Key] = v.Value
		}
	}

	itemType, err := strconv.Atoi(auditMap[_ITEM_TYPE])
	if err != nil {
		log.Println("[error][open-box] cannot get itemType", err)
	}
	typeId, err := strconv.Atoi(auditMap[_ITEM_TYPE_ID])
	if err != nil {
		log.Println("[error][open-box] cannot get itemTypeId", err)
	}

	return &CosmeticAudits{
		UserId:    auditMap[_USER_ID],
		Type:      cosmeticenum.CosmeticType(itemType),
		TypeID:    typeId,
		Source:    cosmeticenum.CES_OPEN_BOX,
		Timestamp: e.Event.Timestamp,
		ItemID:    auditMap[_ITEM_ID],
	}
}

// DO NOT USING CONSTANT
func (b *CosmeticAudits) ToReceivedCosmeticAudits(e *auditproto.SimpleEvent) *CosmeticAudits {
	auditMap := map[string]string{}
	if e.Event != nil {
		for _, v := range e.Event.EventParams {
			auditMap[v.Key] = v.Value
		}
	}

	itemTypeID, err := strconv.Atoi(auditMap[_ITEM_TYPE_ID]) // item type id
	if err != nil {
		log.Println("[error][item-received] cannot get itemTypeId", err)
	}

	source, err := strconv.Atoi(auditMap[_SOURCE])
	if err != nil {
		log.Println("[error][[item-received] cannot get source", err)
	}

	itemType, err := strconv.Atoi(auditMap[_ITEM_TYPE])
	if err != nil {
		log.Println("[error][item-received] cannot get itemType", err)
	}

	return &CosmeticAudits{
		UserId:    auditMap[_USER_ID],
		TypeID:    itemTypeID,
		ItemID:    auditMap[_ITEM_ID],
		Source:    cosmeticenum.CosmeticEventSource(source),
		Type:      cosmeticenum.CosmeticType(itemType),
		Timestamp: e.Event.Timestamp,
	}
}
