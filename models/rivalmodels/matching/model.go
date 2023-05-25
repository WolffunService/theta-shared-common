package matching

import (
	"github.com/WolffunService/thetan-shared-common/enums/thetanrivalerrorenum/thetanrivalrank"
	"github.com/WolffunService/thetan-shared-common/proto/coreproto"
)

type PlayerInfoResponse struct {
	PlayerId          string                       `json:"playerId"`
	IsValid           bool                         `json:"isValid"`
	MinionId          string                       `json:"minionId"`
	SkinId            int32                        `json:"skinId"`
	CosmeticUsingInfo coreproto.CosmeticUsingProto `json:"cosmeticUsingInfo"`
	MinionLevel       int32                        `json:"minionLevel"`
	MinionStatus      int32                        `json:"minionStatus"`
	IsBannedForever   bool                         `json:"isBannedForever"`
	IsBannedFindMatch bool                         `json:"isBannedFindMatch"`
	TrophyMatching    int32                        `json:"trophyMatching"`
	ErrorCode         int32                        `json:"errorCode"`
	BehaviorPoint     int32                        `json:"behaviorPoint"`
	BattleCount       int32                        `json:"battleCount"`
	IsHeroNFT         bool                         `json:"isHeroNFT"`
	TrophiesRank      int32                        `json:"trophiesRank"`
	MatchSearchType   int32                        `json:"matchSearchType"`
	AvatarId          int32                        `json:"avatarId"`
	FrameId           int32                        `json:"frameId"`
	UserName          string                       `json:"userName"`
	NameColorId       int32                        `json:"nameColorId"`
	PlayerScore       float64                      `json:"playerScore"`
	MapIDs            []int32                      `json:"mapIDs"`
	Rank              thetanrivalrank.Rank         `json:"rank"`
	RoundPlayed       map[string]int32             `json:"roundPlayed"`
}
