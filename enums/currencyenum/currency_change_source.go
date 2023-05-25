package currencyenum

// Currency
type CurrencySource int

const (
	CS_NONE    CurrencySource = 0
	CS_EVENT   CurrencySource = 1
	CS_DEPOSIT CurrencySource = 2
	CS_AIRDROP CurrencySource = 3
)

// Cô dì, chú bác, anh, em hết sức lưu ý: CurrencyChangeSource không quá 999,
//để tránh bị trùng với Thetan Rivals Inventory ChangeSource <3

type CurrencyChangeSource int

const (
	CCS_UPGRADE_HERO              CurrencyChangeSource = 1
	CCS_BATTLE_END                CurrencyChangeSource = 2
	CCS_IAP_SHOP                  CurrencyChangeSource = 3
	CCS_LEVEL_UP                  CurrencyChangeSource = 4
	CCS_RANKING_REWARD            CurrencyChangeSource = 5
	CCS_SEASON_END_REWARD         CurrencyChangeSource = 6
	CCS_SHOP_NON_NFT              CurrencyChangeSource = 7
	CCS_SHOP_POWER_POINT          CurrencyChangeSource = 8
	CCS_SHOP_INGAME_THG           CurrencyChangeSource = 9
	CCS_REFERRAL_REWARD           CurrencyChangeSource = 10
	CCS_SHOP_BOX_NON_NFT          CurrencyChangeSource = 11
	CCS_SPECIAL_EVENT_REWARD      CurrencyChangeSource = 12
	CCS_SPECIAL_EVENT_FEE_JOIN    CurrencyChangeSource = 13
	CCS_BUY_SKILL_TURN            CurrencyChangeSource = 14
	CCS_PLAY_AGAIN_BONUS          CurrencyChangeSource = 16
	CCS_FUSION_CASHBACK           CurrencyChangeSource = 17
	CCS_FUSION_LEADERBOARD        CurrencyChangeSource = 18
	CCS_FUSION_REFRESH_OUTPUT     CurrencyChangeSource = 19
	CCS_REFERRAL_REWARD_MKP       CurrencyChangeSource = 20
	CCS_CHANGE_NAME_PROFILE       CurrencyChangeSource = 21
	CCS_TRADING_EVENT_CASHBACK    CurrencyChangeSource = 22
	CCS_TRADING_EVENT_LEADERBOARD CurrencyChangeSource = 23
	CCS_TRADING_EVENT_DAILY       CurrencyChangeSource = 24
	CCS_TRADING_REFRESH_DAILY     CurrencyChangeSource = 25

	CCS_DEPOSIT         CurrencyChangeSource = 50
	CCS_CLAIM           CurrencyChangeSource = 51
	CCS_CLAIM_FEE       CurrencyChangeSource = 52
	CCS_VESTING_CLAIM   CurrencyChangeSource = 53
	CCS_CONVERT_VESTING CurrencyChangeSource = 54

	CCS_CREATOR_PROGRAM_VIEWER  CurrencyChangeSource = 55
	CCS_CREATOR_PROGRAM_CREATOR CurrencyChangeSource = 56

	CCS_STAKING_REWARD CurrencyChangeSource = 60
	CCS_QUEST_REWARD   CurrencyChangeSource = 61
	CCS_OPEN_GIFT_BOX  CurrencyChangeSource = 62

	CCS_MKP_QUEST_STAGE_REWARD  CurrencyChangeSource = 63
	CCS_MKP_QUEST_DAILY_CHECKIN CurrencyChangeSource = 64
	CCS_MKP_QUEST_DAILY_STREAK  CurrencyChangeSource = 65
	CCS_MKP_QUEST_LB_REWARD     CurrencyChangeSource = 66
	CCS_MKP_QUEST_BOX_BUY       CurrencyChangeSource = 67
	CCS_MKP_QUEST_BOX_OPEN      CurrencyChangeSource = 68
	CCS_MKP_QUEST_REFRESH       CurrencyChangeSource = 69

	CCS_BUY_PRIVATE_SALE    CurrencyChangeSource = 70
	CCS_REFUND_PRIVATE_SALE CurrencyChangeSource = 71

	CCS_ADMIN        CurrencyChangeSource = 101
	CCS_ALPHA_REWARD CurrencyChangeSource = 102
	CCS_BETA_REWARD  CurrencyChangeSource = 103
	CCS_ADMIN_REVERT CurrencyChangeSource = 104
	CCS_REFUND       CurrencyChangeSource = 105
)

func CCSToString(c CurrencyChangeSource) string {
	switch c {
	case CCS_UPGRADE_HERO: // 1
		return "Upgrade hero"
	case CCS_BATTLE_END: // 2
		return "Battle end"
	case CCS_IAP_SHOP: // 3
		return "IAP shop"
	case CCS_LEVEL_UP: // 4
		return "Level up"
	case CCS_RANKING_REWARD: // 5
		return "Ranking reward"
	case CCS_SEASON_END_REWARD: // 6
		return "Season end reward"
	case CCS_SHOP_NON_NFT: // 7
		return "Shop non-nft"
	case CCS_SHOP_POWER_POINT: // 8
		return "Shop powerpoint"
	case CCS_SHOP_INGAME_THG: // 9
		return "Shop THG"
	case CCS_REFERRAL_REWARD: // 10
		return "Referal reward"
	case CCS_SHOP_BOX_NON_NFT: // 11
		return "Box non-NFT"
	case CCS_SPECIAL_EVENT_REWARD: // 12
		return "Special event reward"

	case CCS_DEPOSIT: // 50
		return "Deposit"
	case CCS_CLAIM: // 51
		return "Claim"
	case CCS_CLAIM_FEE: // 52
		return "Claim fee"
	case CCS_VESTING_CLAIM: // 53
		return "Vesting claim"

	case CCS_OPEN_GIFT_BOX: // 62
		return "Gift box reward"

	case CCS_ADMIN: // 101
		return "Admin"
	case CCS_ALPHA_REWARD: // 102
		return "Alpha reward"
	case CCS_BETA_REWARD: // 103
		return "Beta reward"
	case CCS_ADMIN_REVERT: // 104
		return "Admin"

	case CCS_CREATOR_PROGRAM_VIEWER: // 150
		return "Creator viewer"
	case CCS_CREATOR_PROGRAM_CREATOR: // 151
		return "Creator reward"
	}

	return "Untrack yet"
}
