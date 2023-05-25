package esmodel

import "time"

type HeroRentalMapping struct {
	User            UserModel `json:"user"`
	OwnerAddress    string    `json:"owner_address"`
	OwnerId         string    `json:"owner_id"`
	OnMarketTime    time.Time `json:"on_market_time"`
	MaxRentalPeriod int       `json:"max_rental_period"`
	ThcRentalPrice  int       `json:"thc_rental_price"`
	ThcFee          int       `json:"thc_fee"`
	RentBattles     int       `json:"rent_battles"`
	HeroId          string    `json:"hero_id"`
	HeroTypeId      int       `json:"hero_type"`
	SkinId          int       `json:"skin"`
	HeroRarity      int       `json:"hero_rarity"`
	SkinRarity      int       `json:"skin_rarity"`
	HeroLevel       int       `json:"hero_level"`
	TrophyClass     int       `json:"trophy_class"`
	ThcBattle       int       `json:"thc_battle"`
	ThcBattleLimit  int       `json:"thc_battle_limit"`
	ExpiredTime     time.Time `json:"expired_time"`
	Timestamp       time.Time `json:"@timestamp"`
}

func (HeroRentalMapping) Index() string {
	return "playerstats-herorental"
}
