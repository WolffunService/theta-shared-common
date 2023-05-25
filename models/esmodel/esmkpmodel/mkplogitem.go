package esmkpmodel

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/models/esmodel/esinfomodel"
)

type MkpLogItemMapping struct {
	MkpLog    esinfomodel.MkpLogModel `json:"mkp_log"`
	Timestamp time.Time               `json:"@timestamp"`
}

func (MkpLogItemMapping) Index() string {
	return "log-marketplace.action.item.v2"
}
