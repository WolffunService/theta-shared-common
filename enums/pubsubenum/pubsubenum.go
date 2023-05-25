package pubsubenum

type (
	PubSubTopic      string
	PubSubSubscriber string
)

const (
	TopicBoxOpenAudit            PubSubTopic = "box-open-audit"
	TopicBoxPurchaseAudit        PubSubTopic = "box-purchase-audit"
	TopicUpgradeHeroAudit        PubSubTopic = "upgrade-hero-audit"
	TopicCurrencyChangeAudit     PubSubTopic = "currency-change-audit"
	TopicTrophyExpAudit          PubSubTopic = "trophy-exp-audit"
	TopicTransactionSuccessAudit PubSubTopic = "transaction-success-audit"
	TopicHeroCreateAudit         PubSubTopic = "hero-create-audit"
	TopicBoxEventAudit           PubSubTopic = "box-event-audit"
	TopicReferRewardAudit        PubSubTopic = "refer-reward-audit"
	TopicBattleEndAudit          PubSubTopic = "battle-end-audit"
	TopicHeroRentalAudit         PubSubTopic = "hero-rental-audit"
	TopicAdminAudit              PubSubTopic = "admin-api-audit"
	TopicThcConvertAudit         PubSubTopic = "thc-convert-audit"

	Topic0 PubSubTopic = "topic-0"
	Topic1 PubSubTopic = "topic-1"

	TopicBotBattleEnd PubSubTopic = "bot-battle-end"

	TopicStakingStake   PubSubTopic = "StakingTopic"
	TopicStakingUnstake PubSubTopic = "UnstakingTopic"
	TopicStakingTimeout PubSubTopic = "NotUseStakingTopic"

	TopicFusion          PubSubTopic = "FusionTopic"
	TopicClaimToken      PubSubTopic = "ClaimTokenTopic"
	TopicThetanBoxPaidV2 PubSubTopic = "ThetanboxPaidV2Topic"
	TopicBuyPrivateSale  PubSubTopic = "BuyPrivateSaleTopic"

	TopicHeroSellingSuccess PubSubTopic = "MKPSuccessfulSellingTopic"
	TopicHeroSellingCancel  PubSubTopic = "MKPCancelSellingTopic"
	TopicHeroSellingTimeout PubSubTopic = "MKPSellingTimeoutTopic"
	TopicHeroBuyingSuccess  PubSubTopic = "MKPSuccessfulBuyingTopic"

	TopicHeroMintSuccess      PubSubTopic = "TokenMintTopic"
	TopicHeroMintTimeout      PubSubTopic = "TokenMintTimeoutTopic_V1"
	TopicHeroMultiMintSuccess PubSubTopic = "AdminMultipleTokensMintTopic"

	TopicHeroLendingSuccess PubSubTopic = "SuccessfulLendingTopic"
	TopicHeroLendingCancel  PubSubTopic = "CancelLendingTopic"
	TopicHeroLendingTimeout PubSubTopic = "LendingTimeoutTopic"
	TopicHeroRentingSuccess PubSubTopic = "SuccessfulRentingTopic"

	TopicHeroUnlockSuccess PubSubTopic = "UnlockHeroTopic"
	TopicHeroUnlockTimeout PubSubTopic = "UnlockHeroTimeout"

	TopicCosmeticSellingSuccess PubSubTopic = "CosmeticMKPSuccessfulSellingTopic"
	TopicCosmeticSellingCancel  PubSubTopic = "CosmeticMKPCancelSellingTopic"
	TopicCosmeticSellingTimeout PubSubTopic = "CosmeticMKPSellingTimeoutTopic"
	TopicCosmeticBuyingSuccess  PubSubTopic = "CosmeticMKPSuccessfulBuyingTopic"

	TopicCosmeticMintSuccess      PubSubTopic = "cosmetic-minted"
	TopicCosmeticMintTimeout      PubSubTopic = "cosmetic-mint-timeout"
	TopicCosmeticMultiMintSuccess PubSubTopic = "cosmetic-admin-multi-minted"

	TopicInternalBuySellRent PubSubTopic = "InternalBuySellRentTopic"

	TopicSendEmailMarketing PubSubTopic = "send-email-marketing"
	TopicNotifySystemAlert  PubSubTopic = "system-alert"
	TopicUpdateMongoConfig  PubSubTopic = "remote-config-update"

	TopicConvert    PubSubTopic = "convert"
	TopicBattleLogs PubSubTopic = "BATTLE_LOGS"

	TopicRivalAnalytics0 PubSubTopic = "rival-analytics-0"
	TopicRivalBattleLogs PubSubTopic = "rival-battle-logs"

	TopicMatchFoundThetanArena    PubSubTopic = "match-found"
	TopicMatchFoundThetanRivals   PubSubTopic = "rival-match-found"
	TopicDeleteTicketThetanArena  PubSubTopic = "delete-ticket"
	TopicDeleteTicketThetanRivals PubSubTopic = "rival-delete-ticket"

	TopicUpdateBotBattleEnd PubSubTopic = "update-bot-battle-end"
	TopicUpdateBotRanking   PubSubTopic = "update-bot-ranking"

	TopicFriendRequest PubSubTopic = "friend-request"
	TopicFriendAccept  PubSubTopic = "friend-accept"
	TopicFriendDecline PubSubTopic = "friend-decline"
)

