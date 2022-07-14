package common

const (

	// Error TODO
	//Renaming all error code to use Error prefix ex:
	//ErrorConflict     = 3  // action cannot be performed
	//ErrorInternal     = 1  // internal error
	//ErrorInvalid      = 2  // validation failed
	//ErrorNotFound     = 4  // entity does not exist
	//ErrorValidate     = 5  // entity does not exist
	//ErrorAuthorize    = 11 //cannot authorize
	//ErrorAuthenticate = 10 //cannot authorize
	//ErrorStorage      = 20
	//ErrorDatabase     = 30
	//ErrorOTPLimit     = 40
	//ErrorOTPInvalid   = 41
	//ErrorInputForm    = 50
	//
	//
	//ErrorUserPassword   = 1000
	//ErrorUserName       = 1001
	//ErrorMailRegistered = 1002

	//common code 1 - 1000
	Error                 = 0
	TokenInvalid          = 1
	TokenExpired          = 2
	UserNotFound          = 3
	UserBanned            = 4
	WrongFormat           = 5
	WrongPrice            = 6
	OutOfLimit            = 7
	CannotClaim           = 8
	NewVersionAvailable   = 9
	UserBannedFindMatch   = 10
	CantGetCurrencyPrice  = 11
	ExchangeRateDifferent = 12
	MaintenanceCashout    = 13
	OutdatedSession       = 14
	MismatchToken         = 15
	BusyServer            = 16
	ClaimedToken          = 17
	MaintainceServer      = 18
	ComingSoon            = 19
	DataNotFound          = 20
	DatabaseException     = 21
	CurrencyException     = 22
	AlreadyClaimedReward  = 23

	// turn + season
	ExpiredTurn      = 100
	ExpiredSeason    = 101
	NotExpiredSeason = 102
	NotStartSeason   = 103

	// price
	ExpiredPrice = 200

	// reward
	RewardNotFound      = 300
	NotYetReceiveReward = 301

	//wolffunId ERROR_CODE 2001 -> 3000:
	EmailNotExist             = 2001
	EmailIsExist              = 2002
	CodeNotValid              = 2003
	EmailNotValid             = 2004
	InCorrectUserId           = 2005
	NotEnoughTicketChangeName = 2006
	UserNameHasTaken          = 2007
	TooManyAccountLogin       = 2008
	AccountHasBeenLinked      = 2009

	//theta data code 3001 -> 4000
	IdInvalid            = 3001
	AddressInvalid       = 3002
	TypeInvalid          = 3003
	HeroIdInvalid        = 3004
	Invalidate           = 3005
	NotEnough            = 3006
	NotEnoughTHC         = 3007
	NotEnoughTHG         = 3008
	NotEnoughPP          = 3009
	NotOwnerHero         = 3010
	HeroNotAvailable     = 3011
	NotRankEnoughCashout = 3012
	NotOldEnoughCashout  = 3013
	//OutOfLimitCashout    = 3014
	NotEnoughGTHG           = 3015
	NotEnoughGTHC           = 3016
	InvalidGTHG             = 3017
	NotEnoughgFeeGTHG       = 3018
	HeroInTheMatch          = 3019
	MinimumCashout          = 3020
	MaximumCashout          = 3021
	WrongHeroModel          = 3022
	NotEnoughMinigameTicket = 3023
	MaximumHeroLevel        = 3024

	//box
	OutdatedBuying = 3100
	OutOfBoxes     = 3101
	NotEnoughBox   = 3102

	//friend
	OtherFullRequestReceived = 3103
	OtherFullFriendList      = 3104
	FullFriendList           = 3105
	ReceiptDataInValid       = 3106

	//report battle end
	ReachLimitReportInWeek = 3107
	PreviouslyReported     = 3108

	// Box
	BoxExchangeRate        = 3109
	MustLinkedWolffunId    = 3110
	MustInstalledGame      = 3111
	NeedMoreTimeToOpen     = 3112
	AlreadyClaimFreeBox    = 3113
	MustConnectCoinbase    = 3114
	GiftCodeInvalid        = 3115
	ErrorRequestProcessing = 3116

	// Rental
	RentedItem            = 3120
	LowRentBattles        = 3121
	HighRentBattles       = 3122
	LowRentPrice          = 3123
	ItemIsForRent         = 3124
	MaxBattleRented       = 3125
	HighRentPrice         = 3126
	ReturningOwner        = 3127
	MaintenanceHeroRental = 3128
	InvalidSignature      = 3129
	InvalidLockDuration   = 3130
	InvalidMinGTHCBatles  = 3131

	// VestingSafe
	MaintenanceVestingSafe = 3140
	AlreadyClaimed         = 3141

	// Ingame Skill Ads
	TheDayHasPassed = 3150
	MaxAdsViewed    = 3151
	SkillIdNotFound = 3152

	// Minigame
	MinigameTooFast = 3200

	// fusion
	ErrWrongSet     = 3300
	ErrFusing       = 3301
	ErrNegativeCost = 3302

	// Cosmetic
	ItemNotAvailable = 3310

	// Hero buy/sell/mint
	LowSellPrice            = 3320
	HighSellPrice           = 3321
	MintLowHeroLevel        = 3322
	MintLowUserRankingLevel = 3323
	MintCooldownTime        = 3324
)

