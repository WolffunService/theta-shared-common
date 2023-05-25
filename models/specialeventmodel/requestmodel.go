package specialeventmodel

import eventtype "github.com/WolffunService/theta-shared-common/enums/specialeventenum"

type JoinReq struct {
	eventtype.EventType `json:"eventType"`
}

type CheckRewardReq struct {
	eventtype.EventType `json:"eventType" form:"eventType"`
	IsDaily             bool `json:"isDaily" form:"isDaily"`
}
