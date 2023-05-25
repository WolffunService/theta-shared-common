package herofusionmodel

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/enums/heroenum"
	"github.com/WolffunService/thetan-shared-common/models/currencymodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ---------- INPUT
type HeroFusionTurn struct {
	mongodb.DefaultModel `bson:",inline"`
	mongodb.DateFields   `json:"-" bson:",inline"`

	// current
	InpHeroes []int `json:"inpHeroes" bson:"inpHeroes"` // skinIDs
	OutHeroes []int `json:"outHeroes" bson:"outHeroes"` // skinIDs
	SeasonID  int64 `json:"seasonId" bson:"seasonId"`

	// cashback & reward
	CashbackMilestones map[int]bool `json:"cbMilestones" bson:"cbMilestones"` // true là có thể claim, vì false là giá trị default cho những milestone không thể claim

	// statistic
	Showed map[int]int `json:"-" bson:"showed"`
	// HeroOV         float64     `json:"heroOv" bson:"heroOv"` // Cache lại lúc refresh cost để tính cho fe
	ChangeRate     float64 `json:"-" bson:"changeRate"` // for analytic
	FusionCostTurn int     `json:"-" bson:"fcostTurn"`  // Bộ đếm changeRate trong công thức fusionCost, chỉ reset khi nhấn fuse thành công

	// refresh state
	RefreshCost currencymodel.SystemCurrency `json:"refreshCost" bson:"refreshCost"`
	RefreshTime int32                        `json:"-" bson:"refreshTime"`

	// fusion state
	FusionCost  *currencymodel.SystemCurrency `json:"fusionCost" bson:"fusionCost"`
	CostExpired int64                         `json:"costExpired" bson:"costExpired"`
	EndTime     int64                         `json:"endTime" bson:"endTime"`
	MatIDs      map[string]bool               `json:"-" bson:"matIds"`
	FuseTime    int32                         `json:"fuseTime" bson:"fuseTime"`

	// fusion progress
	Status   heroenum.HeroFusionStatus `json:"status" bson:"status"`
	Progress *FusionProgress           `json:"-" bson:"progress"`

	// hero pool
	PoolID int `json:"poolId" bson:"poolId"`
}

func (h HeroFusionTurn) GetStrID() string {
	return h.ID.(primitive.ObjectID).Hex()
}

func (HeroFusionTurn) CollectionName() string {
	return "HeroFusionTurns"
}

type FusionProgress struct {
	BlockchainID uint64                       `json:"blockchainId" bson:"blockchainId"`
	Materials    []string                     `json:"materials" bson:"materials"`
	OutputSkinID int                          `json:"outputSkinId" bson:"outputSkinId"`
	FusionCost   currencymodel.SystemCurrency `json:"fusionCost" bson:"fusionCost"`
	SeasonID     int64                        `json:"seasonId" bson:"seasonId"`
	Timestamp    time.Time                    `json:"timestamp" bson:"timestamp"`
}
