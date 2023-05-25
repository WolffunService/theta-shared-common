package thetanerr

type ErrorCode int

func (e ErrorCode) Error() string {
	return ErrMessageByCode(e)
}

const (
	Success     ErrorCode = 0
	ErrInternal ErrorCode = 99

	ErrTokenInvalid          ErrorCode = 1
	ErrTokenExpired          ErrorCode = 2
	ErrUserNotFound          ErrorCode = 3
	ErrUserBanned            ErrorCode = 4
	ErrWrongFormat           ErrorCode = 5
	ErrWrongPrice            ErrorCode = 6
	ErrOutOfLimit            ErrorCode = 7
	ErrCannotClaim           ErrorCode = 8
	ErrNewVersionAvailable   ErrorCode = 9
	ErrUserBannedFindMatch   ErrorCode = 10
	ErrCantGetCurrencyPrice  ErrorCode = 11
	ErrExchangeRateDifferent ErrorCode = 12
	ErrMaintenanceCashOut    ErrorCode = 13
	ErrOutdatedSession       ErrorCode = 14
	ErrMismatchToken         ErrorCode = 15
	ErrBusyServer            ErrorCode = 16
	ErrClaimedToken          ErrorCode = 17
	ErrMaintainedServer      ErrorCode = 18
	ErrComingSoon            ErrorCode = 19

	// HTTP Status Code 400 - 599 (Client Error and Server Error)
	ErrHTTPNotFound            ErrorCode = 404
	ErrHTTPInternalServerError ErrorCode = 500

	// WolffunID Err_CODE 2001 -> 3000:
	ErrEmailNotExist             ErrorCode = 2001
	ErrEmailIsExist              ErrorCode = 2002
	ErrCodeNotValid              ErrorCode = 2003
	ErrEmailNotValid             ErrorCode = 2004
	ErrInCorrectUserId           ErrorCode = 2005
	ErrNotEnoughTicketChangeName ErrorCode = 2006
	ErrUserNameHasTaken          ErrorCode = 2007
	ErrTooManyAccountLogin       ErrorCode = 2008
	ErrAccountHasBeenLinked      ErrorCode = 2009

	//theta data code 3001 -> 4000
	ErrIdInvalid            ErrorCode = 3001
	ErrAddressInvalid       ErrorCode = 3002
	ErrTypeInvalid          ErrorCode = 3003
	ErrHeroIdInvalid        ErrorCode = 3004
	ErrInvalidate           ErrorCode = 3005
	ErrNotEnough            ErrorCode = 3006
	ErrNotEnoughTHC         ErrorCode = 3007
	ErrNotEnoughTHG         ErrorCode = 3008
	ErrNotEnoughPP          ErrorCode = 3009
	ErrNotOwnerHero         ErrorCode = 3010
	ErrHeroNotAvailable     ErrorCode = 3011
	ErrNotRankEnoughCashout ErrorCode = 3012
	ErrNotOldEnoughCashout  ErrorCode = 3013
	OutOfLimitCashout       ErrorCode = 3014
	ErrNotEnoughGTHG        ErrorCode = 3015
	ErrNotEnoughGTHC        ErrorCode = 3016
	ErrInvalidGTHG          ErrorCode = 3017
	ErrNotEnoughgFeeGTHG    ErrorCode = 3018
	ErrHeroInTheMatch       ErrorCode = 3019
	ErrMinimumCashout       ErrorCode = 3020
	ErrMaximumCashout       ErrorCode = 3021

	//box
	ErrOutdatedBuying ErrorCode = 3100
	ErrOutOfBoxes     ErrorCode = 3101
	ErrNotEnoughBox   ErrorCode = 3102

	//friend
	ErrOtherFullRequestReceived ErrorCode = 3103
	ErrOtherFullFriendList      ErrorCode = 3104
	ErrFullFriendList           ErrorCode = 3105
	ErrReceiptDataInValid       ErrorCode = 3106

	//report battle end
	ErrReachLimitReportInWeek ErrorCode = 3107
	ErrPreviouslyReported     ErrorCode = 3108

	// Box
	ErrBoxExchangeRate     ErrorCode = 3109
	ErrMustLinkedWolffunId ErrorCode = 3110
	ErrMustInstalledGame   ErrorCode = 3111
	ErrNeedMoreTimeToOpen  ErrorCode = 3112
	ErrAlreadyClaimFreeBox ErrorCode = 3113
	ErrMustConnectCoinbase ErrorCode = 3114
	ErrGiftCodeInvalid     ErrorCode = 3115
	ErrRequestProcessing   ErrorCode = 3116

	// Rental
	ErrRentedItem            ErrorCode = 3120
	ErrLowRentBattles        ErrorCode = 3121
	ErrHighRentBattles       ErrorCode = 3122
	ErrLowRentPrice          ErrorCode = 3123
	ErrItemIsForRent         ErrorCode = 3124
	ErrMaxBattleRented       ErrorCode = 3125
	ErrHighRentPrice         ErrorCode = 3126
	ErrReturningOwner        ErrorCode = 3127
	ErrMaintenanceHeroRental ErrorCode = 3128

	// VestingSafe
	ErrMaintenanceVestingSafe ErrorCode = 3140
	ErrAlreadyClaimed         ErrorCode = 3141

	// Creator Program
	ErrMaintenanceCreatorProgram ErrorCode = 3150
	ErrAreNotCreator             ErrorCode = 3151
	ErrAreNotViewer              ErrorCode = 3152

	// Staking
	ErrMaintenanceStaking  ErrorCode = 3160
	ErrInvalidPlan         ErrorCode = 3161
	ErrInvalidStakeAmount  ErrorCode = 3162
	ErrOutOfTimeStake      ErrorCode = 3163
	ErrStakePoolIsFull     ErrorCode = 3164
	ErrHeroUnavailable     ErrorCode = 3165
	ErrNotReadyToUnstake   ErrorCode = 3166
	ErrPreviousUnstake     ErrorCode = 3167
	ErrCooldownClaimReward ErrorCode = 3168
	ErrMaxPendingStake     ErrorCode = 3169
	ErrEmptyReward         ErrorCode = 3170
	ErrNotReadyForClaim    ErrorCode = 3171
	ErrMaxHeroesPerStake   ErrorCode = 3172

	// PrivateSale
	ErrPrivateSaleInvalidAmountAllocation ErrorCode = 3180
	ErrPrivateSaleExceedPool              ErrorCode = 3181
	ErrPrivateSaleInvalidTHCPrice         ErrorCode = 3182
	ErrPrivateSaleNotEnoughPTicket        ErrorCode = 3183
	ErrPrivateSaleAlreadyBuyProject       ErrorCode = 3184
	ErrPrivateSaleNotStartYet             ErrorCode = 3185
	ErrPrivateSaleAlreadyEnded            ErrorCode = 3186
	ErrPrivateSaleTHCPriceExpired         ErrorCode = 3187
	ErrPrivateSaleUserProjectNotFound     ErrorCode = 3188
	ErrPrivateSaleInvalidTHCFee           ErrorCode = 3189
	ErrPrivateSaleInactiveProject         ErrorCode = 3190

	ErrNone               ErrorCode = 4002
	ErrUnmarshal          ErrorCode = 4003
	ErrMinionNotFound     ErrorCode = 4004
	ErrMinionNotSelected  ErrorCode = 4005
	ErrMinionNotOwner     ErrorCode = 4006
	ErrUserIsLocked       ErrorCode = 4007
	ErrMinionIsLocked     ErrorCode = 4008
	ErrInputInvalid       ErrorCode = 4009
	ErrBattleNotFound     ErrorCode = 4010
	ErrObjectInvalid      ErrorCode = 4011
	ErrObjectNotAvailable ErrorCode = 4012
	ErrObjectNotFound     ErrorCode = 4013
	ErrDBException        ErrorCode = 4014

	// Quest
	ErrQuestNotReadyForClaim ErrorCode = 4101
	ErrQuestAlreadyClaim     ErrorCode = 4102
	ErrQuestLock             ErrorCode = 4103

	// Cosmetic
	ErrCosmeticProfileAlreadyExist ErrorCode = 4200

	// Referral Friend
	ErrInvalidInviter            ErrorCode = 4301
	ErrInvalidInvitee            ErrorCode = 4302
	ErrInvalidReferralType       ErrorCode = 4303
	ErrAlreadyClaimRewards       ErrorCode = 4304
	ErrReferralFriendMaintenance ErrorCode = 4305

	// Thetan Rival still here

	/* THETAN UGC */

	ErrUGCItemNotFound           ErrorCode = 5001
	ErrUGCInsufficientBalance    ErrorCode = 5002
	ErrUGCInvalidTextureFile     ErrorCode = 5003
	ErrUGCMissingTextureFile     ErrorCode = 5004
	ErrUGCUnsupportedFormat      ErrorCode = 5005
	ErrUGCNotOwnItem             ErrorCode = 5006
	ErrUGCUnsupportedItem        ErrorCode = 5007
	ErrUGCInvalidTextureChecksum ErrorCode = 5008
	ErrUGCRateLimited            ErrorCode = 5009
)

