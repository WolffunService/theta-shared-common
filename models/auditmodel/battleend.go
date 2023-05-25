package auditmodel

import (
	"github.com/WolffunService/theta-shared-common/proto/auditproto"
	"log"
	"strconv"
)

type BattleEndAudit struct {
	ID           interface{} `json:"id" bson:"_id,omitempty"`
	UserId       string      `json:"userid" bson:"userid"`
	MatchId      string      `json:"matchid" bson:"matchid"`
	BattleNumber int         `json:"battlenumber" bson:"battlenumber"`
	HeroId       string      `json:"heroid" bson:"heroid"`
	Rank         int         `json:"rank" bson:"rank"`
	IngameMode   int         `json:"ingamemode" bson:"ingamemode"`
	Result       int         `json:"result" bson:"result"`
	TrophyChange int         `json:"trophychange" bson:"trophychange"`
	Trophy       int         `json:"trophy" bson:"trophy"`
	ExpChange    int         `json:"expchange" bson:"expchange"`
	PowerPoint   int         `json:"powerPoint" bson:"powerPoint"`
	ThetanCoin   float64     `json:"thetanCoin" bson:"thetanCoin"`
	Timestamp    int64       `json:"timestamp" bson:"timestamp"`
	HeroStatus   int         `json:"heroStatus" bson:"herostatus"`
}

func (b *BattleEndAudit) GetBattleEndAudit(e *auditproto.SimpleEvent) *BattleEndAudit {
	auditMap := map[string]string{}
	if e.Event != nil {
		for _, v := range e.Event.EventParams {
			auditMap[v.Key] = v.Value
		}
	}

	ingameMode, err := strconv.Atoi(e.Metadata["ingameMode"])
	if err != nil {
		log.Println("[error][battleend] cannot get ingameMode")
	}
	trophyChange, err := strconv.Atoi(e.Metadata["trophyChange"])
	if err != nil {
		log.Println("[error][battleend] cannot get trophyChange")
	}
	trophy, err := strconv.Atoi(e.Metadata["trophy"])
	if err != nil {
		log.Println("[error][battleend] cannot get trophy")
	}
	expChange, err := strconv.Atoi(e.Metadata["expChange"])
	if err != nil {
		log.Println("[error][battleend] cannot get expChange")
	}
	battleNumber, err := strconv.Atoi(auditMap["battleNumber"])
	if err != nil {
		log.Println("[error][battleend] cannot get battleNumber")
	}
	rank, err := strconv.Atoi(e.Metadata["battleRank"])
	if err != nil {
		log.Println("[error][battleend] cannot get battleRank")
	}
	result, err := strconv.Atoi(auditMap["result"])
	if err != nil {
		log.Println("[error][battleend] cannot get result")
	}
	heroStatus, err := strconv.Atoi(e.Metadata["heroStatus"])
	if err != nil {
		log.Println("[error][battleend] cannot get heroStatus")
	}

	thetanCoin, err := strconv.ParseFloat(e.Metadata["thetanCoin"], 64)
	if err != nil {
		// log.Println("[error][[battleend] cannot get thetanCoin", err)
	}

	powerPoint, err := strconv.Atoi(e.Metadata["powerPoint"])
	if err != nil {
		// log.Println("[error][battleend] cannot get powerPoint")
	}

	return &BattleEndAudit{
		IngameMode:   ingameMode,
		TrophyChange: trophyChange,
		Trophy:       trophy,
		ExpChange:    expChange,
		ThetanCoin:   thetanCoin,
		PowerPoint:   powerPoint,

		UserId:       auditMap["userId"],
		MatchId:      auditMap["matchId"],
		BattleNumber: battleNumber,
		HeroId:       auditMap["heroId"],
		Rank:         rank,
		Result:       result,
		Timestamp:    e.Event.Timestamp,

		HeroStatus: heroStatus,
	}
}
