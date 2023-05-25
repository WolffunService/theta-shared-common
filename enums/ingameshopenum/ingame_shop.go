package ingameshopenum

type IngameShopSource int

const (
	ISS_BuyPowerPoint IngameShopSource = iota + 1
	ISS_BuyBoxNonNft
	ISS_StartPack
	ISS_SpecialOffer
)
