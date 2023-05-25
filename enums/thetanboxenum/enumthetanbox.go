package thetanboxenum

import "github.com/WolffunService/thetan-shared-common/enums/currencyenum"

type BoxEventSource int

const (
	BES_PURCHASE BoxEventSource = 1
	BES_OPEN     BoxEventSource = 2
)

type BoxType int // enum:"BoxType"

const (
	COMMON_BOX                        BoxType = 1
	EPIC_BOX                          BoxType = 2
	LEGENDARY_BOX                     BoxType = 3
	SEASON_BOX                        BoxType = 4 // Christmas
	NEW_YEAR_BOX                      BoxType = 5
	SPECIAL_BOX                       BoxType = 6
	COSMETIC_BOX                      BoxType = 7
	CALA_BOX                          BoxType = 8  // 15/4/2022
	SUMMER_BOX_GIFT                   BoxType = 9  // 12/5/2022: ban kem octopus box bundle
	OCTOPUS_BUNDLE                    BoxType = 10 // 12/5/2022: ban thang 6
	WELCOME_BOX                       BoxType = 11
	GIFT_LINK_BOX                     BoxType = 12
	COINBASE_BOX                      BoxType = 13
	QUEST_BOX                         BoxType = 14
	DAILY_CHECKIN_MKP_QUEST           BoxType = 15 // danh cho daily checkin
	NEW_HOUSE_BOX                     BoxType = 16
	HALLOWEEN_BOX                     BoxType = 17
	EXCLU_UTIL_RAIDON                 BoxType = 18 // EXCLUSIVE UTILMATE RAIDON
	HATTRICK_BUNDLE                   BoxType = 19
	GLORY_BOX                         BoxType = 20
	SAMSUNG_BOX                       BoxType = 21
	GIFTMAS_WHEEL_CLAIM_TICKETS_QUEST BoxType = 22
	GIFTMAS_WHEEL_CLAIM_REWARDS_QUEST BoxType = 23
	BIRTHDAY_BOX                      BoxType = 24
	XMAS_BUNDLE                       BoxType = 25
	PREMIUM_LEGENDARY_BOX             BoxType = 26
)

func (b BoxType) ToCurrencyType() currencyenum.Currency {
	return currencyenum.Currency(b + 1000)
}

func (b BoxType) IsBoxFree() bool {
	return b == WELCOME_BOX ||
		b == GIFT_LINK_BOX ||
		b == COINBASE_BOX ||
		b == QUEST_BOX ||
		b == DAILY_CHECKIN_MKP_QUEST ||
		b == NEW_HOUSE_BOX ||
		b == HALLOWEEN_BOX ||
		b == GLORY_BOX ||
		b == SAMSUNG_BOX ||
		b == GIFTMAS_WHEEL_CLAIM_TICKETS_QUEST ||
		b == GIFTMAS_WHEEL_CLAIM_REWARDS_QUEST ||
		b == BIRTHDAY_BOX
}

func (b BoxType) IsBoxMustInstall() bool {
	return b == WELCOME_BOX ||
		b == GIFT_LINK_BOX ||
		b == NEW_HOUSE_BOX ||
		b == HALLOWEEN_BOX ||
		b == GLORY_BOX ||
		b == SAMSUNG_BOX ||
		b == GIFTMAS_WHEEL_CLAIM_TICKETS_QUEST ||
		b == GIFTMAS_WHEEL_CLAIM_REWARDS_QUEST ||
		b == BIRTHDAY_BOX
}

func (b BoxType) IsSpecialBox() bool {
	return b == OCTOPUS_BUNDLE ||
		b == SUMMER_BOX_GIFT ||
		b == CALA_BOX ||
		b == NEW_YEAR_BOX ||
		b == SEASON_BOX ||
		b == SPECIAL_BOX ||
		b == HATTRICK_BUNDLE ||
		b == EXCLU_UTIL_RAIDON ||
		b == XMAS_BUNDLE ||
		b == PREMIUM_LEGENDARY_BOX
}

type BoxStatus int

const (
	BOX_NEW                 BoxStatus = 0
	BOX_BUY_PROCESSING      BoxStatus = 1
	BOX_BUY_SUCCESS         BoxStatus = 2
	BOX_OPENING             BoxStatus = 3 //send api openbox from fe => save hero to thetanbox
	BOX_RECEIVING           BoxStatus = 4 //da send qua ben market place de chuan bi nhan
	BOX_REWARD_SUCCESS      BoxStatus = 5
	BOX_SEND_RABBIT_SUCCESS BoxStatus = 10
	BOX_SEND_RABBIT_ERROR   BoxStatus = 90
	BOX_CALL_BUY_FAILED     BoxStatus = 96 //call api buy to blockchain error
	BOX_RECEIVE_FAILED      BoxStatus = 97
	BOX_BUY_FAILED          BoxStatus = 98
	BOX_REWARD_FAILED       BoxStatus = 99
)

func (b BoxType) IsLimitBox() bool {
	return b == OCTOPUS_BUNDLE ||
		b == HATTRICK_BUNDLE ||
		b == BIRTHDAY_BOX ||
		b == XMAS_BUNDLE ||
		b == PREMIUM_LEGENDARY_BOX
}

func (b BoxType) IsFreeQuestBox() bool {
	return b == DAILY_CHECKIN_MKP_QUEST ||
		b == GIFTMAS_WHEEL_CLAIM_TICKETS_QUEST ||
		b == GIFTMAS_WHEEL_CLAIM_REWARDS_QUEST
}

func (b BoxType) IsBundleBox() bool {
	return b == HATTRICK_BUNDLE ||
		b == OCTOPUS_BUNDLE ||
		b == XMAS_BUNDLE ||
		b == PREMIUM_LEGENDARY_BOX
}
