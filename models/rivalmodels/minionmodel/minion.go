package minionmodel

import (
	"time"

	"github.com/WolffunService/theta-shared-common/enums/rivalitemenum"
	"github.com/WolffunService/theta-shared-common/enums/thetanrivalerrorenum/thetanrivalconst"
	"github.com/WolffunService/theta-shared-common/enums/thetanrivalerrorenum/thetanrivalminionenum"
	"github.com/WolffunService/theta-shared-common/models/rivalmodels/rivalitem"
	"github.com/mitchellh/mapstructure"
)

type MinionGender uint8

const (
	MinionMale MinionGender = 1 + iota
	MinionFemale
)

type BaseMinion struct {
	Type   int                           `json:"type" yaml:"type"`
	Name   string                        `json:"name" yaml:"name"`
	Rarity thetanrivalconst.MinionRarity `json:"rarity" yaml:"rarity"`
	Gender MinionGender                  `json:"gender" yaml:"gender" mapstructure:"gender"`
}

type Minion struct {
	BaseMinion `json:",inline" yaml:",inline" mapstructure:",squash"`
	Skins      []MinionSkin `json:"skins" yaml:"skins"`
}

type MinionSkin struct {
	Id           int                               `json:"id" yaml:"id" mapstructure:"id"`
	Name         string                            `json:"name" yaml:"name" mapstructure:"name"`
	Rarity       thetanrivalconst.MinionSkinRarity `json:"rarity" yaml:"rarity" mapstructure:"rarity"`
	Release      time.Time                         `json:"release,omitempty" yaml:"release,omitempty" mapstructure:"release"`
	ExclusiveFor rivalitemenum.ExclusiveFeature    `json:"exclusiveFor" yaml:"exclusiveFor,omitempty" mapstructure:"exclusiveFor"`
	Season       int                               `json:"season" yaml:"season" mapstructure:"season"`
	EventId      int                               `json:"eventId" yaml:"eventId,omitempty" mapstructure:"eventId"`
	SetId        int                               `json:"setId" yaml:"setId,omitempty" mapstructure:"setId"`
	UniformId    int                               `json:"uniformId" yaml:"uniformId,omitempty" mapstructure:"uniformId"`
	Theme        string                            `json:"theme" yaml:"theme,omitempty" mapstructure:"theme"`
}

// MinionSkinFlatten all information about minion and skin
type MinionSkinFlatten struct {
	MinionSkin `json:",inline"`

	Minion BaseMinion `json:",inline"`
}

type MinionProperties struct {
	Rarity       thetanrivalconst.MinionSkinRarity `json:"rarity,omitempty" bson:"-"`
	ReleaseDate  *time.Time                        `json:"releaseDate,omitempty" bson:"-"`
	Season       int                               `json:"seasons,omitempty" bson:"-"`
	ExclusiveFor rivalitemenum.ExclusiveFeature    `json:"exclusiveFor,omitempty" bson:"-"`
	EventId      int                               `json:"eventId,omitempty" bson:"-"`
	SetId        int                               `json:"setId,omitempty" bson:"-"`
	UniformId    int                               `json:"uniformId,omitempty" bson:"-"`
	//SSPassType   string                            `json:"ssPassType,omitempty" bson:"-"`
	Theme string `json:"theme,omitempty" bson:"-"`

	// Automatically assigned, prefix "As"
	AsUnlocked  bool `json:"unlocked,omitempty" bson:"-"`
	AsDailyDeal bool `json:"dailyDeal,omitempty" bson:"-"`
	AsMegaBox   bool `json:"megaBox,omitempty" bson:"-"`
	AsEventBox  bool `json:"eventBox,omitempty" bson:"-"`
	AsBuckShop  bool `json:"buckShop,omitempty" bson:"-"`
}

type MinionSkinFull struct {
	MinionSkinFlatten `json:",inline"`

	Props MinionProperties `json:"props"`
}

func (m MinionSkinFull) GetDropGroup() rivalitemenum.DropGroup {
	return rivalitemenum.DGMinion
}

func (minion Minion) ToMinionBot() thetanrivalminionenum.MinionBot {
	bot := thetanrivalminionenum.MinionBot{
		Type:   minion.Type,
		Name:   minion.Name,
		Rarity: minion.Rarity,
		Skins:  []thetanrivalminionenum.MinionSkinBot{},
	}

	return bot
}

func (skin MinionSkin) ToBotSkin() thetanrivalminionenum.MinionSkinBot {
	skinBot := thetanrivalminionenum.MinionSkinBot{
		Id:        skin.Id,
		Name:      skin.Name,
		Rarity:    skin.Rarity,
		UniformId: skin.UniformId,
	}
	return skinBot
}

func (skin MinionSkin) ToItemDetail() (item *rivalitem.ItemDetail, err error) {
	skinWrapped := struct {
		MinionSkin `mapstructure:",squash"`
		Props      any `mapstructure:"props"`
	}{
		MinionSkin: skin,
		Props:      skin,
	}

	item = &rivalitem.ItemDetail{}
	err = mapstructure.Decode(skinWrapped, item)

	// TODO: hard code
	item.Props.Release = skin.Release

	return item, err
}
