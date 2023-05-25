package efilter

import (
	"github.com/WolffunService/theta-shared-common/enums/rivalitemenum"
	"github.com/WolffunService/theta-shared-common/models/rivalmodels/ivtmodel"
	"github.com/WolffunService/theta-shared-common/models/rivalmodels/rivalitem"
	"github.com/samber/lo"
)

type ItemPropsFilter func(item rivalitem.ItemDetail) bool

func ApplyFilters(item rivalitem.ItemDetail, filters ...ItemPropsFilter) bool {
	for _, fn := range filters {
		if !fn(item) {
			return false
		}
	}

	return true
}

func LogicOr(filters ...ItemPropsFilter) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		for i := 0; i < len(filters); i++ {
			if filters[i](item) {
				return true
			}
		}

		return false
	}
}

func LogicAnd(filters ...ItemPropsFilter) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		for i := 0; i < len(filters); i++ {
			if !filters[i](item) {
				return false
			}
		}

		return true
	}
}

func Kind(kind ivtmodel.InventoryKind) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.InventoryKind == kind
	}
}

func DropGroup(dropGroup rivalitemenum.DropGroup) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.DropGroup() == dropGroup
	}
}

func AnyDropGroups(dropGroups ...rivalitemenum.DropGroup) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return lo.Contains(dropGroups, item.DropGroup())
	}
}

func NotDropGroup(dropGroup rivalitemenum.DropGroup) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.DropGroup() != dropGroup
	}
}

func ItemType(typ rivalitemenum.ItemType) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.ItemType == typ
	}
}

func Rarity(rarity rivalitemenum.ItemRarity) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.Rarity == rarity
	}
}

func Season(seasons ...int) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return lo.Contains(seasons, item.Props.Season)
	}
}

func Exclusive(exclusive rivalitemenum.ExclusiveFeature) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.ExclusiveFor == exclusive
	}
}

func EventId(eventId int) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.EventId == eventId
	}
}

func SetId(setId int) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.SetId == setId
	}
}

func UniformId(uniformId int) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.UniformId == uniformId
	}
}

func PassType(passType string) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.PassType == passType
	}
}

func NotPassType(passType string) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.PassType != passType
	}
}

func Theme(theme string) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.Theme == theme
	}
}

func Unlocked(unlocked bool) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.AsUnlocked == unlocked
	}
}

func DailyDeal(dd bool) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.AsDailyDeal == dd
	}
}

func MegaBox(mb bool) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.AsMegaBox == mb
	}
}

func EventBox(eb bool) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.AsEventBox == eb
	}
}

func BuckShop(bs bool) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.AsBuckShop == bs
	}
}

func Advantage(bs bool) ItemPropsFilter {
	return func(item rivalitem.ItemDetail) bool {
		return item.Props.AsAdvantage == bs
	}
}
