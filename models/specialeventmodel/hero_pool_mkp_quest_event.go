package specialeventmodel

import (
	"github.com/WolffunService/thetan-shared-common/enums/heroenum"
	"github.com/func25/slicesol/slicesol"
)

type HeroPoolMKPQuestEvent struct {
	BaseEvent       `json:",inline" bson:",inline"`
	HeroPools       map[heroenum.HeroRarity]slicesol.Sliol[int] `json:"heroPools" bson:"heroPools"`
	HeroesAvailable map[heroenum.HeroRarity]slicesol.Sliol[int] `json:"heroesAvailable" bson:"heroesAvailable"`
}
