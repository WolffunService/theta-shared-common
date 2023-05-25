package thetanrivalmatchmakingenum

type MapLevel int

const (
	MAP_LEVEL_NONE MapLevel = iota
	MAP_LEVEL_EASY
	MAP_LEVEL_NORMAL
	MAP_LEVEL_HARD
)

type MapTypeId int32

const (
	MAP_TYPE_ID_NONE   MapTypeId = 0
	MAP_TYPE_ID_NEWBIE MapTypeId = 1

	MAP_TYPE_ID_OBSTACLE_RACING       MapTypeId = 2
	MAP_TYPE_ID_GREEN_LIGHT_RED_LIGHT MapTypeId = 4
	MAP_TYPE_ID_TILES_FALL            MapTypeId = 6
	MAP_TYPE_ID_MAZE_FALL             MapTypeId = 11
	MAP_TYPE_ID_OR_PET                MapTypeId = 13
	MAP_TYPE_ID_JUNGLE_OR             MapTypeId = 14
	MAP_TYPE_ID_SEESAW_PARK           MapTypeId = 15
	MAP_TYPE_ID_BIG_WHEEL             MapTypeId = 16
	MAP_TYPE_ID_WATERFALL_CLIMB       MapTypeId = 17
	MAP_TYPE_ID_SKI_RUN               MapTypeId = 18
	MAP_TYPE_ID_ROLL_ALONG            MapTypeId = 19
	MAP_TYPE_ID_GEAR_BOX              MapTypeId = 20
	MAP_TYPE_ID_SUNDAY_PARK           MapTypeId = 21
	MAP_TYPE_ID_ZIGZAG_CLIMB          MapTypeId = 22
	MAP_TYPE_ID_BAGGAGE_CLAIM         MapTypeId = 23
	MAP_TYPE_ID_SPIN_AROUND           MapTypeId = 24
	MAP_TYPE_ID_ROOFTOP_RUSH          MapTypeId = 25
	MAP_TYPE_ID_DEEP_WATER            MapTypeId = 26
	MAP_TYPE_ID_RAINBOW_SKATE         MapTypeId = 27
	MAP_TYPE_ID_CROSSING_BRIDGE       MapTypeId = 28
	MAP_TYPE_ID_CANDY_DREAM           MapTypeId = 29
	MAP_TYPE_ID_CROSSY_ROAD           MapTypeId = 30
	MAP_TYPE_ID_AMAZONE_RIVER         MapTypeId = 31

	MAP_TYPE_ID_HOT_POTATO   MapTypeId = 3
	MAP_TYPE_ID_LASER_DANCE  MapTypeId = 5
	MAP_TYPE_ID_SCORING_AREA MapTypeId = 7
	MAP_TYPE_ID_BIG_CANNON   MapTypeId = 9
)

func (t MapTypeId) IsRaceGame() bool {
	switch t {
	case MAP_TYPE_ID_OBSTACLE_RACING,
		MAP_TYPE_ID_GREEN_LIGHT_RED_LIGHT,
		MAP_TYPE_ID_TILES_FALL,
		MAP_TYPE_ID_MAZE_FALL,
		MAP_TYPE_ID_OR_PET,
		MAP_TYPE_ID_JUNGLE_OR,
		MAP_TYPE_ID_SEESAW_PARK,
		MAP_TYPE_ID_BIG_WHEEL,
		MAP_TYPE_ID_WATERFALL_CLIMB,
		MAP_TYPE_ID_SKI_RUN,
		MAP_TYPE_ID_ROLL_ALONG,
		MAP_TYPE_ID_GEAR_BOX,
		MAP_TYPE_ID_SUNDAY_PARK,
		MAP_TYPE_ID_ZIGZAG_CLIMB,
		MAP_TYPE_ID_BAGGAGE_CLAIM,
		MAP_TYPE_ID_SPIN_AROUND,
		MAP_TYPE_ID_ROOFTOP_RUSH,
		MAP_TYPE_ID_DEEP_WATER,
		MAP_TYPE_ID_RAINBOW_SKATE,
		MAP_TYPE_ID_CROSSING_BRIDGE,
		MAP_TYPE_ID_CANDY_DREAM,
		MAP_TYPE_ID_CROSSY_ROAD,
		MAP_TYPE_ID_AMAZONE_RIVER:
		return true
	}
	return false
}

func (t MapTypeId) IsSurvivalGame() bool {
	switch t {
	case MAP_TYPE_ID_HOT_POTATO,
		MAP_TYPE_ID_LASER_DANCE,
		MAP_TYPE_ID_SCORING_AREA,
		MAP_TYPE_ID_BIG_CANNON:
		return true
	}
	return false
}

func GetMapTypeID(mapID int) MapTypeId {
	return MapTypeId(mapID / 100)
}

func (t MapTypeId) GetGameRule() GameRule {
	if t.IsRaceGame() {
		return GAME_RULE_RACING
	}
	if t.IsSurvivalGame() {
		return GAME_RULE_SURVIVAL
	}

	return GAME_RULE_NONE
}

type GameRule int

const (
	GAME_RULE_NONE = iota
	GAME_RULE_RACING
	GAME_RULE_SURVIVAL
)

func (g GameRule) IsGameRuleRacing() bool {
	return g == GAME_RULE_RACING
}

func (g GameRule) IsGameRuleSurvival() bool {
	return g == GAME_RULE_SURVIVAL
}
