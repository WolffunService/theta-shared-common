package mkpquestenum

type QuestSubType int

const (
	startSubQuest QuestSubType = iota
	HeroRarity
	HeroType
	AnyHero
	AnyCosmetic
	THC
	THG
	Anybox
	BoxRarity
	ThetanTournament
	FusionHeroRarity
	WelcomBox
	endSubQuest
)

var questSubTypeText = map[QuestSubType]string{
	HeroRarity:       "HeroRarity",
	HeroType:         "HeroType",
	AnyHero:          "AnyHero",
	AnyCosmetic:      "AnyCosmetic",
	THC:              "THC",
	THG:              "THG",
	Anybox:           "Anybox",
	BoxRarity:        "BoxRarity",
	ThetanTournament: "ThetanTournament",
	FusionHeroRarity: "FusionHeroRarity",
	WelcomBox:        "WelcomBox",
}

var difficultyQuestSubType = []QuestSubType{
	startSubQuest,
	AnyHero,
	HeroRarity,
	Anybox,
	HeroType,
	BoxRarity,
	THC,
	THG,
	AnyCosmetic,
	ThetanTournament,
	FusionHeroRarity,
	endSubQuest,
}

func (q QuestSubType) String() string {
	return questSubTypeText[q]
}

func (q QuestSubType) IsValid() bool {
	return q > startSubQuest && q < endSubQuest
}

func (q QuestSubType) GetDiff() int {
	for i := range difficultyQuestSubType {
		if difficultyQuestSubType[i] == q {
			return i
		}
	}
	return -1
}

var QST = startSubQuest

func (QuestSubType) GetMaxDiff() int {
	return len(difficultyQuestSubType) - 2 // all item len - (start + end) enum
}

func (QuestSubType) GetMinDiff() int {
	//first item slice
	return 1
}

func (QuestSubType) GetAllQuestSubType() []QuestSubType {
	var slice []QuestSubType
	for i := startSubQuest + 1; i < endSubQuest; i++ {
		slice = append(slice, i)
	}
	return slice
}
