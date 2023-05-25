package smcmodel

type BlockchainInfo struct {
	// ---- hero
	HeroSmcAddress       string `json:"heroSmcAddress" bson:"heroSmcAddress"`
	HeroMinterSmcAddress string `json:"heroMinterSmcAddress" bson:"heroMinterSmcAddress"`
	HeroRentalSmcAddress string `json:"heroRentalSmcAddress" bson:"heroRentalSmcAddress"`

	// ---- cosmetic
	CosmeticSmcAddress       string `json:"cosmeticSmcAddress" bson:"cosmeticSmcAddress"`
	CosmeticMinterSmcAddress string `json:"cosmeticMinterSmcAddress" bson:"cosmeticMinterSmcAddress"`

	// ---- buy/ sell
	MkpV3SmcAddress string `json:"mkpV3SmcAddress" bson:"mkpV3SmcAddress"`
}

func (BlockchainInfo) CollName() string {
	return "BlockchainInfos"
}