var errMessage = map[ErrorCode]string{
	// HTTP Status Code 400 - 599 (Client Error and Server Error)
	ErrHTTPNotFound:            "404 - Not Found",
	ErrHTTPInternalServerError: "500 - Internal Server Error",

	// Common code 1 - 1000
	ErrTokenInvalid:          "Token Invalid",
	ErrTokenExpired:          "Token Expired",
	ErrUserNotFound:          "User Not Found",
	ErrUserBanned:            "User Banned",
	ErrWrongFormat:           "Wrong Data Format",
	ErrWrongPrice:            "Wrong Price Config",
	ErrOutOfLimit:            "Your Quantity Is Limited",
	ErrCannotClaim:           "You Can't Claim Now",
	ErrNewVersionAvailable:   "New Version Available",
	ErrUserBannedFindMatch:   "User Has Been Banned FindMatch",
	ErrCantGetCurrencyPrice:  "System Can't Get Currency Price",
	ErrExchangeRateDifferent: "Exchange Rate Is Not The Same At This Time, Try Again!",
	ErrMaintenanceCashOut:    "Feature Cashout Under Maintenance",
	ErrOutdatedSession:       "The session has expired due to inactivity. Please try again",
	ErrMismatchToken:         "You can't claim this token",
	ErrBusyServer:            "Server is too busy now, try again later",
	ErrClaimedToken:          "This token was claimed",
	ErrMaintainedServer:      "Server Currently Is Maintenance",
	ErrComingSoon:            "This feature is under construction",

	// WolffunID Err_CODE 2001 -> 3000:
	ErrEmailNotExist:             "Email Not Exist In Database",
	ErrEmailIsExist:              "Email Exist In Database",
	ErrCodeNotValid:              "Code Not Valid",
	ErrEmailNotValid:             "Email Not Valid",
	ErrInCorrectUserId:           "Incorrect User Id",
	ErrNotEnoughTicketChangeName: "User Not Enough Ticket Change Name",
	ErrUserNameHasTaken:          "This Name Has Taken",
	ErrTooManyAccountLogin:       "Too Many Account Login To This Device",
	ErrAccountHasBeenLinked:      "The Account Has Been Linked",

	//theta data code 3001 -> 4000
	ErrIdInvalid:            "Id Not Valid",
	ErrAddressInvalid:       "Address Not Valid",
	ErrTypeInvalid:          "Type Not Valid",
	ErrHeroIdInvalid:        "HeroId Not Valid",
	ErrInvalidate:           "Not Valid",
	ErrNotEnough:            "Not Enough",
	ErrNotEnoughTHC:         "Not Enough Thetan Coin",
	ErrNotEnoughTHG:         "Not Enough Thetan Gem",
	ErrNotEnoughPP:          "Not Enough Power Point",
	ErrNotOwnerHero:         "You Are Not The Owner Of This Hero",
	ErrHeroNotAvailable:     "Hero Is Not Available",
	ErrNotRankEnoughCashout: "Reach Rank Bronze 1 To Start Claiming Your Token",
	ErrNotOldEnoughCashout:  "Start Claiming Your Token After: %v",
	//OutOfLimitCashout:    "Cash Out Is Out Of Limit",
	ErrNotEnoughGTHG:     "Not Enough InGame Thetan Gem",
	ErrNotEnoughGTHC:     "Not Enough InGame Thetan Coin",
	ErrInvalidGTHG:       "Invalid gTHG Balance!",
	ErrNotEnoughgFeeGTHG: "Insufficient gTHG! Please Deposit To Proceed.",
	ErrHeroInTheMatch:    "Hero In The Match",
	ErrMinimumCashout:    "Min convertable: %v",
	ErrMaximumCashout:    "Max convertable: %v",

	//box
	ErrOutdatedBuying:  "Outdated Buying",
	ErrOutOfBoxes:      "Out Of Thetan Boxes",
	ErrNotEnoughBox:    "Not Enough Quantity Box",
	ErrBoxExchangeRate: "Box price has been updated, please try again!",

	// Gift box
	ErrMustLinkedWolffunId: "Hey, you must link your account with WolffunId first!",
	ErrMustInstalledGame:   "Hey, you must install ThetanArena and log in to your account!",
	ErrNeedMoreTimeToOpen:  "Hmm, you need at least 24 hours before opening your gift box!",
	ErrAlreadyClaimFreeBox: "Oh no, you already claimed this box before!",
	ErrMustConnectCoinbase: "Hey, you must connect your account with the Coinbase wallet!",
	ErrGiftCodeInvalid:     "Oh no, the gift code is not valid or expired!",
	ErrRequestProcessing:   "Your request is being handled!",

	//friend
	ErrOtherFullRequestReceived: "Your friend can't take any more friend requests",
	ErrOtherFullFriendList:      "Your friend has full friend list",
	ErrFullFriendList:           "You has full friend list",
	ErrReceiptDataInValid:       "Receipt Data Invalid",

	// report battle end
	ErrReachLimitReportInWeek: "Reach limit report in week (Max: 7 times/week)",
	ErrPreviouslyReported:     "Previously reported",

	// rental
	ErrRentedItem:            "This hero has been rented by someone",
	ErrLowRentBattles:        "Insufficient number of battles for rent",
	ErrHighRentBattles:       "Exceed the number of rent battles",
	ErrLowRentPrice:          "Rental price is lower than minimum",
	ErrHighRentPrice:         "Rental price is higher than maximum",
	ErrItemIsForRent:         "Item is being rent out",
	ErrMaxBattleRented:       "The rental item has reached its battle limit",
	ErrReturningOwner:        "The rental item is being returned",
	ErrMaintenanceHeroRental: "The rental feature is being maintained",

	// Vesting safe
	ErrMaintenanceVestingSafe: "This feature is under maintenance, please try again in a few minutes",
	ErrAlreadyClaimed:         "You have already received all gTHG today, please try again the next day",

	// Creator Program
	ErrMaintenanceCreatorProgram: "This feature is under maintenance, please try again in a few minutes",
	ErrAreNotCreator:             "You are a Creator. You can’t earn gift as a Viewer",
	ErrAreNotViewer:              "You are not a Creator, please contact us for more information",

	// Staking
	ErrMaintenanceStaking:  "This feature is under maintenance, please try again in a few minutes",
	ErrInvalidPlan:         "The staking plan is invalid",
	ErrInvalidStakeAmount:  "The staking amount is invalid",
	ErrOutOfTimeStake:      "Can not stake this time",
	ErrStakePoolIsFull:     "Staking pool is full",
	ErrHeroUnavailable:     "Your selected heroes is unavailable",
	ErrNotReadyToUnstake:   "Your stake is not ready to unstake",
	ErrPreviousUnstake:     "You are previous unstake",
	ErrCooldownClaimReward: "You only claim reward each every 24 hours",
	ErrMaxPendingStake:     "Waiting for your pending stake processed after new stake",
	ErrEmptyReward:         "Current reward for this stake is empty",
	ErrNotReadyForClaim:    "This stake is not ready for claim reward",
	ErrMaxHeroesPerStake:   "Limit heroes per stake. You can stake maximum 200 heroes per stake",

	ErrNone:               "TR_ERROR_NONE",
	ErrUnmarshal:          "TR_ERROR_UNMARSHAL",
	ErrMinionNotFound:     "TR_ERROR_MINION_NOT_FOUND",
	ErrMinionNotSelected:  "TR_ERROR_MINION_NOT_SELECTED",
	ErrMinionNotOwner:     "TR_ERROR_MINION_NOT_OWNER",
	ErrUserIsLocked:       "TR_ERROR_USER_IS_LOCKED",
	ErrMinionIsLocked:     "TR_ERROR_MINION_IS_LOCKED",
	ErrInputInvalid:       "TR_ERROR_INPUT_INVALID",
	ErrBattleNotFound:     "TR_ERROR_BATTLE_NOT_FOUND",
	ErrObjectInvalid:      "TR_ERROR_%v_INVALID",
	ErrObjectNotAvailable: "TR_ERROR_%v_NOT_AVAILABLE",
	ErrObjectNotFound:     "TR_ERROR_%v_NOT_FOUND",
	ErrDBException:        "TR_ERROR_DB_EXCEPTION",

	// Quest
	ErrQuestNotReadyForClaim: "This quest is not ready for claim now",
	ErrQuestAlreadyClaim:     "You have already claimed this quest",
	ErrQuestLock:             "Your quest is locked due to insufficient battles",

	// Cosmetic
	ErrCosmeticProfileAlreadyExist: "You already owned this profile item",

	// Referral Friend
	ErrInvalidInviter:            "Inviter is not valid",
	ErrInvalidInvitee:            "Invitee is not valid",
	ErrInvalidReferralType:       "Referral type is not valid",
	ErrAlreadyClaimRewards:       "You are already claim all rewards now",
	ErrReferralFriendMaintenance: "This feature is under maintenance, please try again in a few minutes",

	/* THETAN UGC */

	ErrUGCItemNotFound:           "Your item not found",
	ErrUGCInsufficientBalance:    "Your C-Point balance is insufficient",
	ErrUGCInvalidTextureFile:     "Your texture file provided invalid",
	ErrUGCMissingTextureFile:     "You have not uploaded the texture file",
	ErrUGCUnsupportedFormat:      "Your texture file provided is unsupported format",
	ErrUGCNotOwnItem:             "You do not own this item",
	ErrUGCUnsupportedItem:        "Your item currently is unsupported",
	ErrUGCInvalidTextureChecksum: "The texture file and MD5 checksum you provided are invalid",
	ErrUGCRateLimited:            "We are rate limited",
}
