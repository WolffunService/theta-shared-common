package common

const (

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
	OutOfLimitCashout    = 3014
	NotEnoughGTHG        = 3015
	NotEnoughGTHC        = 3016
	InvalidGTHG          = 3017
	NotEnoughgFeeGTHG    = 3018
	HeroInTheMatch       = 3019

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

	// Box =))
	BoxExchangeRate = 3109
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
	OutOfLimitCashout:    "Cash Out Is Out Of Limit",
	NotEnoughGTHG:        "Not Enough InGame Thetan Gem",
	NotEnoughGTHC:        "Not Enough InGame Thetan Coin",
	InvalidGTHG:          "Invalid gTHG Balance!",
	NotEnoughgFeeGTHG:    "Insufficient gTHG! Please Deposit To Proceed.",
	HeroInTheMatch:       "Hero In The Match",

	//box
	OutdatedBuying:  "Outdated Buying",
	OutOfBoxes:      "Out Of Thetan Boxes",
	NotEnoughBox:    "Not Enough Quantity Box",
	BoxExchangeRate: "Box price has been updated, please try again!",

	//friend
	OtherFullRequestReceived: "Your friend can't take any more friend requests",
	OtherFullFriendList:      "Your friend has full friend list",
	FullFriendList:           "You has full friend list",
	ReceiptDataInValid:       "Receipt Data Invalid",

	// report battle end
	ReachLimitReportInWeek: "Reach limit report in week (Max: 7 times/week)",
	PreviouslyReported:     "Previously reported",
}

// StatusText returns a text for the common error code. It returns the empty
// string if the code is unknown.
func ErrorText(code int) string {
	return errorText[code]
}
