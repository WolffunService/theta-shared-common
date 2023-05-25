package dailyquestenum

type QuestType int

const (
	PlayBattle QuestType = 1 + iota
	DefeatEnemy
	WinBattle
	DestroyBox
	DestroyGiantBox
	DoubleKill
	WinZeroDeath
	CollectStar
	CollectSuperStar
	CollectBattery
	DestroyTower
	Feature
	UseFury
)
