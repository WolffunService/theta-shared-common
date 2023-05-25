package esmodel

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/enums/currencyenum"
	"github.com/WolffunService/thetan-shared-common/enums/ingameshopenum"
)

type IAPStatMapping struct {
	User       UserModel                       `json:"user"`
	PackID     int                             `json:"pack_id"` //enum
	Source     ingameshopenum.IngameShopSource `json:"source"`  //enum
	Price      float64                         `json:"price"`   //
	PriceType  currencyenum.Currency           `json:"price_type"`
	PowerPoint int                             `json:"power_point"`
	HeroID     string                          `json:"hero_id"` //object id, ko trùng
	SkinID     int                             `json:"skin_id"` //enum, có trùng
	Timestamp  time.Time                       `json:"@timestamp"`
}

func (IAPStatMapping) Index() string {
	return "playerstats-iap"
}
