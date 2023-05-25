package esinfomodel

import "time"

type CosmeticInfoModel struct {
	CosmeticName   string    `json:"cosmetic_name"`
	CosmeticType   int       `json:"cosmetic_type"`
	CosmeticTypeId int       `json:"cosmetic_type_id"`
	CosmeticRarity int       `json:"cosmetic_rarity"`
	Status         int       `json:"status"`
	Created        time.Time `json:"created"`
}
