package tradingeventmodel

import (
	"errors"

	"github.com/WolffunService/thetan-shared-common/enums/heroenum"
	"github.com/WolffunService/thetan-shared-common/enums/tradingeventenum"
)

type QuestData struct {
	SubID           int                  `json:"subId" bson:"subId"`
	HeroTrophyClass *int                 `json:"heroTrophyClass,omitempty" bson:"heroTrophyClass,omitempty"`
	HeroLevel       *int                 `json:"heroLevel,omitempty" bson:"heroLevel,omitempty"`
	HeroClass       *heroenum.HeroRole   `json:"heroClass,omitempty" bson:"heroClass,omitempty"`
	SkinRarity      *heroenum.SkinRarity `json:"skinRarity,omitempty" bson:"skinRarity,omitempty"`
}

func (q *QuestData) InjectData(questType tradingeventenum.QuestType, data int, index int) error {
	q.SubID = index
	switch questType {
	case tradingeventenum.QT_HERO_TROPHY_CLASS:
		q.HeroTrophyClass = &data
	case tradingeventenum.QT_HERO_LEVEL:
		q.HeroLevel = &data
	case tradingeventenum.QT_HERO_CLASS:
		q.HeroClass = (*heroenum.HeroRole)(&data)
	case tradingeventenum.QT_SKIN_RARITY:
		q.SkinRarity = (*heroenum.SkinRarity)(&data)
	case tradingeventenum.QT_ALL_HEROES:
		break
	default:
		return errors.New("wrong internal config")
	}

	return nil
}

type TradingQuest struct {
	Type      tradingeventenum.QuestType `json:"type,omitempty" bson:"type"`
	QuestData `json:",inline" bson:",inline"`
}
