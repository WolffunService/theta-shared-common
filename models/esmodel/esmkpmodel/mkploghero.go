package esmkpmodel

import (
	"github.com/WolffunService/theta-shared-common/models/esmodel/esinfomodel"
	"time"
)

type MkpLogHeroMapping struct {
	MkpLog         esinfomodel.MkpLogModel         `json:"mkp_log"`
	BlockchainInfo esinfomodel.BlockchainInfoModel `json:"blockchain_info"`
	HeroInfo       esinfomodel.HeroInfoModel       `json:"hero_info"`
	RentInfo       esinfomodel.RentInfoModel       `json:"rent_info"`
	Timestamp      time.Time                       `json:"@timestamp"`
}

func (MkpLogHeroMapping) Index() string {
	return "log-marketplace.action.hero.v2"
}
