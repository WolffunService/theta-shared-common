package auditmodel

import "github.com/WolffunService/theta-shared-common/proto/auditproto"

type TrophyAudits struct {
	UserId       string `json:"userId" bson:"userId"`
	Trophy       int    `json:"trophy" bson:"trophy"`
	TrophyChange int    `json:"trophyChange" bson:"trophyChange"`
	Timestamp    int64  `json:"timestamp" bson:"timestamp"`
	Source       int    `json:"source" bson:"source"`
}

func (t TrophyAudits) FromProto(a *auditproto.TrophyExpChangeResult) TrophyAudits {
	return TrophyAudits{
		UserId:       a.UserId,
		Trophy:       int(a.UserTrophy),
		TrophyChange: int(a.TrophyChange),
		Timestamp:    a.Timestamp,
		Source:       int(a.Source),
	}
}
