package rivalitem

import (
	"fmt"
	"strconv"
	"time"

	"github.com/WolffunService/thetan-shared-common/enums/rivalitemenum"
	"github.com/mitchellh/mapstructure"
)

type ItemG[T any] struct {
	Name         string `json:"name" yaml:"name" mapstructure:"name" bson:"name"`
	Desc         string `json:"desc" yaml:"desc,omitempty" mapstructure:"desc" bson:"desc"`
	SimplifyItem `json:",inline" yaml:",inline" mapstructure:",squash" bson:",inline"`
	Tags         []rivalitemenum.Tag `json:"tags" yaml:"tags" mapstructure:"tags" bson:"tags"`
	Props        T                   `json:"props" yaml:"props" mapstructure:"props" bson:"props"`
}

// refactor Item -> ItemG sau
type Item ItemG[map[string]any]

func (i Item) ToSimple() *SimpleItem {
	return &SimpleItem{
		SimplifyItem: SimplifyItem{
			ItemType:      i.ItemType,
			TypeID:        i.TypeID,
			Amount:        i.Amount,
			InventoryKind: i.InventoryKind,
		},
		Props: i.Props,
	}
}

func (i Item) ToItemDetail() (item *ItemDetail, err error) {
	item = &ItemDetail{}
	err = mapstructure.Decode(i, item)

	// TODO: hard-code
	if val, found := i.Props["release"]; found {
		item.Props.Release = val.(time.Time)
	}

	return item, err
}

func (i Item) Rarity() rivalitemenum.ItemRarity {
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

func (i ItemG[T]) ToItem() (item *Item, err error) {
	item = &Item{}
	err = mapstructure.Decode(i, item)
	return item, err
}

// func NewItem(name string, itemType rivalitemenum.ItemType, typeID int, tags ...Tag) Item {
// 	return Item{
// 		Name:     name,
// 		Desc:     "",
// 		ItemType: itemType,
// 		TypeID:   typeID,
// 		Tags:     tags,
// 		Props:    make(map[string]any),
// 		Amount:   utilmodel.NumberDefault(0),
// 	}
// }

// func (i Item) Clone() (*Item, error) {
// 	item := &Item{}
// 	err := copier.CopyWithOption(item, i, copier.Option{DeepCopy: true})
// 	return item, err // return item, copier.Copy(item, i)
// }

// func (i SimpleItem) Clone() (*SimpleItem, error) {
// 	item := &SimpleItem{}
// 	err := copier.CopyWithOption(item, i, copier.Option{DeepCopy: true})
// 	return item, err // return item, copier.Copy(item, i)
// }

// func (i Item) Amount() float64, er
