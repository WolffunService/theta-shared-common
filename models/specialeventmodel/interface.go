package specialeventmodel

import (
	"time"

	eventtype "github.com/WolffunService/theta-shared-common/enums/specialeventenum"

	"github.com/WolffunService/theta-shared-common/database/mongodb"
)

type IEvent interface {
	GetID() interface{}
	GetEventId() int
	GetEventType() eventtype.EventType
	IsActive() bool
	Ended() bool
	CollectionName() string
	SetID(id interface{})
	SetEventId(eventId int)
}

type BaseEvent struct {
	mongodb.DefaultModel `json:"-" bson:",inline"`
	mongodb.DateFields   `json:"-" bson:",inline"`
	CreatorId            string `json:"-" bson:"creatorId,omitempty"`
	EventId              int    `json:"eventId" bson:"eventId"` //pk

	//admin input
	EventType eventtype.EventType `json:"eventType"binding:"enum=eventtype.EventType" bson:"eventType"` //pk
	StartIn   int64               `json:"startIn" bson:"startIn"`
	Duration  int64               `json:"duration" bson:"duration"`
}

func (b BaseEvent) GetEventId() int {
	return b.EventId
}

func (b BaseEvent) GetEventType() eventtype.EventType {
	return b.EventType
}

func (b BaseEvent) IsActive() bool {
	curTimeStamp := time.Now().UTC().Unix()
	return curTimeStamp >= b.StartIn && b.EndIn()-10 > curTimeStamp //10s for cronjob
}

func (b BaseEvent) Ended() bool {
	curTimeStamp := time.Now().UTC().Unix()
	return b.EndIn()-10 < curTimeStamp //10s for cronjob
}

func (b BaseEvent) EndIn() int64 {
	return b.StartIn + b.Duration
}

func (b BaseEvent) CollectionName() string {
	return "ServerSpecialEventData"
}

func (b *BaseEvent) SetEventId(eventId int) {
	b.EventId = eventId
}