const (
	SubscriberStakingStake   PubSubSubscriber = "StakingTopic-sub"
	SubscriberStakingUnstake PubSubSubscriber = "UnstakingTopic-sub"
	SubscriberStakingTimeout PubSubSubscriber = "NotUseStakingTopic-sub"

	SubscriberCosmeticSellingStop PubSubSubscriber = "CosmeticIgnoreSignatureTopic-sub"
	SubscriberCosmeticMinted      PubSubSubscriber = "CosmeticMintedTopic-sub"
	SubscriberCosmeticBuy         PubSubSubscriber = "CosmeticMatchTransactionTopic-sub"

	SubscriberFusionSuccess   PubSubSubscriber = "FusionTopic-sub"
	SubscriberFusionSuccessES PubSubSubscriber = "FusionTopic-ES-sub"
	SubscriberFusionTimeout   PubSubSubscriber = "FusionTimeoutTopic-sub"

	SubscriberClaimTokenSuccess   PubSubSubscriber = "ClaimTokenTopic-sub"
	SubscriberClaimTokenSuccessES PubSubSubscriber = "ClaimTokenTopic-ES-sub"
	SubscriberClaimTokenTimeout   PubSubSubscriber = "ClaimTokenTimeoutTopic-sub"

	SubscriberThetanBoxPaidSuccess   PubSubSubscriber = "ThetanboxPaidV2Topic-sub"
	SubscriberThetanBoxPaidSuccessES PubSubSubscriber = "ThetanboxPaidV2Topic-ES-sub"
	SubscriberThetanBoxPaidTimeout   PubSubSubscriber = "ThetanboxPaidV2TimeoutTopic-sub"

	SubscriberTEBoxPaid  PubSubSubscriber = "ThetanBoxPaidV2Topic-TE-sub"
	SubscriberTEHeroBuy  PubSubSubscriber = "MKPSuccessfulBuyingTopic-TE-sub"
	SubscriberTEHeroRent PubSubSubscriber = "MKPSuccessfulRentTopic-TE-sub"

	SubscriberBuyPrivateSaleSuccess   PubSubSubscriber = "BuyPrivateSaleTopic-sub"
	SubscriberBuyPrivateSaleSuccessES PubSubSubscriber = "BuyPrivateSaleTopic-ES-sub"
	SubscriberBuyPrivateSaleTimeout   PubSubSubscriber = "BuyPrivateSaleTimeoutTopic-sub"

	SubscriberHeroSellingSuccess PubSubSubscriber = "MKPSuccessfulSellingTopic-sub"
	SubscriberHeroSellingCancel  PubSubSubscriber = "MKPCancelSellingTopic-sub"
	SubscriberHeroSellingTimeout PubSubSubscriber = "MKPSellingTimeoutTopic-sub"
	SubscriberHeroBuyingSuccess  PubSubSubscriber = "MKPSuccessfulBuyingTopic-sub"

	SubscriberHeroMintSuccess           PubSubSubscriber = "TokenMintTopic-sub"
	SubscriberHeroAdminMultiMintSuccess PubSubSubscriber = "AdminMultipleTokensMintTopic-sub"
	SubscriberHeroMintTimeout           PubSubSubscriber = "TokenMintTimeoutTopic_V1-sub"

	SubscriberHeroLendingSuccess PubSubSubscriber = "SuccessfulLendingTopic-sub"
	SubscriberHeroLendingCancel  PubSubSubscriber = "CancelLendingTopic-sub"
	SubscriberHeroLendingTimeout PubSubSubscriber = "LendingTimeoutTopic-sub"

	SubscriberHeroRentingSuccess PubSubSubscriber = "SuccessfulRentingTopic-sub"

	SubscriberHeroUnlockSuccess PubSubSubscriber = "UnlockHeroTopic-sub"
	SubscriberHeroUnlockTimeout PubSubSubscriber = "UnlockHeroTimeout-sub"

	SubscriberCosmeticSellingSuccess PubSubSubscriber = "CosmeticMKPSuccessfulSellingTopic-sub"
	SubscriberCosmeticSellingCancel  PubSubSubscriber = "CosmeticMKPCancelSellingTopic-sub"
	SubscriberCosmeticSellingTimeout PubSubSubscriber = "CosmeticMKPSellingTimeoutTopic-sub"

	SubscriberCosmeticBuyingSuccess PubSubSubscriber = "CosmeticMKPSuccessfulBuyingTopic-sub"

	SubscriberCosmeticMintSuccess      PubSubSubscriber = "cosmetic-minted-sub"
	SubscriberCosmeticMintTimeout      PubSubSubscriber = "cosmetic-mint-timeout-sub"
	SubscriberCosmeticMultiMintSuccess PubSubSubscriber = "cosmetic-admin-multi-minted-sub"

	//SubscriberInHeroSellingSuccessES               PubSubSubscriber = "InHeroSellingSuccessES-sub"
	//SubscriberInHeroSellingCancelES                PubSubSubscriber = "InHeroCancelSellingES-sub"
	//SubscriberInHeroBuyingSuccessES                PubSubSubscriber = "InHeroBuyingSuccessES-sub"
	//SubscriberInHeroMintSuccessES                  PubSubSubscriber = "InHeroMintSuccessES-sub"
	//SubscriberInHeroAdminMultipleMintSuccessES     PubSubSubscriber = "InHeroAdminMultipleMintSuccessES-sub"
	//SubscriberInHeroLendingSuccessES               PubSubSubscriber = "InHeroLendingSuccessES-sub"
	//SubscriberInHeroLendingCancelES                PubSubSubscriber = "InHeroCancelLendingES-sub"
	//SubscriberInHeroRentingSuccessES               PubSubSubscriber = "InHeroRentingSuccessES-sub"
	//SubscriberInHeroUnlockSuccessES                PubSubSubscriber = "InHeroUnlockSuccessES-sub"
	//SubscriberInCosmeticSellingSuccessES           PubSubSubscriber = "InCosmeticSellingSuccessES-sub"
	//SubscriberInCosmeticSellingCancelES            PubSubSubscriber = "InCosmeticCancelSellingES-sub"
	//SubscriberInCosmeticBuyingSuccessES            PubSubSubscriber = "InCosmeticBuyingSuccessES-sub"
	//SubscriberInCosmeticMintSuccessES              PubSubSubscriber = "InCosmeticMintSuccessES-sub"
	//SubscriberInCosmeticAdminMultipleMintSuccessES PubSubSubscriber = "InCosmeticAdminMultipleMintSuccessES-sub"
	//SubscriberInHeroRentedReleasedES               PubSubSubscriber = "InHeroRentedReleasedES-sub"
	SubscriberInternalBuySellRentES PubSubSubscriber = "InternalBuySellRentES-sub"

	SubscriberSendEmailMarketing PubSubSubscriber = "send-email-marketing-sub"
	SubscriberNotifySystemAlert  PubSubSubscriber = "system-alert-sub"
	SubscriberUpdateMongoConfig  PubSubSubscriber = "mongo-config-thetan-support-sub"

	SubscriberIngameDeposit PubSubSubscriber = "ingame-deposit-sub"
	SubscriberBattleLogs    PubSubSubscriber = "BATTLE_LOGS-sub"

	SubscriberRivalAnalytics0 PubSubSubscriber = "rival-analytics-0-sub"
	SubscriberRivalBattleLogs PubSubSubscriber = "rival-battle-logs-sub"

	SubscriberMatchFoundThetanRivals   PubSubSubscriber = "rival-match-found-sub"
	SubscriberDeleteTicketThetanRivals PubSubSubscriber = "rival-delete-ticket-sub"

	SubscriberUpdateBotBattleEnd PubSubSubscriber = "update-bot-battle-end-sub"
	SubscriberUpdateRanking      PubSubSubscriber = "update-bot-ranking-sub"

	SubFriendRequest PubSubSubscriber = "friend-request-sub"
	SubFriendAccept  PubSubSubscriber = "friend-accept-sub"
	SubFriendDecline PubSubSubscriber = "friend-decline-sub"
)

func (topic PubSubTopic) String() string {
	return string(topic)
}

func (subscriber PubSubSubscriber) String() string {
	return string(subscriber)
}
