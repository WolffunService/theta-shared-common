package herofusionmodel

import (
	"encoding/json"
	"time"

	"github.com/WolffunService/theta-shared-common/database/mongodb"
	"github.com/WolffunService/theta-shared-common/models/currencymodel"
)

type HeroFusionSeason struct {
	mongodb.IDIntField `json:",inline" bson:",inline"`
	mongodb.DateFields `json:"-" bson:",inline"`
	TotalCash          currencymodel.SystemCurrency `json:"-" bson:"totalCash"`
	MaxCash            float64                      `json:"-" bson:"maxCash"`
	MinPool            float64                      `json:"-" bson:"minPool"`
	Cashback           []FusionCashbackConfig       `json:"cashback" bson:"cashback"`
	StartTime          time.Time                    `json:"startTime" bson:"startTime"`
	EndTime            time.Time                    `json:"endTime" bson:"endTime"`
}

func (HeroFusionSeason) CollectionName() string {
	return "HeroFusionSeasons"
}

func (f HeroFusionSeason) IsRunning() bool {
	return f.StartTime.Before(time.Now()) && f.EndTime.After(time.Now())
}

type FusionCashbackConfig struct {
	Spend  currencymodel.SystemCurrency `json:"spend" bson:"spend"`
	Rate   float64                      `json:"rate,omitempty" bson:"rate"`
	Reward currencymodel.SystemCurrency `json:"reward" bson:"reward"` // auto
}

var _ json.Marshaler = (*FusionCashbackConfig)(nil)

func (f *FusionCashbackConfig) MarshalJSON() ([]byte, error) {
	newModel := *f
	newModel.Rate = 0

	return json.Marshal(newModel)
}
