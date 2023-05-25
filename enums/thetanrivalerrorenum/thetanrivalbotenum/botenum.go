package thetanrivalbotenum

type UserType int

const (
	USER_TYPE_TUTORIAL UserType = iota + 1
	USER_TYPE_NEWBIE
	USER_TYPE_OLDBIE
)

type BotRate struct {
	UserType UserType `json:"userType"`
	Rate     float64  `json:"rate"`
}

type RivalBotStatus int

const (
	RivalBotStatusActive    RivalBotStatus = 1
	RivalBotStatusNotActive RivalBotStatus = 2
	RivalBotStatusDeleted   RivalBotStatus = 3
)

type BotBrainType int

const (
	BotBrainTypeObstacleRacing          BotBrainType = 2
	BotBrainTypeHotPotato               BotBrainType = 3
	BotBrainTypeGreenLightRedLight      BotBrainType = 4
	BotBrainTypeLaserDance              BotBrainType = 5
	BotBrainTypeTileFall                BotBrainType = 6
	BotBrainTypeScoringArea             BotBrainType = 7
	BotBrainTypeObstacleRacingOnsenRoad BotBrainType = 13
)

type BrainTypeRival int32

const (
	BrainTypeRivalNone BrainTypeRival = iota
	BrainTypeRivalTutorial
	BrainTypeRivalNovice
	BrainTypeRivalAmateur
	BrainTypeRivalSemiPro
	BrainTypeRivalPro
)
