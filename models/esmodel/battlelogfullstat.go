package esmodel

import "time"

type BattleLogFullMapping struct {
	User                 UserModel `json:"user"`
	IsBot                bool      `json:"is_bot"`
	Brain                string    `json:"brain"`
	TrophyMatching       int       `json:"trophy_matching"`
	TrophyReward         int       `json:"trophy_reward"`
	TeamId               int       `json:"team"`
	Result               int       `json:"result"`
	Rank                 int       `json:"rank"`
	IsSpecialEvent       bool      `json:"is_special_event"`
	MatchId              string    `json:"match_id"`
	TimeStart            time.Time `json:"start"`
	TimeEnd              time.Time `json:"end"`
	InGameMode           int       `json:"ingame_mode"`
	Map                  int       `json:"map"`
	Region               int       `json:"region"`
	Team1Score           int       `json:"team_1_score"`
	Team2Score           int       `json:"team_2_score"`
	Score                int       `json:"score"`
	Kill                 int       `json:"kill"`
	Death                int       `json:"death"`
	Support              int       `json:"support"`
	DamageOut            int       `json:"damage_out"`
	DeathMatchScore      int       `json:"death_match_score"`
	DeathMatchMinusScore int       `json:"death_match_minus_score"`
	Star                 int       `json:"star"`
	KeepingTimeSuperStar int       `json:"keeping_time_super_star"`
	RobotSummonTime      int       `json:"robot_summon_time"`
	TowerDamage          int       `json:"tower_damage"`
	BatteryPickupTime    int       `json:"battery_pickup_time"`
	HeroLevelExp         int       `json:"hero_level_exp"`
	SerialKiller         bool      `json:"serial_killer"`
	TeamMVP              bool      `json:"team_mvp"`
	Nanny                bool      `json:"nanny"`
	MVP                  bool      `json:"mvp"`
	TripleKill           int       `json:"triple_kill"`
	MegaKill             int       `json:"mega_kill"`
	HeroId               string    `json:"hero_id"`
	HeroType             int       `json:"hero_type"`
	HeroLevel            int       `json:"hero_level"`
	Skin                 int       `json:"skin"`
	Skill1Id             int       `json:"skill_1"`
	Skill2Id             int       `json:"skill_2"`
	Skill1Level          int       `json:"skill_1_level"`
	Skill2Level          int       `json:"skill_2_level"`
	THC                  int       `json:"thc"`
	PowerPoint           int       `json:"power_point"`
	Exp                  int       `json:"exp"`
	Timestamp            time.Time `json:"@timestamp"`
}

func (BattleLogFullMapping) Index() string {
	return "playerstats-battle"
}
