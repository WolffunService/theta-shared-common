package dailyquestenum

type QuestSubType int

const (
	Hero QuestSubType = 1 + iota
	Skill
	SkillCombo
	Mode
	PlayAgainCount
	PartyUp
	LikeTeammate
	LikeOpponent
	ShareRoom
)
