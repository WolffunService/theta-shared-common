package tradingeventenum

type QuestType int

const (
	_ QuestType = iota
	QT_HERO_TROPHY_CLASS
	QT_HERO_LEVEL
	QT_HERO_CLASS
	QT_SKIN_RARITY
	QT_ALL_HEROES
)
