package thetaleaderboard

import (
	"context"

	podium "github.com/WolffunService/thetan-shared-common/proto/pb/podium.api.v1"
)

func TrackLB(ctx context.Context, userId string, score int32, leaderboards ...string) error {
	req := &podium.UpsertScoreMultiLeaderboardsRequest{}
	req.MemberPublicId = userId

	scoreMultiChange := &podium.UpsertScoreMultiLeaderboardsRequest_ScoreMultiChange{}
	scoreMultiChange.Score = float64(score)
	scoreMultiChange.Leaderboards = leaderboards
	req.ScoreMultiChange = scoreMultiChange

	return TrackLeaderboardMulti(ctx, req)
}
