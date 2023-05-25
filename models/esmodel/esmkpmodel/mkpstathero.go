package esmkpmodel

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/models/esmodel"
	"github.com/WolffunService/thetan-shared-common/models/esmodel/esinfomodel"
)

type HeroMapping struct {
	ID             string                          `json:"id"`
	User           esmodel.UserModel               `json:"user"`
	MkpStat        esinfomodel.MKPStatModel        `json:"mkp_stat"`
	BlockchainInfo esinfomodel.BlockchainInfoModel `json:"blockchain_info"`
	HeroInfo       esinfomodel.HeroInfoModel       `json:"hero_info"`
	Buyer          esinfomodel.BuyerInfoModel      `json:"buyer"`
	RentInfo       esinfomodel.RentInfoModel       `json:"rent_info"`
	Renter         esinfomodel.RenterInfoModel     `json:"renter"`
	PricePerBattle int64                           `json:"price_per_battle"`
	Timestamp      time.Time                       `json:"@timestamp"`
}

func (HeroMapping) Index() string {
	return "mkpstats-hero"
}
