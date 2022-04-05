package playerstatenum

import "fmt"

type StatName string

const (
	Test            StatName = "test_%d"
	PlayerBattle    StatName = "player_battle"
	PlayerWinBattle StatName = "player_win_battle"
	HeroBattle      StatName = "hero_battle_%d"  //heroId
	SkillBattle     StatName = "skill_battle_%d" //skillId
	ModeBattle      StatName = "mode_battle_%d"  //modeId
)

func (s StatName) GetWithCode(code int) string {
	return fmt.Sprintf(string(s), code)
}

func (s StatName) String() string {
	return fmt.Sprintf(string(s))
}
