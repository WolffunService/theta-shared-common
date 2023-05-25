package minionenum

import "fmt"

type MinionSource int

const (
	MSOnBoarding    MinionSource = 1001
	MSRankingReward MinionSource = 1002
	MSBoxOpen       MinionSource = 1003
	MSIAPShop       MinionSource = 1004
	MSSeasonPass    MinionSource = 1005

	// Admin change source 1900 -> 1999
	MSAdminSend MinionSource = 1900
)

var sourceName = map[MinionSource]string{
	MSOnBoarding:    "On-boarding",
	MSRankingReward: "Ranking reward",
	MSBoxOpen:       "Open box",
	MSIAPShop:       "IAP shop",

	MSAdminSend: "Admin send",
}

func (source MinionSource) String() string {
	if name, found := sourceName[source]; found {
		return name
	}

	return fmt.Sprintf("Unknown (%d)", source)
}
