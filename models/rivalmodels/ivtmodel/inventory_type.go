package ivtmodel

import "github.com/WolffunService/theta-shared-common/enums/rivalitemenum"

// InventoryType id of inventory's, must unique in same InventoryKind
type InventoryType uint32

const (
	// CurrencyTHC Thetan Coin
	CurrencyTHC InventoryType = 11

	// CurrencyTHG Thetan Gem
	CurrencyTHG InventoryType = 12

	// CurrencyRBuck Rivals Buck
	CurrencyRBuck InventoryType = 21

	// CurrencyGold Rivals Gold
	CurrencyGold InventoryType = 22

	// CurrencySSPoint Season Point
	CurrencySSPoint InventoryType = 23

	// CurrencySSBooster Season Booster
	CurrencySSBooster InventoryType = 24

	// CurrencyEHC Enhancer
	CurrencyEHC InventoryType = 25

	// CurrencyEXP Season experience
	CurrencyEXP InventoryType = 26

	CurrencyUSD InventoryType = 27

	CurrencyTrophy InventoryType = 28
)

var _textCurrency = map[InventoryType]string{
	CurrencyRBuck: "rival_buck",

	CurrencyGold: "gold",

	CurrencySSPoint: "season_point",

	CurrencySSBooster: "season_boost",

	CurrencyEHC: "enhancer",

	CurrencyEXP: "exp",

	CurrencyUSD: "usd",

	CurrencyTrophy: "trophy",
}

func (i InventoryType) String() string {
	return _textCurrency[i]
}

func (i InventoryType) ToInt() int {
	return int(i)
}

func (i InventoryType) SeasonalCurrency() bool {
	return i == CurrencyEXP || i == CurrencySSPoint || i == CurrencySSBooster
}

func ItemTypeToInventoryType(itemType rivalitemenum.ItemType) InventoryType {
	switch itemType {
	case rivalitemenum.ITThetanCoin:
		return CurrencyTHC
	case rivalitemenum.ITThetanGem:
		return CurrencyTHG
	case rivalitemenum.ITEnhancer:
		return CurrencyEHC
	case rivalitemenum.ITRivalBuck:
		return CurrencyRBuck
	case rivalitemenum.ITGold:
		return CurrencyGold
	case rivalitemenum.ITSeasonPoint:
		return CurrencySSPoint
	case rivalitemenum.ITSeasonBooster:
		return CurrencySSBooster
	}
	return 0
}
