package boxrewardenum

import (
	"github.com/WolffunService/thetan-shared-common/enums/heroenum"
	"strings"
)

type BoxRewardType uint32

const (
	Hero BoxRewardType = 1 << iota
	Cosmetic
	Currency

	Common
	Epic
	Legendary

	Normal
	Rare
	Mythical

	GameTHC
	GameTHG
	GamePP

	// ALTERNATIVE //

	HeroCommon    = Hero | Common
	HeroEpic      = Hero | Epic
	HeroLegendary = Hero | Legendary

	HeroCommonNormal   = Hero | Common | Normal
	HeroCommonRare     = HeroCommon | Rare
	HeroCommonMythical = HeroCommon | Mythical

	HeroEpicNormal   = HeroEpic | Normal
	HeroEpicRare     = HeroEpic | Rare
	HeroEpicMythical = HeroEpic | Mythical

	HeroLegendaryNormal   = HeroLegendary | Normal
	HeroLegendaryRare     = HeroLegendary | Rare
	HeroLegendaryMythical = HeroLegendary | Mythical

	CosmeticCommon    = Cosmetic | Common
	CosmeticEpic      = Cosmetic | Epic
	CosmeticLegendary = Cosmetic | Legendary

	CurrencyGTHC = Currency | GameTHC
	CurrencyGTHG = Currency | GameTHG
	CurrencyGPP  = Currency | GamePP
)

var listTypes []BoxRewardType
var names map[BoxRewardType]string
var heroRarityMapping map[BoxRewardType]heroenum.HeroRarity

func init() {
	listTypes = append(listTypes,
		Hero,
		Cosmetic,
		Currency,

		Common,
		Epic,
		Legendary,

		Normal,
		Rare,
		Mythical,

		GameTHC,
		GameTHG,
		GamePP,
	)

	names = map[BoxRewardType]string{
		Hero:     "Hero",
		Cosmetic: "Cosmetic",
		Currency: "Currency",

		Common:    "Common",
		Epic:      "Epic",
		Legendary: "Legendary",

		Normal:   "Normal",
		Rare:     "Rare",
		Mythical: "Mythical",

		GameTHC: "GameTHC",
		GameTHG: "GameTHG",
		GamePP:  "GamePP",
	}

	heroRarityMapping = map[BoxRewardType]heroenum.HeroRarity{
		Common:    heroenum.COMMON,
		Epic:      heroenum.EPIC,
		Legendary: heroenum.LEGENDARY,
	}
}

func (rt BoxRewardType) Has(another BoxRewardType) bool {
	return rt&another != 0
}

func (rt BoxRewardType) String() string {
	if n, found := names[rt]; found {
		return n
	}

	var fullName []string
	for _, rewardType := range listTypes {
		if rt.Has(rewardType) {
			fullName = append(fullName, names[rewardType])
		}
	}

	return strings.Join(fullName, " ")
}

func (rt BoxRewardType) GetHeroRarity() heroenum.HeroRarity {
	for rewardType, rarity := range heroRarityMapping {
		if rt.Has(Hero) && rt.Has(rewardType) {
			return rarity
		}
	}

	// Return default rarity
	return heroenum.UNKNOWN
}
