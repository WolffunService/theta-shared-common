package gamereleaseenum

type GameRelease string

const (
	Unreleased   GameRelease = ""
	GRBeta       GameRelease = "beta"
	GRSoftLaunch GameRelease = "soft-launch"
)
