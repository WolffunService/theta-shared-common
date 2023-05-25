package ivtmodel

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/WolffunService/theta-shared-common/enums/inventoryenum"
	"github.com/WolffunService/theta-shared-common/enums/rivalboxenum"
	"github.com/WolffunService/theta-shared-common/enums/rivalitemenum"
	"github.com/WolffunService/theta-shared-common/enums/thetanrivalerrorenum/thetanrivalconst"
)

var ErrInventoryInvalidFormat = errors.New("inventory invalid format")

const InventorySeparator = "_"

// Inventory data models
type Inventory struct {
	Name   string                         `json:"name,omitempty"`
	Kind   InventoryKind                  `json:"kind"`
	Type   InventoryType                  `json:"type"`
	Desc   string                         `json:"desc,omitempty"`
	Tags   []rivalitemenum.Tag            `json:"tags,omitempty"`
	Extras map[inventoryenum.PropName]any `json:"extras,omitempty"`
}

type InventoryFull struct {
	Inventory `json:",inline" bson:",inline"`
	Props     InventoryProperties `json:"props" bson:"props"`
}

func ExtractInventory(inventoryKey string) (InventoryKind, InventoryType, error) {
	split := strings.Split(inventoryKey, InventorySeparator)
	if len(split) != 2 {
		return IKUndefined, 0, ErrInventoryInvalidFormat
	}

	kind, err := strconv.Atoi(split[0])
	if err != nil {
		return IKUndefined, 0, nil
	}

	typ, err := strconv.Atoi(split[1])
	if err != nil {
		return IKUndefined, 0, nil
	}

	return InventoryKind(kind), InventoryType(typ), nil
}

func (i Inventory) UniqueId() string {
	return UniqueId(i.Kind, i.Type)
}

func UniqueId(kind InventoryKind, typ InventoryType) string {
	return fmt.Sprintf("%d%s%d", kind, InventorySeparator, typ)
}

func (i Inventory) InExclusiveFor(expected string) bool {
	// Validate exclusive is nil pointer
	if i.Extras == nil || i.Extras[inventoryenum.ExclusiveFor] == nil {
		return false
	}

	// Check has expected exclusive in slice
	return i.Extras[inventoryenum.ExclusiveFor] != nil && i.Extras[inventoryenum.ExclusiveFor].(string) == expected
}

// Deprecated: GetCosmeticRarity is deprecated. Please use GetItemRarity instead
func (i Inventory) GetCosmeticRarity() thetanrivalconst.CosmeticRarity {
	r := i.GetItemRarity()
	//if r == rivalitemenum.IRNormal {
	//	return thetanrivalconst.CosmeticRarityNormal
	//}

	return thetanrivalconst.CosmeticRarity(r)
}

// GetItemRarity Get rarity of item, if not config return rarity None
func (i Inventory) GetItemRarity() rivalitemenum.ItemRarity {
	if r, found := i.Extras[inventoryenum.Rarity]; found {
		rarity := r.(int)
		return rivalitemenum.ItemRarity(rarity)
	}

	return rivalitemenum.IRNormal
}

func (i Inventory) GetDropGroup() rivalitemenum.DropGroup {
	switch i.Kind {
	case IKMinion:
		return rivalitemenum.DGMinion
	case IKCosmeticProfile:
		return rivalitemenum.DGProfile
	case IKCosmeticAddIn:
		switch i.GetItemType() {
		case rivalitemenum.ITGlow, rivalitemenum.ITVoice, rivalitemenum.ITDance, rivalitemenum.ITVehicle:
			return rivalitemenum.DGFullEvolve
		case rivalitemenum.ITFootprint, rivalitemenum.ITBackBling, rivalitemenum.ITFlyCraft:
			return rivalitemenum.DGHalfEvolve
		}
	}

	return rivalitemenum.DGUnknown
}

// GetItemType detect item type via item id.
func (i Inventory) GetItemType() rivalitemenum.ItemType {
	// TODO: Refactor idea?
	//@Tinh comment idea la dung items cua PhuongGa de load (structure lai thanh `map` de load nhanh hon
	switch i.Kind {
	case IKCurrency:
		switch i.Type {
		case CurrencyEHC:
			return rivalitemenum.ITEnhancer
		case CurrencyGold:
			return rivalitemenum.ITGold
		case CurrencyRBuck:
			return rivalitemenum.ITRivalBuck
		case CurrencySSPoint:
			return rivalitemenum.ITSeasonPoint
		case CurrencySSBooster:
			return rivalitemenum.ITSeasonBooster
		}

	case IKBox:
		switch rivalboxenum.RivalBoxType(i.Type) {
		case rivalboxenum.RivalBox:
			return rivalitemenum.ITRivalBox
		case rivalboxenum.BigBox:
			return rivalitemenum.ITBigBox
		case rivalboxenum.MegaBox:
			return rivalitemenum.ITMegaBox
		case rivalboxenum.EventBox:
			return rivalitemenum.ITEventBox
		}

	case IKCosmeticProfile:
		switch i.Type / 10_000 {
		case 0:
			return rivalitemenum.ITAvatar
		case 2:
			return rivalitemenum.ITAvatarFrame
		case 3:
			return rivalitemenum.ITEmoticon
		case 5:
			return rivalitemenum.ITNameColor

		}

	case IKCosmeticAddIn:
		switch i.Type / 100_0000 {
		case 10:
			return rivalitemenum.ITGlow
		case 12:
			return rivalitemenum.ITDance
		case 13:
			return rivalitemenum.ITFootprint
		case 14:
			return rivalitemenum.ITVehicle
		case 15:
			return rivalitemenum.ITBackBling
		case 16:
			return rivalitemenum.ITFlyCraft
		case 17:
			return rivalitemenum.ITVoice

		}
	}

	return rivalitemenum.ITNone
}
