package thetaerror

const (
	//common code 1 - 1000
	ErrorNil                   = 0
	ErrorInternal              = 99
	ErrorTokenInvalid          = 1
	ErrorTokenExpired          = 2
	ErrorUserNotFound          = 3
	ErrorUserBanned            = 4
	ErrorWrongFormat           = 5
	ErrorWrongPrice            = 6
	ErrorOutOfLimit            = 7
	ErrorCannotClaim           = 8
	ErrorNewVersionAvailable   = 9
	ErrorUserBannedFindMatch   = 10
	ErrorCantGetCurrencyPrice  = 11
	ErrorExchangeRateDifferent = 12
	ErrorMaintenanceCashout    = 13
	ErrorOutdatedSession       = 14
	ErrorMismatchToken         = 15
	ErrorBusyServer            = 16
	ErrorClaimedToken          = 17
	ErrorMaintainceServer      = 18
	ErrorComingSoon            = 19

	//wolffunId ERROR_CODE 2001 -> 3000:
	ErrorEmailNotExist             = 2001
	ErrorEmailIsExist              = 2002
	ErrorCodeNotValid              = 2003
	ErrorEmailNotValid             = 2004
	ErrorInCorrectUserId           = 2005
	ErrorNotEnoughTicketChangeName = 2006
	ErrorUserNameHasTaken          = 2007
	ErrorTooManyAccountLogin       = 2008
	ErrorAccountHasBeenLinked      = 2009

	//theta data code 3001 -> 4000
	ErrorIdInvalid            = 3001
	ErrorAddressInvalid       = 3002
	ErrorTypeInvalid          = 3003
	ErrorHeroIdInvalid        = 3004
	ErrorInvalidate           = 3005
	ErrorNotEnough            = 3006
	ErrorNotEnoughTHC         = 3007
	ErrorNotEnoughTHG         = 3008
	ErrorNotEnoughPP          = 3009
	ErrorNotOwnerHero         = 3010
	ErrorHeroNotAvailable     = 3011
	ErrorNotRankEnoughCashout = 3012
	ErrorNotOldEnoughCashout  = 3013
	//OutOfLimitCashout    = 3014
	ErrorNotEnoughGTHG     = 3015
	ErrorNotEnoughGTHC     = 3016
	ErrorInvalidGTHG       = 3017
	ErrorNotEnoughgFeeGTHG = 3018
	ErrorHeroInTheMatch    = 3019
	ErrorMinimumCashout    = 3020
	ErrorMaximumCashout    = 3021

	//box
	ErrorOutdatedBuying = 3100
	ErrorOutOfBoxes     = 3101
	ErrorNotEnoughBox   = 3102

	//friend
	ErrorOtherFullRequestReceived = 3103
	ErrorOtherFullFriendList      = 3104
	ErrorFullFriendList           = 3105
	ErrorReceiptDataInValid       = 3106

	//report battle end
	ErrorReachLimitReportInWeek = 3107
	ErrorPreviouslyReported     = 3108

	// Box
	ErrorBoxExchangeRate = 3109

	// Rental
	ErrorRentedItem            = 3120
	ErrorLowRentBattles        = 3121
	ErrorHighRentBattles       = 3122
	ErrorLowRentPrice          = 3123
	ErrorItemIsForRent         = 3124
	ErrorMaxBattleRented       = 3125
	ErrorHighRentPrice         = 3126
	ErrorReturningOwner        = 3127
	ErrorMaintenanceHeroRental = 3128

	// VestingSafe
	ErrorMaintenanceVestingSafe = 3140
	ErrorAlreadyClaimed         = 3141

	// Creator Program
	ErrorMaintenanceCreatorProgram = 3150

	// Staking
	ErrorMaintenanceStaking  = 3160
	ErrorInvalidPlan         = 3161
	ErrorInvalidStakeAmount  = 3162
	ErrorOutOfTimeStake      = 3163
	ErrorStakePoolIsFull     = 3164
	ErrorHeroUnavailable     = 3165
	ErrorNotReadyToUnstake   = 3166
	ErrorPreviousUnstake     = 3167
	ErrorCooldownClaimReward = 3168
	ErrorMaxPendingStake     = 3169
	ErrorEmptyReward         = 3170
	ErrorNotReadyForClaim    = 3171

	// PrivateSale
	ErrorPrivateSaleInvalidAmountAllocation = 3180
	ErrorPrivateSaleExceedPool              = 3181
	ErrorPrivateSaleInvalidTHCPrice         = 3182
	ErrorPrivateSaleNotEnoughPTicket        = 3183
	ErrorPrivateSaleAlreadyBuyProject       = 3184
	ErrorPrivateSaleNotStartYet             = 3185
	ErrorPrivateSaleAlreadyEnded            = 3186
	ErrorPrivateSaleTHCPriceExpired         = 3187
	ErrorPrivateSaleUserProjectNotFound     = 3188
	ErrorPrivateSaleInvalidTHCFee           = 3189
)

