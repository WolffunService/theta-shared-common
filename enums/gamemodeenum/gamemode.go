package gamemodeenum

// --------------
//GameMode
type GameMode int

const (
	TEAM_MODE              GameMode = iota
	BATTLE_ROYALE_MODE_OLD          //12
	BATTLE_ROYALE_MODE_OLD_DUAL
	BATTLE_ROYALE_MODE_NEW //42
)

func (g GameMode) IsTeamMode() bool {
	return g == TEAM_MODE
}

//InGameMode
type InGameMode int

const (
	TEAM_COLLECT_STAR        InGameMode = 0
	SOLO_SURVIVAL            InGameMode = 1
	DUAL_SURVIVAL            InGameMode = 2
	TEAM_COLLECT_STAR_4_VS_4 InGameMode = 3
	KING                     InGameMode = 5
	DEATH_MATCH              InGameMode = 6
	DEATH_MATCH_3_VS_3       InGameMode = 7
	FLAG                     InGameMode = 8
	TOWER                    InGameMode = 9
	BATTLE_ROYALE            InGameMode = 12
	SQUAD_BATTLE_ROYALE      InGameMode = 13
	DUO_BATTLE_ROYALE        InGameMode = 14
	TRIO_BATTLE_ROAYLE       InGameMode = 15
	THETAN_RIVALS            InGameMode = 20
	NONE_MODE                InGameMode = -1
)

// GetModeType return 0 if this is team mode or 1 if solo mode
func (i InGameMode) GetModeType() int {
	if i == TEAM_COLLECT_STAR_4_VS_4 ||
		i == TEAM_COLLECT_STAR ||
		i == DEATH_MATCH ||
		i == DEATH_MATCH_3_VS_3 ||
		i == TOWER ||
		i == FLAG {
		return 0
	}

	return 1
}

func (i InGameMode) GetGameMode() GameMode {
	gameMode := TEAM_MODE
	switch i {
	case SOLO_SURVIVAL:
		gameMode = BATTLE_ROYALE_MODE_OLD
	case DUAL_SURVIVAL:
		gameMode = BATTLE_ROYALE_MODE_OLD_DUAL
	}
	return gameMode
}

func (inGameMode InGameMode) IsSoloMode() bool {
	return inGameMode == SOLO_SURVIVAL
}

func (inGameMode InGameMode) IsDualMode() bool {
	return inGameMode == DUAL_SURVIVAL
}

func (inGameMode InGameMode) IsTeamMode() bool {
	if inGameMode == TEAM_COLLECT_STAR_4_VS_4 ||
		inGameMode == TEAM_COLLECT_STAR ||
		inGameMode == DEATH_MATCH ||
		inGameMode == DEATH_MATCH_3_VS_3 ||
		inGameMode == TOWER ||
		inGameMode == FLAG {
		return true
	}
	return false
}

func (inGameMode InGameMode) IsBRMode() bool {
	return inGameMode == SOLO_SURVIVAL || inGameMode == DUAL_SURVIVAL
}

type GameIds int

const (
	GameIdsThetanArena  GameIds = 1
	GameIdsThetanRivals GameIds = 2
)
