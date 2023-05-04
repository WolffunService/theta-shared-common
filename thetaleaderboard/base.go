package thetaleaderboard

import (
	"context"
	"fmt"

	podium "github.com/WolffunService/theta-shared-common/proto/pb/podium.api.v1"
	"github.com/WolffunService/theta-shared-common/thetalog"
	"google.golang.org/grpc"
)

var podiumClient podium.PodiumClient

func NewClientConn(uri string) *grpc.ClientConn {
	thetalog.Info().Op("NewClientConn").Msgf("Try connect ", uri)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err == nil {
		podiumClient = podium.NewPodiumClient(conn)
	} else {
		thetalog.Err(err).Op("NewClientConn").Send()
	}
	return conn
}

func validatePodium() error {
	if podiumClient == nil {
		return fmt.Errorf("podiumClient is nil")
	}
	return nil
}

func TrackLeaderboardMulti(ctx context.Context, req *podium.UpsertScoreMultiLeaderboardsRequest) error {
	//add queue?
	if err := validatePodium(); err != nil {
		return err
	}
	_, err := podiumClient.UpsertScoreMultiLeaderboards(ctx, req)
	return err
}

func TrackLeaderboard(ctx context.Context, req *podium.UpsertScoreRequest) error {
	if err := validatePodium(); err != nil {
		return err
	}
	_, err := podiumClient.UpsertScore(ctx, req)
	return err
}

func GetTopLeaderboard(ctx context.Context, req *podium.GetTopMembersRequest) *podium.GetTopMembersResponse {
	res, err := podiumClient.GetTopMembers(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

func DeleteMemberLeaderboard(ctx context.Context, req *podium.RemoveMemberRequest) *podium.RemoveMemberResponse {
	res, err := podiumClient.RemoveMember(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}