var errorText = map[int]string{
	//common code 1 - 1000
	TokenInvalid:          "Token Invalid",
	TokenExpired:          "Token Expired",
	UserNotFound:          "User Not Found",
	UserBanned:            "User Banned",
	WrongFormat:           "Wrong Data Format",
	WrongPrice:            "Wrong Price Config",
	OutOfLimit:            "Your Quantity Is Limited",
	CannotClaim:           "You Can't Claim Now",
	NewVersionAvailable:   "New Version Available",
	UserBannedFindMatch:   "User Has Been Banned FindMatch",
	CantGetCurrencyPrice:  "System Can't Get Currency Price",
	ExchangeRateDifferent: "Exchange Rate Is Not The Same At This Time, Try Again!",
	MaintenanceCashout:    "Feature Cashout Under Maintenance",
	OutdatedSession:       "The session has expired due to inactivity. Please try again",
	MismatchToken:         "You can't claim this token",
	BusyServer:            "Server is too busy now, try again later",
	ClaimedToken:          "This token was claimed",
	MaintainceServer:      "Server Currently Is Maintenance",
	ComingSoon:            "This feature is under construction",
	DataNotFound:          "Data not found",
	DatabaseException:     "Database exception occurred. Please try again",
	CurrencyException:     "Process currency exception occurred. Please try again",
	AlreadyClaimedReward:  "You have already claimed",

	// general
	// turn + season
	ExpiredTurn:      "Turned expired!",
	ExpiredSeason:    "Event has ended!",
	NotExpiredSeason: "Unable to proceed! Please wait for the event to end.",
	NotStartSeason:   "Unable to proceed! Please wait for the event to start.",

	// price
	ExpiredPrice: "Price updated! Please refresh to proceed.",

	// reward
	RewardNotFound:      "Unable to claim! No reward is available.",
	NotYetReceiveReward: "Please claim all rewards from the previous event to join a new one.",

	//wolffunId ERROR_CODE 2001 -> 3000:
	EmailNotExist:             "Email Not Exist In Database",
	EmailIsExist:              "Email Exist In Database",
	CodeNotValid:              "Code Not Valid",
	EmailNotValid:             "Email Not Valid",
	InCorrectUserId:           "Incorrect User Id",
	NotEnoughTicketChangeName: "User Not Enough Ticket Change Name",
	UserNameHasTaken:          "This Name Has Taken",
	TooManyAccountLogin:       "Too Many Account Login To This Device",
	AccountHasBeenLinked:      "The Account Has Been Linked",

	//theta data code 3001 -> 4000
	IdInvalid:            "Id Not Valid",
	AddressInvalid:       "Address Not Valid",
	TypeInvalid:          "Type Not Valid",
	HeroIdInvalid:        "HeroId Not Valid",
	Invalidate:           "Not Valid",
	NotEnough:            "Not Enough",
	NotEnoughTHC:         "Not Enough Thetan Coin",
	NotEnoughTHG:         "Not Enough Thetan Gem",
	NotEnoughPP:          "Not Enough Power Point",
	NotOwnerHero:         "You Are Not The Owner Of This Hero",
	HeroNotAvailable:     "Hero Is Not Available",
	NotRankEnoughCashout: "Reach Rank Bronze 1 To Start Claiming Your Token",
	NotOldEnoughCashout:  "Start Claiming Your Token After: %v",
	//OutOfLimitCashout:    "Cash Out Is Out Of Limit",
	NotEnoughGTHG:     "Not Enough InGame Thetan Gem",
	NotEnoughGTHC:     "Not Enough InGame Thetan Coin",
	InvalidGTHG:       "Invalid gTHG Balance!",
	NotEnoughgFeeGTHG: "Insufficient gTHG! Please Deposit To Proceed.",
	HeroInTheMatch:    "Hero In The Match",
	MinimumCashout:    "Min convertable: %v",
	MaximumCashout:    "Max convertable: %v",

	//box
	OutdatedBuying:  "Outdated Buying",
	OutOfBoxes:      "Out Of Thetan Boxes",
	NotEnoughBox:    "Not Enough Quantity Box",
	BoxExchangeRate: "Box price has been updated, please try again!",

	// Gift box
	MustLinkedWolffunId:    "Hey, you must link your account with WolffunId first!",
	MustInstalledGame:      "Hey, you must install ThetanArena and log in to your account!",
	NeedMoreTimeToOpen:     "Hmm, you need at least 24 hours before opening your gift box!",
	AlreadyClaimFreeBox:    "Oh no, you already claimed this box before!",
	MustConnectCoinbase:    "Hey, you must connect your account with the Coinbase wallet!",
	GiftCodeInvalid:        "Oh no, the gift code is not valid or expired!",
	ErrorRequestProcessing: "Your request is being handled!",

	//friend
	OtherFullRequestReceived: "Your friend can't take any more friend requests",
	OtherFullFriendList:      "Your friend has full friend list",
	FullFriendList:           "You has full friend list",
	ReceiptDataInValid:       "Receipt Data Invalid",

	// report battle end
	ReachLimitReportInWeek: "Reach limit report in week (Max: 7 times/week)",
	PreviouslyReported:     "Previously reported",

	// rental
	RentedItem:            "This hero has been rented by someone",
	LowRentBattles:        "Insufficient number of battles for rent",
	HighRentBattles:       "Exceed the number of rent battles",
	LowRentPrice:          "Rental price is lower than minimum",
	HighRentPrice:         "Rental price is higher than maximum",
	ItemIsForRent:         "Item is being rent out",
	MaxBattleRented:       "The rental item has reached its battle limit",
	ReturningOwner:        "The rental item is being returned",
	MaintenanceHeroRental: "The rental feature is being maintained",
	InvalidSignature:      "The signature is wrong",
	InvalidLockDuration:   "The lock duration is wrong",
	InvalidMinGTHCBatles:  "Rental battle is lower than minimum",

	// vesting safe
	MaintenanceVestingSafe: "This feature is under maintenance, please try again in a few minutes",
	AlreadyClaimed:         "You have already received all gTHG today, please try again the next day",

	// Ingame Skill Ads
	TheDayHasPassed: "The day has passed, please syndata and try again.",
	MaxAdsViewed:    "Max Ads Viewed",
	SkillIdNotFound: "SkillId not found, please syndata and try again.",

	// Minigame
	MinigameTooFast:         "You play too fast, please try again.",
	NotEnoughMinigameTicket: "Not enough ticket to play, please buy ticket.",

	// fusion
	ErrWrongSet:     "Wrong input heroes!",
	ErrFusing:       "Complete the available fusion to start the others",
	ErrNegativeCost: "Heroes used to fuse has greater total value than output hero. Please change your input!",

	// Cosmetic
	ItemNotAvailable: "Item Is Not Available",

	// Hero buy/sell
	LowSellPrice:            "Sell price is lower than minimum",
	HighSellPrice:           "Sell price is higher than maximum",
	MintLowHeroLevel:        "Hero can be minted from level 2.",
	MintLowUserRankingLevel: "Reach Private rank to mint item.",
	MintCooldownTime:        "Mint is cooling down, please try it later.",
}

const unknownErrorMessage = "An internal error has occurred. Please contact technical support!"

// ErrorText returns a text for the common error code. It returns the default
// string if the code is unknown.
func ErrorText(code int) string {
	if errorMessage, found := errorText[code]; found {
		return errorMessage
	}
	return unknownErrorMessage
}
