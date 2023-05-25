package auditmodel

import (
	"fmt"
	"log"
	"strconv"

	"github.com/WolffunService/theta-shared-common/enums/currencyenum"
	"github.com/WolffunService/theta-shared-common/enums/heroenum"
	"github.com/WolffunService/theta-shared-common/models/currencymodel"
	"github.com/WolffunService/theta-shared-common/proto/auditproto"
)

type HeroAuditCreator struct{}

type HeroEvent struct {
	ID        interface{}              `json:"id" bson:"_id,omitempty"`
	UserId    string                   `json:"userId" bson:"userId"`
	Timestamp int64                    `json:"timestamp" bson:"timestamp"`
	HeroId    string                   `json:"heroId" bson:"heroId"`
	SkinId    int32                    `json:"skinId" bson:"skinId"`
	Source    heroenum.HeroEventSource `json:"source" bson:"source"`
	Status    heroenum.HeroEventStatus `json:"status" bson:"status"`
	TradeInfo *TradeInfo               `json:"tradeInfo,omitempty" bson:"tradeInfo,omitempty"`
}

type HeroRentalEvent struct {
	ID            interface{}        `json:"id"`
	LastModified  int64              `json:"lastModified"`
	RentBattles   int                `json:"rentBattles"`
	RenterAddress string             `json:"renterAddress"`
	SkinId        int32              `json:"skinId"`
	State         heroenum.RentState `json:"state"`
	TokenName     string             `json:"tokenName"`
	TokenId       int64              `json:"tokenId"`
	RefId         string             `json:"refId"`
}

type TradeInfo struct {
	OtherAddress  string                       `json:"address,omitempty" bson:"address,omitempty"`
	OtherUserId   string                       `json:"-" bson:"otherUserId,omitempty"`
	Price         currencymodel.SystemCurrency `json:"price" bson:"price"`
	TransactionId string                       `json:"-" bson:"transactionId,omitempty"`
	Fee           float32                      `json:"-" bson:"fee,omitempty"`
}

type UpgradeHeroEvent struct {
	UserId       string `json:"userId" bson:"userId"`
	HeroLevel    int    `json:"heroLevel" bson:"heroLevel"`
	HeroId       string `json:"heroId" bson:"heroId"`
	BattleNumber int    `json:"battleNumber" bson:"battleNumber"`
	Timestamp    int64  `json:"timestamp" bson:"timestamp"`
}

// NewOpenEvent return BoxEvent after convert and a bool determine that you need write event to database
func (HeroAuditCreator) NewCreationEvent(res *auditproto.HeroCreation) *HeroEvent {
	return &HeroEvent{
		UserId:    res.UserId,
		Timestamp: res.Timestamp,
		HeroId:    res.HeroId,
		SkinId:    res.SkinId,
		Source:    heroenum.HeroEventSource(res.Source),
		Status:    heroenum.HESt_Succeeded,
	}
}

func (HeroAuditCreator) NewCreationEventFromSimpleEvent(e *auditproto.SimpleEvent) *HeroEvent {
	auditMap := map[string]string{}
	if e.Event != nil {
		for _, v := range e.Event.EventParams {
			auditMap[v.Key] = v.Value
		}
	}

	itemId, err := strconv.Atoi(auditMap["itemId"])
	if err != nil {
		log.Println("[error][audit] cannot parse itemId:", itemId, err)
	}

	source, err := strconv.Atoi(auditMap["source"])
	if err != nil {
		log.Println("[error][audit] cannot parse source:", source, err)
	}

	if source == 11 {
		fmt.Println(e.Metadata)
	}

	var sc *currencymodel.SystemCurrency = nil
	if priceValue, err := strconv.Atoi(e.Metadata["priceValue"]); err == nil {
		priceType, err := strconv.Atoi(e.Metadata["priceType"])
		if err != nil {
			log.Println("[error][audit] cannot parse priceType:", err)
		}

		priceDecimals, err := strconv.Atoi(e.Metadata["priceDecimals"])
		if err != nil {
			log.Println("[error][audit] cannot parse priceDecimals:", err)
		}

		sc = &currencymodel.SystemCurrency{
			Value:    int64(priceValue),
			Type:     currencyenum.Currency(priceType),
			Decimals: priceDecimals,
		}
	}

	var tradeInfo *TradeInfo = nil
	if sc != nil {
		tradeInfo = &TradeInfo{
			Price: *sc,
		}
	}

	return &HeroEvent{
		UserId:    auditMap["userId"],
		HeroId:    auditMap["id"],
		SkinId:    int32(itemId),
		Source:    heroenum.HeroEventSource(source),
		Status:    heroenum.HESt_Succeeded,
		Timestamp: e.Event.Timestamp,
		TradeInfo: tradeInfo,
	}
}

func (HeroAuditCreator) NewTradingEvent_Purchased(res *auditproto.TransactionSuccess) *HeroEvent {
	return &HeroEvent{
		UserId:    res.UserId,
		Timestamp: res.Timestamp,
		HeroId:    res.HeroId,
		SkinId:    res.SkinId,
		Source:    heroenum.HES_Purchase,
		Status:    heroenum.HESt_Succeeded,
		TradeInfo: &TradeInfo{
			TransactionId: res.TransactionId,
			Fee:           res.RealFee,
			OtherUserId:   res.SellerId,
			OtherAddress:  res.SellerAddress,
			Price: currencymodel.SystemCurrency{
				Value:    res.PurchasePrice,
				Type:     currencyenum.Currency(res.PurchaseCurrency),
				Decimals: int(res.Decimals),
			},
		},
	}
}

func (HeroAuditCreator) NewTradingEvent_Sold(res *auditproto.TransactionSuccess) *HeroEvent {
	return &HeroEvent{
		UserId:    res.SellerId,
		Timestamp: res.Timestamp,
		HeroId:    res.HeroId,
		SkinId:    res.SkinId,
		Source:    heroenum.HES_Sold,
		Status:    heroenum.HESt_Succeeded,
		TradeInfo: &TradeInfo{
			TransactionId: res.TransactionId,
			OtherUserId:   res.UserId,
			OtherAddress:  res.UserAddress,
			Fee:           res.RealFee,
			Price: currencymodel.SystemCurrency{
				Type:     currencyenum.Currency(res.PurchaseCurrency),
				Value:    res.PurchasePrice,
				Decimals: int(res.Decimals),
			},
		},
	}
}

func (HeroAuditCreator) NewUpgradeHeroEvent(res *auditproto.UpgradeHeroResult) *UpgradeHeroEvent {
	return &UpgradeHeroEvent{
		UserId:       res.UserId,
		HeroId:       res.HeroId,
		HeroLevel:    int(res.HeroLevel),
		BattleNumber: int(res.BattleNumber),
		Timestamp:    res.Timestamp,
	}
}
