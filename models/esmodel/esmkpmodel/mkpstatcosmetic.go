package esmkpmodel

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/models/esmodel"
	"github.com/WolffunService/thetan-shared-common/models/esmodel/esinfomodel"
)

type CosmeticMapping struct {
	ID             string                          `json:"id"`
	User           esmodel.UserModel               `json:"user"`
	MkpStat        esinfomodel.MKPStatModel        `json:"mkp_stat"`
	BlockchainInfo esinfomodel.BlockchainInfoModel `json:"blockchain_info"`
	CosmeticInfo   esinfomodel.CosmeticInfoModel   `json:"cosmetic_info"`
	Buyer          esinfomodel.BuyerInfoModel      `json:"buyer"`
	Timestamp      time.Time                       `json:"@timestamp"`
}

func (CosmeticMapping) Index() string {
	return "mkpstats-cosmetic"
}
