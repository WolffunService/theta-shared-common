package specialeventenum

type EventStatus int

const (
	LIVE EventStatus = iota
	WARM_UP
	ENDED
)
