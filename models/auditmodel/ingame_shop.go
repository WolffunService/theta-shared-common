package auditmodel

import (
	"log"
	"strconv"

	"github.com/WolffunService/theta-shared-common/enums/currencyenum"
	"github.com/WolffunService/theta-shared-common/enums/ingameshopenum"
	"github.com/WolffunService/theta-shared-common/proto/auditproto"
)

type IngameShopAudits struct {
	UserId              string                          `json:"userId" bson:"userId"`
	PackId              int                             `json:"packId" bson:"packId"`
	Source              ingameshopenum.IngameShopSource `json:"source" bson:"source"`
	Price               float64                         `json:"price" bson:"price"`
	PriceType           currencyenum.Currency           `json:"priceType" bson:"priceType"`
	IngameShopPP        *ingameShopPP                   `json:"shopPP,omitempty" bson:"shopPP,omitempty"`
	IngameShopBoxNonNft *ingameShopBoxNonNft            `json:"shopBoxNonNft,omitempty" bson:"shopBoxNonNft,omitempty"`
	Timestamp           int64                           `json:"timestamp" bson:"timestamp"`
}

func (i *IngameShopAudits) FromSimpleEvent(a *auditproto.SimpleEvent) IngameShopAudits {
	var err error
	source, err := strconv.Atoi(a.Metadata["source"])
	if err != nil {
		log.Println("[error][analytic][topic-0][ingame-shop] cannot parse metadata[source]")
	}

	packId, err := strconv.Atoi(a.Metadata["packId"])
	if err != nil {
		log.Println("[error][analytic][topic-0][ingame-shop] cannot parse metadata[packId]")
	}

	price, err := strconv.ParseFloat(a.Metadata["price"], 64)
	if err != nil {
		log.Println("[error][analytic][topic-0][ingame-shop] cannot parse metadata[price]")
	}

	priceType, err := strconv.Atoi(a.Metadata["priceType"])
	if err != nil {
		log.Println("[error][analytic][topic-0][ingame-shop] cannot parse metadata[priceType]")
	}

	var shopPP *ingameShopPP = nil
	var shopBoxNonNft *ingameShopBoxNonNft = nil
	if source == int(ingameshopenum.ISS_BuyBoxNonNft) {
		skinId, err := strconv.Atoi(a.Metadata["skinId"])
		if err != nil {
			log.Println("[error][analytic][topic-0][ingame-shop] cannot parse metadata[skinId]")
		}
		shopBoxNonNft = &ingameShopBoxNonNft{
			SkinId: skinId,
			HeroId: a.Metadata["heroId"],
		}
	} else if source == int(ingameshopenum.ISS_BuyPowerPoint) {
		powerPoint, err := strconv.Atoi(a.Metadata["powerPoint"])
		if err != nil {
			log.Println("[error][analytic][topic-0][ingame-shop] cannot parse metadata[powerPoint]")
		}
		shopPP = &ingameShopPP{
			PowerPoint: int64(powerPoint),
		}
	}

	return IngameShopAudits{
		UserId:              a.Metadata["userId"],
		PackId:              packId,
		Source:              ingameshopenum.IngameShopSource(source),
		Price:               price,
		PriceType:           currencyenum.Currency(priceType),
		IngameShopPP:        shopPP,
		IngameShopBoxNonNft: shopBoxNonNft,
		Timestamp:           a.Event.Timestamp,
	}
}

type ingameShopPP struct {
	PowerPoint int64 `json:"powerPoint"`
}

type ingameShopBoxNonNft struct {
	HeroId string `json:"heroId"`
	SkinId int    `json:"skinId"`
}
