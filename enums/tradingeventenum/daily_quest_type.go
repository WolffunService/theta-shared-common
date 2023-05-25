package tradingeventenum

type DailyQuestType int

const (
	_ DailyQuestType = iota
	DQT_SELLER
	DQT_UNBOX
)

type ClaimStatus int

const (
	LOCK ClaimStatus = iota
	CAN_NOT
	CAN
	CLAIMED
)
