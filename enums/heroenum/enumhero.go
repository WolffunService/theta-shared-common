package heroenum

// for audits
type HeroEventSource int32 // heroEventSource (not HeroEventSource)

const (
	HES_Open_Box HeroEventSource = 1
	// HES_List            HeroEventSource = 2
	HES_Sold         HeroEventSource = 3
	HES_Purchase     HeroEventSource = 4
	HES_Non_NFT_Shop HeroEventSource = 5
	// HES_Mint            HeroEventSource = 6
	HES_New_User_Reward HeroEventSource = 7
	HES_First_Free      HeroEventSource = 8
	HES_Ranking         HeroEventSource = 9
	// HES_Unlist          HeroEventSource = 10
	HES_FUSION       HeroEventSource = 11
	HES_Wheel_Reward HeroEventSource = 12

	HES_ADMIN         HeroEventSource = 101
	HES_BETA          HeroEventSource = 102
	HES_GG_IOS_REVIEW HeroEventSource = 103
	HES_PARTNER       HeroEventSource = 104
)

type HeroEventStatus int32

const (
	HESt_Processing HeroEventStatus = 1
	HESt_Succeeded  HeroEventStatus = 2
	HESt_Failed     HeroEventStatus = 3
)

type RentState int

const (
	RS_OWNER_FOR_RENT      RentState = 1
	RS_OWNER_STOP_FOR_RENT RentState = 2
	RS_OWNER_RENTED        RentState = 3
	RS_OWNER_RETURN        RentState = 4
	RS_RENTER_RENTED       RentState = 5
	RS_RENTER_RETURN       RentState = 6
)

// --- HERO STATUS
type HeroStatus int

const (
	HS_DEFAULT                HeroStatus = 0
	HS_NOT_MINT               HeroStatus = 1
	HS_MINTING                HeroStatus = 2
	HS_AVAILABLE              HeroStatus = 3
	HS_RENTED                 HeroStatus = 4
	HS_FUSING                 HeroStatus = 5
	HS_RENTED_RELEASED        HeroStatus = 6
	HS_USER_NOT_MINT          HeroStatus = 7
	HS_USER_MINTING           HeroStatus = 8
	HS_ON_MARKET              HeroStatus = 10
	HS_FOR_RENT               HeroStatus = 11
	HS_IN_TRANSACTION_PROCESS HeroStatus = 20
	HS_IN_STAKING             HeroStatus = 30
	HS_LOCK                   HeroStatus = 99
	HS_IS_PREPARE_SELLING     HeroStatus = 110
	HS_IS_PREPARE_LENDING     HeroStatus = 111
	HS_IS_PREPARE_UNLOCK      HeroStatus = 112
)

func (status HeroStatus) IsInProcess() bool {
	return status == HS_IS_PREPARE_LENDING ||
		status == HS_IS_PREPARE_SELLING
}

func (status HeroStatus) IsActive() bool {
	return IsHeroActive(status)
}

func (status HeroStatus) CanUpgrade() bool {
	return IsHeroCanUpgrade(status)
}

func (status HeroStatus) CanSelling() bool {
	return status == HS_AVAILABLE
}

func (status HeroStatus) CanMintByAdmin() bool {
	return status == HS_NOT_MINT
}

func (status HeroStatus) CanMintByUser() bool {
	return status == HS_USER_NOT_MINT
}

func (status HeroStatus) CanLending() bool {
	return status == HS_AVAILABLE
}

func (status HeroStatus) CanUnlock() bool {
	return status == HS_RENTED_RELEASED
}

func (status HeroStatus) IsHeroNFT() bool {
	return status != HS_DEFAULT
}

func IsHeroActive(status HeroStatus) bool {
	return status != HS_ON_MARKET &&
		status != HS_IN_TRANSACTION_PROCESS &&
		status != HS_LOCK &&
		status != HS_FOR_RENT &&
		status != HS_IN_STAKING &&
		!status.IsInProcess()
}

func IsHeroCanUpgrade(status HeroStatus) bool {
	return status != HS_ON_MARKET &&
		status != HS_IN_TRANSACTION_PROCESS &&
		status != HS_LOCK &&
		status != HS_FOR_RENT &&
		status != HS_IN_STAKING &&
		!status.IsInProcess()
}

type HeroRole int

const (
	TANK     HeroRole = 0
	ASSASSIN HeroRole = 1
	MARKSMAN HeroRole = 2
)

type HeroFusionStatus int

const (
	HFS_AVAILABLE  HeroFusionStatus = 0 // default
	HFS_PROCESSING HeroFusionStatus = 1
)

type HeroRarity int // #enum: "Hero Rarity"

const (
	UNKNOWN HeroRarity = -1

	COMMON    HeroRarity = 0
	EPIC      HeroRarity = 1
	LEGENDARY HeroRarity = 2
)

type SkinRarity int // #enum: "Skin Rarity"

const (
	SR_s SkinRarity = iota - 1

	DEFAULT
	RARE
	MYTHIC

	SR_e
)
