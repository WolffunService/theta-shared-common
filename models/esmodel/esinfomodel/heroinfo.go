package esinfomodel

import "time"

type HeroInfoModel struct {
	HeroName         string    `json:"hero_name"`
	SkinName         string    `json:"skin_name"`
	HeroTypeId       int       `json:"hero_type"`
	SkinId           int       `json:"skin"`
	HeroRarity       int       `json:"hero_rarity"`
	SkinRarity       int       `json:"skin_rarity"`
	HeroLevel        int       `json:"hero_level"`
	TrophyClass      int       `json:"trophy_class"`
	ThcBattle        int       `json:"thc_battle"`
	ThcBattleLimit   int       `json:"thc_battle_limit"`
	ThcBattlePercent int       `json:"thc_battle_percent"`
	HeroRole         int       `json:"hero_role"`
	Status           int       `json:"status"`
	Hp               int       `json:"hp"`
	Damage           int       `json:"damage"`
	AttackSpeed      int64     `json:"attack_speed"`
	Speed            int64     `json:"speed"`
	FuryRequired     int       `json:"fury_required"`
	Created          time.Time `json:"created"`
}
