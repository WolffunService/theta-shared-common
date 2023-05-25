package rivalitemenum

import (
	"errors"
	"fmt"

	"github.com/WolffunService/thetan-shared-common/enums/thetanrivalerrorenum/thetanrivalconst"
)

type ItemRarity int

const (
	IRNone   ItemRarity = iota
	IRNormal            // trong khi chờ update Rarity thì rarity sẽ 1->3
	IRRare
	IRMythical
)

var (
	ErrUndefinedItemRarity = errors.New("undefined item")
	ErrUndefinedType       = errors.New("undefined type")
)
var _textItemRare = map[ItemRarity]string{
	IRNone:     "unk",
	IRNormal:   "normal",
	IRRare:     "rare",
	IRMythical: "mythical",
}

func (i ItemRarity) String() string {
	return _textItemRare[i]
}

func getItemRare(itemRare string) ItemRarity {
	ir := IRNone
	for rarity, rareStr := range _textItemRare {
		if rareStr == itemRare {
			ir = rarity
			break
		}
	}
	return ir
}

func ToItemRarity(itemType any) (ItemRarity, error) {
	switch v := itemType.(type) {
	case thetanrivalconst.CosmeticRarity:
		return getRarityByCosmetic(v)
	case thetanrivalconst.MinionSkinRarity:
		return getRarityByMinionSkin(v)
	case string:
		return getItemRare(v), nil
	default:
		return IRNone, fmt.Errorf("type %v : %w", itemType, ErrUndefinedType)
	}
}

func getRarityByCosmetic(cr thetanrivalconst.CosmeticRarity) (ItemRarity, error) {
	switch cr {
	case thetanrivalconst.CosmeticRarityNormal:
		return IRNormal, nil
	case thetanrivalconst.CosmeticRarityRare:
		return IRRare, nil
	case thetanrivalconst.CosmeticRarityMythical:
		return IRMythical, nil
	}
	return IRNone, fmt.Errorf("type %v : %w", cr, ErrUndefinedItemRarity)
}

func getRarityByMinionSkin(cr thetanrivalconst.MinionSkinRarity) (ItemRarity, error) {
	switch cr {
	case thetanrivalconst.MinionSkinRarityNormal:
		return IRNormal, nil
	case thetanrivalconst.MinionSkinRarityRare:
		return IRRare, nil
	case thetanrivalconst.MinionSkinRarityMythical:
		return IRMythical, nil
	}
	return IRNone, fmt.Errorf("type %v : %w", cr, ErrUndefinedItemRarity)
}
