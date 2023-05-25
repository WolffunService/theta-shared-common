package mkpquestenum

type QuestType int

// must input difficultyQuestType if u wanna insert new questtype
const (
	startQuestType QuestType = iota
	PutUpForSale
	PutUpForRent
	ClaimTHC
	ClaimTHG
	SpendTHCBox
	SpendTHCBuy
	SpendTHCRent
	SellSuccess
	RentedOutSuccess
	ObtainHero
	RentHero
	FusionEvent
	Upgrade
	BuyBox
	WelcomeBox
	ConvertToken
	TournamentEvent

	//insert New questType here...

	endQuestType
)

var questTypeText = map[QuestType]string{
	PutUpForSale:     "PutUpForSale",
	PutUpForRent:     "PutUpForRent",
	ClaimTHC:         "ClaimTHC",
	ClaimTHG:         "ClaimTHG",
	SpendTHCBox:      "SpendTHCBox",
	SpendTHCBuy:      "SpendTHCBuy",
	SpendTHCRent:     "SpendTHCRent",
	SellSuccess:      "SellSuccess",
	RentedOutSuccess: "RentedOutSuccess",
	ObtainHero:       "ObtainHero",
	RentHero:         "RentHero",
	FusionEvent:      "FusionEvent",
	Upgrade:          "Upgrade",
	BuyBox:           "BuyBox",
	WelcomeBox:       "WelcomeBox",
	ConvertToken:     "ConvertToken",
	TournamentEvent:  "TournamentEvent",
}

//index is difficult of quest
var difficultyQuestTypeSample1 = []QuestType{
	startQuestType,

	PutUpForSale,
	ObtainHero,
	SpendTHCRent,
	BuyBox,
	PutUpForRent,
	SpendTHCBuy,
	ClaimTHC,
	FusionEvent,
	TournamentEvent,
	RentHero,
	ClaimTHG,
	ConvertToken,
	SpendTHCBox,
	SellSuccess,
	Upgrade,
	RentedOutSuccess,

	endQuestType,
}

var difficultyQuestTypeSample2 = []QuestType{
	startQuestType,

	ConvertToken,
	Upgrade,
	SpendTHCRent,
	PutUpForRent,
	PutUpForSale,
	BuyBox,
	SpendTHCBox,
	SpendTHCBuy,
	RentHero,
	RentedOutSuccess,
	ObtainHero,
	ClaimTHC,
	SellSuccess,
	ClaimTHG,
	TournamentEvent,
	FusionEvent,

	endQuestType,
}

var difficultyQuestTypeSample3 = []QuestType{
	startQuestType,

	ConvertToken,
	ObtainHero,
	Upgrade,
	PutUpForSale,
	SpendTHCBuy,
	RentHero,
	PutUpForRent,
	SpendTHCRent,
	ClaimTHC,
	SpendTHCBox,
	ClaimTHG,
	SellSuccess,
	BuyBox,
	RentedOutSuccess,
	TournamentEvent,
	FusionEvent,

	endQuestType,
}

//sample 1 2 3
var ABQuest = map[int][]QuestType{
	1: difficultyQuestTypeSample1,
	2: difficultyQuestTypeSample2,
	3: difficultyQuestTypeSample3,
}

func (q QuestType) IsValid() bool {
	return q > startQuestType && q < endQuestType
}

func (q QuestType) String() string {
	return questTypeText[q]
}

func (q QuestType) GetDiff(sample int) int {

	dqt := ABQuest[sample]

	for i := range dqt {
		if dqt[i] == q {
			return i
		}
	}
	return -1
}

var QT = startQuestType

func (QuestType) GetMaxDiff(sample int) int {
	dqt := ABQuest[sample]
	return len(dqt) - 2 // all item len - (start + end) enum
}

func (QuestType) GetMinDiff() int {
	//first item slice
	return 1
}

func (QuestType) GetAllQuestType() []QuestType {
	var slice []QuestType
	for i := startQuestType + 1; i < endQuestType; i++ {
		slice = append(slice, i)
	}
	return slice
}
