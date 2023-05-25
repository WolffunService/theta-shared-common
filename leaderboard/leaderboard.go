package leaderboard

import (
	"context"
	"fmt"
	"log"

	podium "github.com/WolffunService/thetan-shared-common/proto/podium"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var podiumClient podium.PodiumClient

func NewClientConn(uri string) *grpc.ClientConn {
	fmt.Println("Try connect ", uri)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err == nil {
		podiumClient = podium.NewPodiumClient(conn)
	} else {
		log.Fatal(err)
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

func GetTopLeaderboard(ctx context.Context, req *podium.GetTopMembersRequest) (*podium.GetTopMembersResponse, error) {
	if err := validatePodium(); err != nil {
		return nil, err
	}

	topMembers, err := podiumClient.GetTopMembers(ctx, req)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}
	return topMembers, nil
}

func GetMineLeaderboard(ctx context.Context, publicID, lbType string) (*podium.GetMemberResponse, error) {
	request := &podium.GetMemberRequest{}
	request.LeaderboardId = lbType
	request.MemberPublicId = publicID
	return getMineLeaderboard(ctx, request)
}

func getMineLeaderboard(ctx context.Context, req *podium.GetMemberRequest) (*podium.GetMemberResponse, error) {
	if err := validatePodium(); err != nil {
		return nil, err
	}

	member, err := podiumClient.GetMember(ctx, req)
	if err != nil {
		//log.Println(err)
		return nil, err
	}
	return member, nil
}

func GetTotalMember(ctx context.Context, req *podium.TotalMembersRequest) (*podium.TotalMembersResponse, error) {
	if err := validatePodium(); err != nil {
		return nil, err
	}

	topMembers, err := podiumClient.TotalMembers(ctx, req)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}
	return topMembers, nil
}

func DeleteMemberLeaderboard(ctx context.Context, req *podium.RemoveMemberRequest) *podium.RemoveMemberResponse {
	res, err := podiumClient.RemoveMember(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

func TrackLB(ctx context.Context, userId string, score int, leaderboards ...string) error {
	req := &podium.UpsertScoreMultiLeaderboardsRequest{}
	req.MemberPublicId = userId

	scoreMultiChange := &podium.UpsertScoreMultiLeaderboardsRequest_ScoreMultiChange{}
	scoreMultiChange.Score = float64(score)
	scoreMultiChange.Leaderboards = leaderboards
	req.ScoreMultiChange = scoreMultiChange

	return TrackLeaderboardMulti(ctx, req)
}
