package esmatchmaking

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/models/esmodel"
	"github.com/WolffunService/thetan-shared-common/models/esmodel/esinfomodel"
)

type ReconnectMatchMapping struct {
	User               esmodel.UserModel         `json:"user"`
	HeroInfo           esinfomodel.HeroInfoModel `json:"hero_info"`
	InGameInfo         InGameInfo                `json:"in_game_info"`
	MatchId            string                    `json:"match_id"`
	PartyId            string                    `json:"party_id"`
	TrophyMatching     int                       `json:"trophy_matching"`
	TicketCreatedTime  int64                     `json:"ticket_created_time"`
	WaitingSecond      int                       `json:"waiting_second"`
	ErrorCode          int                       `json:"error_code"`
	FindMatchTime      int64                     `json:"find_match_time"`
	MatchFoundTime     int64                     `json:"match_found_time"`
	StartMatchTime     int64                     `json:"start_match_time"`
	CancelMatchTime    int64                     `json:"cancel_match_time"`
	FindMatchErrorTime int64                     `json:"find_match_error_time"`
	EndMatchTime       int64                     `json:"end_match_time"`
	Timestamp          time.Time                 `json:"@timestamp"`
}

func (ReconnectMatchMapping) Index() string {
	return "matchmaking-reconnectmatch"
}
