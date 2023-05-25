package statistictypeenum

type StatisticType int

const (
	_Start StatisticType = iota
	THC_BOX
	THC_DEP
	THC_PRESALE

	THC_RENTAL

	THG_BOX
	THG_DEP
	THG_EVENT
	THG_AIRDROP
	THG_STAKE

	WBNB_HERO
	THC_HERO
	BUSD_BOX
	_End
)

var _text = map[StatisticType]string{
	THC_BOX:     "thcBox",
	THC_DEP:     "thcDep",
	THC_PRESALE: "thcPresale",

	THC_RENTAL: "thcRental",

	THG_BOX:     "thgBox",
	THG_DEP:     "thgDep",
	THG_EVENT:   "thgEvent",
	THG_AIRDROP: "thgAirdrop",
	THG_STAKE:   "thgStake",

	WBNB_HERO: "wbnbHero",
	THC_HERO:  "thcHero",
	BUSD_BOX:  "busdBox",
}

func (s StatisticType) IsValid() bool {
	return s > _Start && s < _End && _text[s] != ""
}

func (s StatisticType) String() string {
	return "statisticsCurrency." + _text[s]
}
