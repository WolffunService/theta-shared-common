package mkpenum

type MKPRefType int

const (
	MKP_REF_TYPE_HERO          MKPRefType = 0
	MKP_REF_TYPE_PET           MKPRefType = 1
	MKP_REF_TYPE_COSMETIC      MKPRefType = 2
	MKP_REF_TYPE_BOX           MKPRefType = 3
	MKP_REF_TYPE_FUSION        MKPRefType = 4
	MKP_REF_TYPE_PRIVATE_SALE  MKPRefType = 5
	MKP_REF_TYPE_TICKET_BATTLE MKPRefType = 6
	MKP_REF_TYPE_THC           MKPRefType = 7
	MKP_REF_TYPE_THG           MKPRefType = 8
	MKP_REF_TYPE_GTHC          MKPRefType = 9
	MKP_REF_TYPE_GTHG          MKPRefType = 10
	MKP_REF_TYPE_NONE          MKPRefType = 99
)

type MarketplaceAction int

const (
	MKP_ACTION_NONE MarketplaceAction = 1

	MKP_ACTION_SELLING         MarketplaceAction = 11
	MKP_ACTION_BUYING          MarketplaceAction = 12
	MKP_ACTION_CANCEL_SELLING  MarketplaceAction = 13
	MKP_ACTION_SELLING_TIMEOUT MarketplaceAction = 14

	MKP_ACTION_LENDING         MarketplaceAction = 21
	MKP_ACTION_RENTING         MarketplaceAction = 22
	MKP_ACTION_CANCEL_LENDING  MarketplaceAction = 23
	MKP_ACTION_LENDING_TIMEOUT MarketplaceAction = 24
	MKP_ACTION_UNLOCK          MarketplaceAction = 25
	MKP_ACTION_RENTED_RELEASED MarketplaceAction = 26

	MKP_ACTION_MINTED       MarketplaceAction = 31
	MKP_ACTION_MINT_TIMEOUT MarketplaceAction = 32
	MKP_ACTION_USER_MINTED  MarketplaceAction = 33
	MKP_ACTION_OPEN_BOX     MarketplaceAction = 34

	MKP_ACTION_FUSION   MarketplaceAction = 51
	MKP_ACTION_WITHDRAW MarketplaceAction = 52
	MKP_ACTION_UPGRADE  MarketplaceAction = 53
	MKP_ACTION_DEPOSIT  MarketplaceAction = 54

	MKP_ACTION_BUY_BOX_THC MarketplaceAction = 60
	MKP_ACTION_BUY_BOX_THG MarketplaceAction = 61

	MKP_ACTION_BUY_PRIVATE_SALE      MarketplaceAction = 70
	MKP_ACTION_BUY_TOURNAMENT_TICKET MarketplaceAction = 71

	MKP_ACTION_SYNC_DATA_OWNER  MarketplaceAction = 80
	MKP_ACTION_SYNC_DATA_STATUS MarketplaceAction = 81

	MKP_ACTION_SEASON_END     MarketplaceAction = 85
	MKP_ACTION_RANKING_REWARD MarketplaceAction = 86
	MKP_ACTION_ADMIN          MarketplaceAction = 87
)

type ItemType int

const (
	CONST_ITEM_TYPE_BANNER ItemType = 1
)

type ItemStatus int

const (
	CONST_ITEM_STATUS_DISPLAY      ItemStatus = 1
	CONST_ITEM_STATUS_NONE_DISPLAY ItemStatus = 2
)

type ESGroup int

const (
	CONST_ES_GROUP_COUNT_TIME_24H    ESGroup = 1
	CONST_ES_GROUP_COUNT_TIME_7DAYS  ESGroup = 2
	CONST_ES_GROUP_COUNT_TIME_30DAYS ESGroup = 3
	CONST_ES_GROUP_COUNT_TIME_ALL    ESGroup = 4
)

type ItemFormat int

const (
	CONST_ITEM_FORMAT_FIXED  ItemFormat = 1
	CONST_ITEM_FORMAT_CONFIG ItemFormat = 2
)
