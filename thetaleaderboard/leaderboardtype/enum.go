package leaderboardtype

import "fmt"

//TODO implement this type
type LeaderboardType string

func (l LeaderboardType) GetCode(seasonId int32) string {
	return fmt.Sprintf(string(l), seasonId)
}
func (l LeaderboardType) GetCountryCode(country string, seasonId int32) string {
	return fmt.Sprintf(string(l), seasonId, country)
}
