package esmkpmodel

import (
	"github.com/WolffunService/thetan-shared-common/models/esmodel/esinfomodel"
	"time"
)

type MkpLogCosmeticMapping struct {
	MkpLog         esinfomodel.MkpLogModel         `json:"mkp_log"`
	BlockchainInfo esinfomodel.BlockchainInfoModel `json:"blockchain_info"`
	CosmeticInfo   esinfomodel.CosmeticInfoModel   `json:"cosmetic_info"`
	Timestamp      time.Time                       `json:"@timestamp"`
}

func (MkpLogCosmeticMapping) Index() string {
	return "log-marketplace.action.cosmetic.v2"
}
