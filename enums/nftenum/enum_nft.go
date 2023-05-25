package nftenum

type RefType int

const (
	RT_HERO     RefType = 0
	RT_PET      RefType = 1
	RT_COSMETIC RefType = 2
)

type NFTStatus int

const (
	NFTS_AVAILABLE NFTStatus = 1
	NFTS_SELLING   NFTStatus = 2
	NFTS_LENDING   NFTStatus = 3
	NFTS_RENTED    NFTStatus = 4
)
