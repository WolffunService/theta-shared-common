package rivalitem

import (
	"fmt"
	"strconv"

	"github.com/WolffunService/thetan-shared-common/enums/rivalitemenum"
	"github.com/WolffunService/thetan-shared-common/models/rivalmodels/ivtmodel"
	"github.com/WolffunService/thetan-shared-common/models/utilmodel"
)

type SimpleItemG[T any] struct {
	SimplifyItem `json:",inline" yaml:",inline" mapstructure:",squash" bson:",inline"`
	Props        T `json:"props" yaml:"props" mapstructure:"props" bson:"props"`
}

type SimpleItem SimpleItemG[map[string]any]

func (i SimpleItem) Rarity() rivalitemenum.ItemRarity {
	rarityRaw, ok := i.Props["rarity"]
	if !ok {
		return rivalitemenum.IRNormal
	}

	rarity, err := strconv.Atoi(fmt.Sprintf("%d", rarityRaw))
	if err != nil {
		return rivalitemenum.IRNormal
	}

	return rivalitemenum.ItemRarity(rarity)
}

type SimplifyItem struct {
	ItemType      rivalitemenum.ItemType `json:"itemType" yaml:"itemType" mapstructure:"itemType" bson:"itemType"`
	TypeID        int                    `json:"typeId" yaml:"typeId" mapstructure:"typeId" bson:"typeId"`
	Amount        utilmodel.Number       `json:"amount" yaml:"-" mapstructure:"amount" bson:"amount"`
	InventoryKind ivtmodel.InventoryKind `json:"inventoryKind" yaml:"-" mapstructure:"inventoryKind" bson:"inventoryKind"`
}

func (s SimplifyItem) ToSimple() SimpleItem {
	return SimpleItem{SimplifyItem: s, Props: make(map[string]any)}
}

// func (i SimpleItem) Add(j SimpleItem) (SimpleItem, error) {
// 	if i.InventoryKind != j.InventoryKind {
// 		return SimpleItem{}, fmt.Errorf("inventory kind not match")
// 	}

// 	if i.TypeID != i.TypeID {
// 		return SimpleItem{}, fmt.Errorf("inventory kind not match")
// 	}

// 	if i.ItemType != i.ItemType {
// 		return SimpleItem{}, fmt.Errorf("inventory kind not match")
// 	}

// 	k := i
// 	k.Amount.Add(j.Amount)
// 	return k, nil
// }
