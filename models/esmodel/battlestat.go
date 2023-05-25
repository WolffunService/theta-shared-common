package esmodel

import "time"

type BattleStatMapping struct {
	User       UserModel `json:"user"`
	IngameMode string    `json:"ingame_mode"`
	Mode       string    `json:"mode"`
	Hero       string    `json:"hero"`
	Skill1     string    `json:"skill_1"`
	Skill2     string    `json:"skill_2"`
	Result     string    `json:"result"`
	Trophy     int       `json:"trophy"`
	Region     string    `json:"region"`
	Timestamp  time.Time `json:"@timestamp"`
}

func (BattleStatMapping) Index() string {
	return "playerstats-battle"
}
