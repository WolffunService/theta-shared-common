package ivtmodel

import (
	"sort"
	"strings"
	"sync"

	"github.com/WolffunService/thetan-shared-common/enums/gamereleaseenum"
	"github.com/WolffunService/thetan-shared-common/enums/rivalitemenum"
	"github.com/WolffunService/thetan-shared-common/enums/thetanrivalerrorenum/thetanrivalconst"
)

type InventoryProperties struct {
	Rarity       thetanrivalconst.CosmeticRarity `json:"rarity,omitempty" bson:"-"`
	Release      gamereleaseenum.GameRelease     `json:"release,omitempty" bson:"-"`
	Season       int                             `json:"season,omitempty" bson:"-"`
	ExclusiveFor rivalitemenum.ExclusiveFeature  `json:"exclusiveFor,omitempty" bson:"-"`
	EventId      int                             `json:"eventId,omitempty" bson:"-"`
	SetId        int                             `json:"setId,omitempty" bson:"-"`
	UniformId    int                             `json:"uniformId,omitempty" bson:"-"`
	SSPassType   string                          `json:"ssPassType,omitempty" bson:"-"`
	Theme        string                          `json:"theme,omitempty" bson:"-"`

	// Automatically assigned, prefix "As"
	AsUnlocked  bool `json:"unlocked,omitempty" bson:"-"`
	AsDailyDeal bool `json:"dailyDeal,omitempty" bson:"-"`
	AsMegaBox   bool `json:"megaBox,omitempty" bson:"-"`
	AsEventBox  bool `json:"eventBox,omitempty" bson:"-"`
	AsBuckShop  bool `json:"buckShop,omitempty" bson:"-"`
}

type ItemPicker struct {
	key  string
	once sync.Once

	DropGroups   []rivalitemenum.DropGroup
	Qty          int
	Filters      any
	ExpectedRate map[string]float64
	OrderFn      func(items any) (any, func(int, int) bool)
}

func (p *ItemPicker) Key() string {
	// Build key once time
	p.once.Do(func() {
		dropGroups := make([]string, len(p.DropGroups))
		for i := range p.DropGroups {
			dropGroups[i] = string(p.DropGroups[i])
		}

		sort.Slice(dropGroups, func(i, j int) bool {
			return dropGroups[i] < dropGroups[j]
		})

		p.key = strings.Join(dropGroups, "-")
	})

	return p.key
}
