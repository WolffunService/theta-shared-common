package esmkpmodel

import (
	"time"

	"github.com/WolffunService/theta-shared-common/models/esmodel/esinfomodel"
)

type MkpLogItemFirstDataMapping struct {
	MkpLog    esinfomodel.MkpLogModel `json:"mkp_log"`
	Timestamp time.Time               `json:"@timestamp"`
}

func (MkpLogItemFirstDataMapping) Index() string {
	return "log-marketplace.action.item.v2.firstdata"
}
