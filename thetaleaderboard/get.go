package thetaleaderboard

import (
	"context"
	podium "github.com/WolffunGame/theta-shared-common/proto/pb/podium.api.v1"
)

func GetLeaderboardMember(ctx context.Context, userId, leaderboardType string) (*podium.GetMemberResponse, error) {
	req := &podium.GetMemberRequest{}
	req.MemberPublicId = userId
	req.LeaderboardId = leaderboardType
	if err := validatePodium(); err != nil {
		return nil, err
	}
	res, err := podiumClient.GetMember(ctx, req)
	return res, err
}

func GetTopLeaderboardMember(ctx context.Context, leaderboardId string, pageNumber int32, pageSize int32) (*podium.GetTopMembersResponse, error) {
	req := podium.GetTopMembersRequest{
		LeaderboardId: leaderboardId,
		PageNumber:    pageNumber,
		PageSize:      pageSize,
	}

	if err := validatePodium(); err != nil {
		return nil, err
	}

	res, err := podiumClient.GetTopMembers(ctx, &req)

	return res, err
}

func GetTotalMember(ctx context.Context, leaderboardId string) (*podium.TotalMembersResponse, error) {
	req := podium.TotalMembersRequest{
		LeaderboardId: leaderboardId,
	}

	if err := validatePodium(); err != nil {
		return nil, err
	}

	return podiumClient.TotalMembers(ctx, &req)

}
