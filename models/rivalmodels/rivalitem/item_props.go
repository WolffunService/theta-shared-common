package rivalitem

import (
	"time"

	"github.com/WolffunService/theta-shared-common/enums/rivalitemenum"
)

type ItemProps struct {
	ItemManualProps       `json:",inline" yaml:",inline" mapstructure:",squash" bson:"-" `
	ItemAutoAssignedProps `json:",inline" yaml:",inline" mapstructure:",squash" bson:"-" `
}

type ItemManualProps struct {
	// Required field
	Rarity       rivalitemenum.ItemRarity       `json:"rarity" yaml:"rarity" mapstructure:"rarity" bson:"-"`
	Release      time.Time                      `json:"release" yaml:"release" mapstructure:"release" bson:"-"`
	Season       int                            `json:"season" yaml:"season" mapstructure:"season" bson:"-"`
	ExclusiveFor rivalitemenum.ExclusiveFeature `json:"exclusiveFor" yaml:"exclusiveFor" mapstructure:"exclusiveFor" bson:"-"`

	// Optional field
	EventId   int    `json:"eventId,omitempty" yaml:"eventId,omitempty" mapstructure:"eventId" bson:"-"`
	SetId     int    `json:"setId,omitempty" yaml:"setId,omitempty" mapstructure:"setId" bson:"-"`
	UniformId int    `json:"uniformId,omitempty" yaml:"uniformId,omitempty" mapstructure:"uniformId" bson:"-"`
	PassType  string `json:"ssPassType,omitempty" yaml:"ssPassType,omitempty" mapstructure:"ssPassType" bson:"-"`
	Theme     string `json:"theme,omitempty" yaml:"theme,omitempty" mapstructure:"theme" bson:"-"`
}

type ItemAutoAssignedProps struct {
	// IsAssigned Cờ để kiểm tra xem các props đã được tính toán chưa.
	IsAssigned bool `json:"-" yaml:"-" mapstructure:"-" bson:"-"`

	// Automatically assigned, prefix "As"
	AsUnlocked  bool `json:"unlocked,omitempty" yaml:"-" mapstructure:"-" bson:"-"`
	AsDailyDeal bool `json:"dailyDeal,omitempty" yaml:"-" mapstructure:"-" bson:"-"`
	AsMegaBox   bool `json:"megaBox,omitempty" yaml:"-" mapstructure:"-" bson:"-"`
	AsEventBox  bool `json:"eventBox,omitempty" yaml:"-" mapstructure:"-" bson:"-"`
	AsBuckShop  bool `json:"buckShop,omitempty" yaml:"-" mapstructure:"-" bson:"-"`
	AsAdvantage bool `json:"advantage,omitempty" yaml:"-" mapstructure:"-" bson:"-"`
}
