package rivalitemenum

type DropGroup string

const (
	DGUnknown    DropGroup = "unknown"
	DGGold       DropGroup = "gold"
	DGEnhancer   DropGroup = "enhancer"
	DGMinion     DropGroup = "minion"
	DGProfile    DropGroup = "profile"
	DGFullEvolve DropGroup = "full-evolve"
	DGHalfEvolve DropGroup = "half-evolve"
	DGSPoint     DropGroup = "season-point"
)
