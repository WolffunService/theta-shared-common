package battleresultenum

type BattleResult int // #enum: "Battle Result"

const (
	LOSE    BattleResult = -1
	DRAW    BattleResult = 0
	WIN     BattleResult = 1
	UNKNOWN BattleResult = -999
)

func (r BattleResult) GetName() string {
	switch r {
	case LOSE:
		return "LOSE"
	case DRAW:
		return "DRAW"
	case WIN:
		return "WIN"
	}
	return "UNKNOWN"
}
