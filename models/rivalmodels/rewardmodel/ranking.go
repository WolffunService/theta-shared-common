package rewardmodel

import "github.com/WolffunService/theta-shared-common/models/rivalmodels/ivtmodel"

type SystemRewardType uint

const (
	MinionItem SystemRewardType = iota
	CurrencyItem
	CosmeticProfileItem
	CosmeticAddInItem
	BoxItem
)

func (r SystemRewardType) GetKind() ivtmodel.InventoryKind {
	switch r {
	case CurrencyItem:
		return ivtmodel.IKCurrency
	//case CosmeticProfileItem:
	//	return ivtmodel.IKCosmeticProfile
	case CosmeticAddInItem:
		return ivtmodel.IKCosmeticAddIn
	case BoxItem:
		return ivtmodel.IKBox
	}
	panic("Undefined item")
}