var errorText = map[int]string{
	//common code 1 - 1000
	ErrorTokenInvalid:          "Token Invalid",
	ErrorTokenExpired:          "Token Expired",
	ErrorUserNotFound:          "User Not Found",
	ErrorUserBanned:            "User Banned",
	ErrorWrongFormat:           "Wrong Data Format",
	ErrorWrongPrice:            "Wrong Price Config",
	ErrorOutOfLimit:            "Your Quantity Is Limited",
	ErrorCannotClaim:           "You Can't Claim Now",
	ErrorNewVersionAvailable:   "New Version Available",
	ErrorUserBannedFindMatch:   "User Has Been Banned FindMatch",
	ErrorCantGetCurrencyPrice:  "System Can't Get Currency Price",
	ErrorExchangeRateDifferent: "Exchange Rate Is Not The Same At This Time, Try Again!",
	ErrorMaintenanceCashout:    "Feature Cashout Under Maintenance",
	ErrorOutdatedSession:       "The session has expired due to inactivity. Please try again",
	ErrorMismatchToken:         "You can't claim this token",
	ErrorBusyServer:            "Server is too busy now, try again later",
	ErrorClaimedToken:          "This token was claimed",
	ErrorMaintainceServer:      "Server Currently Is Maintenance",
	ErrorComingSoon:            "This feature is under construction",

	//wolffunId ERROR_CODE 2001 -> 3000:
	ErrorEmailNotExist:             "Email Not Exist In Database",
	ErrorEmailIsExist:              "Email Exist In Database",
	ErrorCodeNotValid:              "Code Not Valid",
	ErrorEmailNotValid:             "Email Not Valid",
	ErrorInCorrectUserId:           "Incorrect User Id",
	ErrorNotEnoughTicketChangeName: "User Not Enough Ticket Change Name",
	ErrorUserNameHasTaken:          "This Name Has Taken",
	ErrorTooManyAccountLogin:       "Too Many Account Login To This Device",
	ErrorAccountHasBeenLinked:      "The Account Has Been Linked",

	//theta data code 3001 -> 4000
	ErrorIdInvalid:            "Id Not Valid",
	ErrorAddressInvalid:       "Address Not Valid",
	ErrorTypeInvalid:          "Type Not Valid",
	ErrorHeroIdInvalid:        "HeroId Not Valid",
	ErrorInvalidate:           "Not Valid",
	ErrorNotEnough:            "Not Enough",
	ErrorNotEnoughTHC:         "Not Enough Thetan Coin",
	ErrorNotEnoughTHG:         "Not Enough Thetan Gem",
	ErrorNotEnoughPP:          "Not Enough Power Point",
	ErrorNotOwnerHero:         "You Are Not The Owner Of This Hero",
	ErrorHeroNotAvailable:     "Hero Is Not Available",
	ErrorNotRankEnoughCashout: "Reach Rank Bronze 1 To Start Claiming Your Token",
	ErrorNotOldEnoughCashout:  "Start Claiming Your Token After: %v",
	//OutOfLimitCashout:    "Cash Out Is Out Of Limit",
	ErrorNotEnoughGTHG:     "Not Enough InGame Thetan Gem",
	ErrorNotEnoughGTHC:     "Not Enough InGame Thetan Coin",
	ErrorInvalidGTHG:       "Invalid gTHG Balance!",
	ErrorNotEnoughgFeeGTHG: "Insufficient gTHG! Please Deposit To Proceed.",
	ErrorHeroInTheMatch:    "Hero In The Match",
	ErrorMinimumCashout:    "Min convertable: %v",
	ErrorMaximumCashout:    "Max convertable: %v",

	//box
	ErrorOutdatedBuying:  "Outdated Buying",
	ErrorOutOfBoxes:      "Out Of Thetan Boxes",
	ErrorNotEnoughBox:    "Not Enough Quantity Box",
	ErrorBoxExchangeRate: "Box price has been updated, please try again!",

	//friend
	ErrorOtherFullRequestReceived: "Your friend can't take any more friend requests",
	ErrorOtherFullFriendList:      "Your friend has full friend list",
	ErrorFullFriendList:           "You has full friend list",
	ErrorReceiptDataInValid:       "Receipt Data Invalid",

	// report battle end
	ErrorReachLimitReportInWeek: "Reach limit report in week (Max: 7 times/week)",
	ErrorPreviouslyReported:     "Previously reported",

	// rental
	ErrorRentedItem:            "This hero has been rented by someone",
	ErrorLowRentBattles:        "Insufficient number of battles for rent",
	ErrorHighRentBattles:       "Exceed the number of rent battles",
	ErrorLowRentPrice:          "Rental price is lower than minimum",
	ErrorHighRentPrice:         "Rental price is higher than maximum",
	ErrorItemIsForRent:         "Item is being rent out",
	ErrorMaxBattleRented:       "The rental item has reached its battle limit",
	ErrorReturningOwner:        "The rental item is being returned",
	ErrorMaintenanceHeroRental: "The rental feature is being maintained",

	// vesting safe
	ErrorMaintenanceVestingSafe: "This feature is under maintenance, please try again in a few minutes",
	ErrorAlreadyClaimed:         "You have already received all gTHG today, please try again the next day",

	// Creator Program
	ErrorMaintenanceCreatorProgram: "This feature is under maintenance, please try again in a few minutes",

	// Staking
	ErrorMaintenanceStaking:  "This feature is under maintenance, please try again in a few minutes",
	ErrorInvalidPlan:         "The staking plan is invalid",
	ErrorInvalidStakeAmount:  "The staking amount is invalid",
	ErrorOutOfTimeStake:      "Can not stake this time",
	ErrorStakePoolIsFull:     "Staking pool is full",
	ErrorHeroUnavailable:     "Your selected heroes is unavailable",
	ErrorNotReadyToUnstake:   "Your stake is not ready to unstake",
	ErrorPreviousUnstake:     "You are previous unstake",
	ErrorCooldownClaimReward: "You only claim reward each every 24 hours",
	ErrorMaxPendingStake:     "Waiting for your pending stake processed after new stake",
	ErrorEmptyReward:         "Current reward for this stake is empty",
	ErrorNotReadyForClaim:    "This stake is not ready for claim reward",
}
