package auditmodel

import (
	"github.com/WolffunService/theta-shared-common/enums/userenum"
	"github.com/WolffunService/theta-shared-common/proto/auditproto"
	"log"
	"strconv"
)

type BanUnbanAudit struct {
	Executor    string                `json:"executor" bson:"executor"`
	UserId      string                `json:"userId" bson:"userId"`
	Type        userenum.BanAuditType `json:"type" bson:"type"`
	Reason      string                `json:"reason" bson:"reason"`
	ExpiredTime int64                 `json:"expiredTime,omitempty" bson:"expiredTime,omitempty"`
	Timestamp   int64                 `json:"timestamp" bson:"timestamp"`
}

func (s *BanUnbanAudit) FromSimpleEvent(e *auditproto.SimpleEvent) *BanUnbanAudit {
	auditMap := map[string]string{}
	if e.Event != nil {
		for _, v := range e.Event.EventParams {
			auditMap[v.Key] = v.Value
		}
	}

	typee, err := strconv.Atoi(e.Metadata["type"])
	if err != nil {
		log.Println("[error][analytic][ban-unban] cannot convert type to int", e.Metadata["type"])
	}

	expiredTime, err := strconv.ParseInt(e.Metadata["expiredTime"], 10, 0)
	if err != nil {
		log.Println("[error][analytic][ban-unban] cannot convert expiredTime to int", e.Metadata["expiredTime"])
	}

	return &BanUnbanAudit{
		Executor:    e.Metadata["executor"],
		UserId:      e.Metadata["userId"],
		Reason:      e.Metadata["reason"],
		Type:        userenum.BanAuditType(typee),
		ExpiredTime: int64(expiredTime),
		Timestamp:   e.Event.Timestamp,
	}
}
