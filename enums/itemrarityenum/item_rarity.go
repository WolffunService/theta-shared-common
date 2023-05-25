package itemrarityenum

import (
	"errors"
	"fmt"

	"github.com/WolffunService/thetan-shared-common/enums/heroenum"
	"github.com/WolffunService/thetan-shared-common/enums/thetanboxenum"
)

var (
	ErrUndefinedItemRarity = errors.New("undefined item")
	ErrUndefinedType       = errors.New("undefined type")
)

type ItemRarity int // #enum: "Item Rarity(General rarity)"

func (i ItemRarity) IsValid() bool {
	return i > start && i < end
}

const (
	start ItemRarity = iota - 1
	COMMON
	EPIC
	LEGENDARY

	end
)

var text = map[ItemRarity]string{
	COMMON:    "COMMON",
	EPIC:      "EPIC",
	LEGENDARY: "LEGENDARY",
}

func (i ItemRarity) String() string {
	return text[i]
}

func (i ItemRarity) ToHeroRarity() (heroenum.HeroRarity, error) {
	switch i {
	case COMMON:
		return heroenum.COMMON, nil
	case EPIC:
		return heroenum.EPIC, nil
	case LEGENDARY:
		return heroenum.LEGENDARY, nil
	}
	return -1, ErrUndefinedItemRarity
}

func (i ItemRarity) ToBoxRarity() (thetanboxenum.BoxType, error) {
	switch i {
	case COMMON:
		return thetanboxenum.COMMON_BOX, nil
	case EPIC:
		return thetanboxenum.EPIC_BOX, nil
	case LEGENDARY:
		return thetanboxenum.LEGENDARY_BOX, nil
	}
	return -1, ErrUndefinedItemRarity
}

func GetItemRarity(allRarity any) (ItemRarity, error) {
	switch v := allRarity.(type) {
	case heroenum.HeroRarity:
		return getItemRarityByHeroRarity(v)
	case thetanboxenum.BoxType:
		return getItemRarityByBoxRarity(v)
	default:
		return start, errors.New(fmt.Sprintf("undefined item %T %v", allRarity, allRarity))
	}
}

func getItemRarityByHeroRarity(rarity heroenum.HeroRarity) (ItemRarity, error) {
	switch rarity {
	case heroenum.COMMON:
		return COMMON, nil
	case heroenum.EPIC:
		return EPIC, nil
	case heroenum.LEGENDARY:
		return LEGENDARY, nil
	}
	return -1, ErrUndefinedItemRarity
}

func getItemRarityByBoxRarity(rarity thetanboxenum.BoxType) (ItemRarity, error) {
	switch rarity {
	case thetanboxenum.COMMON_BOX:
		return COMMON, nil
	case thetanboxenum.EPIC_BOX:
		return EPIC, nil
	case thetanboxenum.LEGENDARY_BOX:
		return LEGENDARY, nil
	}
	return -1, ErrUndefinedItemRarity
}
