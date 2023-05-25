package esmatchmaking

type InGameInfo struct {
	GameMode        int    `json:"game_mode"`
	InGameMode      int    `json:"ingame_mode"`
	Map             int    `json:"map"`
	TicketRegions   string `json:"ticket_regions"`
	HeroId          string `json:"hero_id"`
	Skill1Id        int    `json:"skill_1"`
	Skill2Id        int    `json:"skill_2"`
	Skill1Level     int    `json:"skill_1_level"`
	Skill2Level     int    `json:"skill_2_level"`
	TeamId          int    `json:"team_id"`
	IsSpecialEvent  bool   `json:"is_special_event"`
	RankingLevel    int    `json:"ranking_level"`
	TrophyRanking   int    `json:"trophy_ranking"`
	BatttleNumber   int    `json:"batttle_number"`
	Ping            int    `json:"ping"`
	CountPlayer     int    `json:"count_player"`
	CountBot        int    `json:"count_bot"`
	AverageTrophies int    `json:"average_trophies"`
	BestRegion      int    `json:"best_region"`
	BehaviorPoint   int    `json:"behavior_point"`
	ArrBrainType    string `json:"arr_brain_type"`
}
