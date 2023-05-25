package metric

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/enums/inventoryenum"
	"github.com/WolffunService/thetan-shared-common/enums/rivalitemenum"
	"github.com/kamva/mgm/v3"
)

type DBPrefix string
type MetricDropGroup string
type MetricAction string

const (
	DBPrefixThetanRivals DBPrefix = "Rivals_"
)

const (
	DropGroupMinion     MetricDropGroup = "m"
	DropGroupEnhancer   MetricDropGroup = "e"
	DropGroupProfile    MetricDropGroup = "p"
	DropGroupHalfEvolve MetricDropGroup = "he"
	DropGroupFullEvolve MetricDropGroup = "fe"
	DropGroupGold       MetricDropGroup = "g"
	DropGroupSPoint     MetricDropGroup = "s"

	ActionBeginSession    MetricAction = "bs"
	ActionEndSession      MetricAction = "es"
	ActionEndShortSession MetricAction = "ess"
	ActionShopVisit       MetricAction = "sv"
	ActionFirstPlay       MetricAction = "fp"
	ActionLobbyStart      MetricAction = "ls"
	ActionLobbyEnd        MetricAction = "le"
)

func ItemDGToMetricDG(dg rivalitemenum.DropGroup) MetricDropGroup {
	switch dg {
	case rivalitemenum.DGMinion:
		return DropGroupMinion
	case rivalitemenum.DGProfile:
		return DropGroupProfile
	case rivalitemenum.DGFullEvolve:
		return DropGroupFullEvolve
	case rivalitemenum.DGHalfEvolve:
		return DropGroupHalfEvolve
	case rivalitemenum.DGGold:
		return DropGroupGold
	case rivalitemenum.DGSPoint:
		return DropGroupSPoint
	case rivalitemenum.DGEnhancer:
		return DropGroupEnhancer

	}

	return ""
}

type CommonMetadata struct {
	UserId string
	Stage  int
}

type PurchaseData struct {
	Stage  int
	ItemID string
	Value  float64
	Battle int
	Gap    int
}

type BattleData struct {
	Stage int
	Round int
	Rank  int
}

type ItemData struct {
	Stage     int
	ItemID    string
	DropGroup MetricDropGroup
	Source    inventoryenum.ChangeSource
	Amount    float64
}

type ActionData struct {
	Action MetricAction
}

type Record[TS TimeSeriesModel] struct {
	mgm.IDField `json:",inline" bson:",inline"`
	Timestamp   time.Time `json:"timestamp" bson:"timestamp"`
	TrueTime    time.Time `json:"true_time" bson:"true_time"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Data        TS        `json:"data" bson:",inline"`
	Others      any       `json:"others" bson:"others"`
}

func (ItemData) CollectionName() string {
	return string(DBPrefixThetanRivals) + "Item"
}

func (BattleData) CollectionName() string {
	return string(DBPrefixThetanRivals) + "Battle"
}

func (PurchaseData) CollectionName() string {
	return string(DBPrefixThetanRivals) + "Purchase"
}

func (ActionData) CollectionName() string {
	return string(DBPrefixThetanRivals) + "Action"
}

func (r Record[TS]) CollectionName() string {
	if casted, ok := any(r.Data).(mgm.CollectionNameGetter); ok {
		return casted.CollectionName()
	}

	//Should not happen
	return string(DBPrefixThetanRivals) + "Common"
}

type TimeSeriesModel interface {
	BattleData | PurchaseData | ItemData | ActionData

	// CollectionName Đừng tin thằng Trung, syntax này IDE nó bị lừa nhưng compiler thì không
	// CollectionName() string
}

type ActionDataLobby struct {
	TownID       string `json:"townID" bson:"townID"`
	UserLanguage string `json:"userLanguage" bson:"userLanguage"`
	UserAge      int    `json:"userAge" bson:"userAge"`
}
