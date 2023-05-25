package thetanrivalconst

import "fmt"

type MinionRarity uint8

const (
	MinionRarityCommon MinionRarity = 1 + iota
	MinionRarityEpic
	MinionRarityLegendary
)

func (r MinionRarity) String() string {
	switch r {
	case MinionRarityCommon:
		return "Common"
	case MinionRarityEpic:
		return "Epic"
	case MinionRarityLegendary:
		return "Legendary"
	}

	return fmt.Sprintf("Unknown (%d)", r)
}

type MinionSkinRarity uint8

const (
	MinionSkinRarityNormal MinionSkinRarity = 1 + iota
	MinionSkinRarityRare
	MinionSkinRarityMythical
)

func (r MinionSkinRarity) String() string {
	switch r {
	case MinionSkinRarityNormal:
		return "Normal"
	case MinionSkinRarityRare:
		return "Rare"
	case MinionSkinRarityMythical:
		return "Mythical"
	}

	return fmt.Sprintf("Unknown (%d)", r)
}

type CosmeticRarity uint8

const (
	CosmeticRarityNormal CosmeticRarity = 1 + iota
	CosmeticRarityRare
	CosmeticRarityMythical
)

func (r CosmeticRarity) String() string {
	switch r {
	case CosmeticRarityNormal:
		return "Normal"
	case CosmeticRarityRare:
		return "Rare"
	case CosmeticRarityMythical:
		return "Mythical"
	}

	return fmt.Sprintf("Unknown (%d)", r)
}

type BattleRank int

const (
	BattleRank4 BattleRank = 4
	BattleRank3 BattleRank = 3
	BattleRank2 BattleRank = 2
	BattleRank1 BattleRank = 1
)

type RewardType int

const (
	RewardTypeTrophy RewardType = 1 + iota
	RewardTypeSeasonPoint
	RewardTypeEXP
)
