package pubsubenum

// TopicSubscriptions register all topics and subscriptions
var TopicSubscriptions = map[PubSubTopic][]PubSubSubscriber{
	TopicBoxOpenAudit:            {},
	TopicBoxPurchaseAudit:        {},
	TopicUpgradeHeroAudit:        {},
	TopicCurrencyChangeAudit:     {},
	TopicTrophyExpAudit:          {},
	TopicTransactionSuccessAudit: {},
	TopicHeroCreateAudit:         {},
	TopicBoxEventAudit:           {},
	TopicReferRewardAudit:        {},
	TopicBattleEndAudit:          {},
	TopicHeroRentalAudit:         {},
	TopicAdminAudit:              {},
	TopicThcConvertAudit:         {},

	Topic0: {},
	Topic1: {},

	TopicStakingStake: {
		SubscriberStakingStake,
	},
	TopicStakingUnstake: {
		SubscriberStakingUnstake,
	},
	TopicStakingTimeout: {
		SubscriberStakingTimeout,
	},

	TopicFusion: {
		SubscriberFusionSuccess,
		SubscriberFusionSuccessES,
	},
	TopicClaimToken: {
		SubscriberClaimTokenSuccess,
		SubscriberClaimTokenSuccessES,
	},
	TopicThetanBoxPaidV2: {
		SubscriberThetanBoxPaidSuccess,
		SubscriberThetanBoxPaidSuccessES,
	},
	TopicBuyPrivateSale: {
		SubscriberBuyPrivateSaleSuccess,
		SubscriberBuyPrivateSaleSuccessES,
	},

	TopicHeroSellingSuccess: {SubscriberHeroSellingSuccess},
	TopicHeroSellingCancel:  {SubscriberHeroSellingCancel},
	TopicHeroSellingTimeout: {SubscriberHeroSellingTimeout},
	TopicHeroBuyingSuccess:  {SubscriberHeroBuyingSuccess},

	TopicHeroMintSuccess:      {SubscriberHeroMintSuccess},
	TopicHeroMintTimeout:      {SubscriberHeroMintTimeout},
	TopicHeroMultiMintSuccess: {SubscriberHeroAdminMultiMintSuccess},

	TopicHeroLendingSuccess: {SubscriberHeroLendingSuccess},
	TopicHeroLendingCancel:  {SubscriberHeroLendingCancel},
	TopicHeroLendingTimeout: {SubscriberHeroLendingTimeout},
	TopicHeroRentingSuccess: {SubscriberHeroRentingSuccess},

	TopicHeroUnlockSuccess: {SubscriberHeroUnlockSuccess},
	TopicHeroUnlockTimeout: {SubscriberHeroUnlockTimeout},

	TopicCosmeticSellingSuccess: {SubscriberCosmeticSellingSuccess},
	TopicCosmeticSellingCancel:  {SubscriberCosmeticSellingCancel},
	TopicCosmeticSellingTimeout: {SubscriberCosmeticSellingTimeout},
	TopicCosmeticBuyingSuccess:  {SubscriberCosmeticBuyingSuccess},

	TopicCosmeticMintSuccess:      {SubscriberCosmeticMintSuccess},
	TopicCosmeticMintTimeout:      {SubscriberCosmeticMintTimeout},
	TopicCosmeticMultiMintSuccess: {SubscriberCosmeticMultiMintSuccess},

	TopicInternalBuySellRent: {SubscriberInternalBuySellRentES},

	TopicSendEmailMarketing:       {SubscriberSendEmailMarketing},
	TopicNotifySystemAlert:        {SubscriberNotifySystemAlert},
	TopicUpdateMongoConfig:        {SubscriberUpdateMongoConfig},
	TopicBattleLogs:               {SubscriberBattleLogs},
	TopicRivalBattleLogs:          {SubscriberRivalBattleLogs},
	TopicRivalAnalytics0:          {SubscriberRivalAnalytics0},
	TopicMatchFoundThetanArena:    {},
	TopicDeleteTicketThetanArena:  {},
	TopicMatchFoundThetanRivals:   {SubscriberMatchFoundThetanRivals},
	TopicDeleteTicketThetanRivals: {SubscriberDeleteTicketThetanRivals},

	TopicUpdateBotBattleEnd: {SubscriberUpdateBotBattleEnd},
	TopicUpdateBotRanking:   {SubscriberUpdateRanking},

	TopicFriendRequest: {SubFriendRequest},
	TopicFriendAccept:  {SubFriendAccept},
	TopicFriendDecline: {SubFriendDecline},
}
