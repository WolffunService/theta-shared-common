package tradingeventmodel

import (
	"encoding/json"
	"time"

	"github.com/WolffunService/theta-shared-common/database/mongodb"
	"github.com/WolffunService/theta-shared-common/enums/currencyenum"
	"github.com/WolffunService/theta-shared-common/enums/tradingeventenum"
	"github.com/WolffunService/theta-shared-common/models/currencymodel"
)

//functask: xử lý vụ
/*
4 cái event types á giả sử event chạy được mấy tháng r
nếu chị thêm event thứ 5 là nó ra cái event 5 hoài á
*/
type TradingEventSeason struct {
	mongodb.IDIntField `json:",inline" bson:",inline"`
	mongodb.DateFields `json:"-" bson:",inline"`

	// cash
	TotalCash currencymodel.SystemCurrency `json:"-" bson:"totalCash"`

	// reward
	Cashback    SeasonCashback `json:"cashback" bson:"cashback"`
	Leaderboard []LBRank       `json:"-" bson:"lbReward"`
	MaxLBPool   int64          `json:"-" bson:"maxLBPool"` // real value of gTHG, lỡ lưu bằng int64
	MinLBPool   int64          `json:"-" bson:"minLBPool"` // real value of gTHG

	// setup
	StartTime  time.Time                                 `json:"startTime" bson:"startTime"`
	EndTime    time.Time                                 `json:"endTime" bson:"endTime"`
	Quest      TradingQuest                              `json:"quest" bson:"quest"`
	WarmUp     []TradingQuest                            `json:"warmUp" bson:"warmUp"`
	QuestStats map[tradingeventenum.QuestType]QuestStats `json:"-" bson:"questStats"`

	ExRate float64 `json:"-" bson:"exRate"`
}

type QuestStats struct {
	Showed    uint         `json:"showed" bson:"showed"`
	SubQuests map[int]uint `json:"subquest" bson:"subquest"`
}

func (*TradingEventSeason) CollectionName() string {
	return "TradingEventSeasons"
}

func (f TradingEventSeason) IsRunning() bool {
	now := time.Now().UTC()
	return f.StartTime.Before(now) && f.EndTime.After(now)
}

func (f TradingEventSeason) IsPast() bool {
	return f.EndTime.Before(time.Now().UTC())
}

// GetDay return relative day from startTime
func (f TradingEventSeason) GetDay(isProduction ...bool) int {
	if len(isProduction) == 0 || isProduction[0] {
		return int(time.Since(f.StartTime).Hours())/24 + 1
	}

	return int(time.Since(f.StartTime).Minutes()/15) + 1 // warmup
}

func (f TradingEventSeason) GetEndOfDay(isProduction ...bool) time.Time {
	n := f.GetDay(isProduction...)

	if f.IsPast() {
		return f.EndTime
	} else if f.StartTime.After(time.Now()) {
		return f.StartTime
	}

	if len(isProduction) == 0 || isProduction[0] {
		return f.StartTime.Add(time.Duration(n) * 24 * time.Hour)
	}

	return f.StartTime.Add(time.Duration(n) * 15 * time.Minute) // warmup
}

func (f TradingEventSeason) IsInRewardRank(rank int) bool {
	for i := range f.Leaderboard {
		if int(f.Leaderboard[i].Min) <= rank && rank <= int(f.Leaderboard[i].Max) {
			return true
		}
	}

	return false
}

// ----------------- Cashback
type CashbackConfig struct {
	Score  currencymodel.SC `json:"score" bson:"score"`
	Rate   float64          `json:"rate,omitempty" bson:"rate"`
	Reward currencymodel.SC `json:"reward" bson:"reward"` // auto
}

var _ json.Marshaler = (*CashbackConfig)(nil)

func (f *CashbackConfig) MarshalJSON() ([]byte, error) {
	newModel := *f
	newModel.Rate = 0

	return json.Marshal(newModel)
}

type SeasonCashback []CashbackConfig

func (s SeasonCashback) InjectRewards(exRate float64, feePercent float64, rewardType currencyenum.Currency) SeasonCashback {
	var pre float64 = 0
	for i := range s {
		// validate cashback reward
		gap := s[i].Score.GetRealValue() - pre
		s[i].Reward.SystemCurrency = currencymodel.ConvertFloatToSystemCurrency(
			gap*feePercent/100/2*s[i].Rate*exRate,
			rewardType,
		)

		pre = s[i].Score.GetRealValue()
	}

	return s
}

type LBRank struct {
	Min        float64 `json:"min" bson:"min"`
	Max        float64 `json:"max" bson:"max"`
	RangeShare float64 `json:"share" bson:"share"`
}
