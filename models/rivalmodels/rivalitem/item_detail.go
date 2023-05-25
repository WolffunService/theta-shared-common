package rivalitem

import (
	"github.com/WolffunService/theta-shared-common/enums/rivalitemenum"
	"github.com/WolffunService/theta-shared-common/models/rivalmodels/ivtmodel"
)

// ItemDetail Sử dụng ItemGeneric của Phương và ItemProps của Trung
type ItemDetail ItemG[ItemProps]

// DropGroup TODO: Cứuuuuuuuuuuuuu tui, cứu tui
func (item ItemDetail) DropGroup() rivalitemenum.DropGroup {
	switch item.InventoryKind {
	case ivtmodel.IKMinion:
		return rivalitemenum.DGMinion
	case ivtmodel.IKCosmeticProfile:
		return rivalitemenum.DGProfile
	case ivtmodel.IKCosmeticAddIn:
		switch item.ItemType {
		case rivalitemenum.ITGlow, rivalitemenum.ITVoice, rivalitemenum.ITDance, rivalitemenum.ITVehicle:
			return rivalitemenum.DGFullEvolve
		case rivalitemenum.ITFootprint, rivalitemenum.ITBackBling, rivalitemenum.ITFlyCraft:
			return rivalitemenum.DGHalfEvolve
		}
	}

	return rivalitemenum.DGUnknown
}

func (item ItemDetail) GetItemRarity() rivalitemenum.ItemRarity {
	return item.Props.Rarity
}
