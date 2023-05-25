package auditmodel

import (
	"encoding/json"
	"log"
	"time"

	"github.com/WolffunService/thetan-shared-common/enums/adminenum"
	"github.com/WolffunService/thetan-shared-common/models/thetanboxmodel"
	"github.com/WolffunService/thetan-shared-common/proto/auditproto"
)

type AdminAuditCreator struct{}

type AdminAuditEvent struct {
	AdminId      string            `bson:"adminId" json:"adminId"`
	Ip           string            `bson:"ip" json:"ip"`
	Email        string            `bson:"email" json:"email"`
	Role         int32             `bson:"role" json:"role"`
	FunctionType int32             `bson:"functionType" json:"functionType"`
	Timestamp    int64             `bson:"timestamp" json:"timestamp"`
	SendBoxEvent *AdminSendBoxData `bson:"sendBoxEvent,omitempty" json:"sendBoxEvent,omitempty"`
}

type AdminSendBoxData struct {
	UserAddress string                   `bson:"userAddress" json:"userAddress"`
	UserEmail   string                   `bson:"userEmail" json:"userEmail"`
	Boxes       []thetanboxmodel.BoxData `bson:"boxes" json:"boxes"`
	Date        string                   `json:"date" bson:"date"`
}

func (AdminAuditCreator) NewSendBoxEvent(res *auditproto.AdminEvent) *AdminAuditEvent {
	if res.Function != adminenum.AFT_SEND_BOX {
		log.Println("[error] - cannot create NewSendBoxEvent because this function (id: ", res.Function, ") is not AdminSendBox")
		return nil
	}

	var sendBoxData AdminSendBoxData
	if err := json.Unmarshal(res.JsonData, &sendBoxData); err != nil {
		log.Println("[error] -", err)
		return nil
	}

	sendBoxData.Date = time.Unix(res.Timestamp, 0).UTC().Format("2006-01-02")

	return &AdminAuditEvent{
		AdminId:      res.AdminId,
		Ip:           res.IpAddress,
		Email:        res.Email,
		Role:         res.Role,
		FunctionType: res.Function,
		SendBoxEvent: &sendBoxData,
		Timestamp:    res.Timestamp,
	}
}
