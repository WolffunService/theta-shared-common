package dailyquestenum

import "fmt"

type StageRewardStatus int

const (
	RewardInProgress StageRewardStatus = 1 + iota
	RewardPending
	RewardClaimed
)

func (s StageRewardStatus) String() string {
	switch s {
	case RewardInProgress:
		return "In progress"
	case RewardPending:
		return "Pending"
	case RewardClaimed:
		return "Claimed"
	default:
		return fmt.Sprintf("Unknow(%d)", s)
	}
}
