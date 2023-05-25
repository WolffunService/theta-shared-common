package auditmodel

import (
	"log"
	"strconv"

	"github.com/WolffunService/thetan-shared-common/proto/auditproto"
)

type StartMatch struct {
	MatchId    string `json:"matchId" bson:"matchId"`
	UserId     string `json:"userId" bson:"userId"`
	HeroId     string `json:"heroId" bson:"heroId"`
	IngameMode int    `json:"ingameMode" bson:"ingameMode"`
	Timestamp  int64  `json:"timestamp" bson:"timestamp"`
}

func (s *StartMatch) FromSimpleEvent(e *auditproto.SimpleEvent) *StartMatch {
	ingameMode, err := strconv.Atoi(e.Metadata["ingameMode"])
	if err != nil {
		log.Println("[error][analytic][start-match] cannot convert ingameMode to int", e.Metadata["ingameMode"])
	}

	return &StartMatch{
		MatchId:    e.Metadata["matchId"],
		HeroId:     e.Metadata["heroId"],
		UserId:     e.Metadata["userId"],
		IngameMode: ingameMode,
		Timestamp:  e.Event.Timestamp,
	}
}
