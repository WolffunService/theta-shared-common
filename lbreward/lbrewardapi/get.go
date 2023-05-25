package lbrewardapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"
	"github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"
)

const (
	endpointLeaderboardGetFullData       = "leaderboard/fullData/list"
	endpointLeaderboardGetFullDataMember = "leaderboard/fullData/member"
)

var lbGoServiceInternalAddr string

func Init(internalAddress string) error {
	if len(internalAddress) == 0 {
		return errors.New("[error] lbrewardapi: please specify leaderboard-go internal address")
	}

	lbGoServiceInternalAddr = internalAddress

	return nil
}

func GetLeaderboardFullData(req *leaderboardmodel.RequestGetLBFullData) (*leaderboardmodel.LeaderBroadResponseFullData, error) {

	data, statusCode, err := newGetRequestParams(lbGoServiceInternalAddr, endpointLeaderboardGetFullData, req)
	if err != nil {
		fmt.Printf("GetLeaderboardFullData error %s", err)
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request fail %v", statusCode))
	}

	var res *lbrewardmodel.LeaderBroadResponseFullDataWrap
	errParseJson := json.Unmarshal(data, &res)
	if errParseJson != nil {
		return nil, errParseJson
	}

	if res.Success == false || res.Data == nil {
		return nil, errors.New(fmt.Sprintf("error when get leaderboardReward from leaderboard-go. res.Success: %v, res.Data: %v", res.Success, res.Data))
	}

	return res.Data, nil
}

func GetMemberLeaderboardFullData(req *leaderboardmodel.RequestGetLBFullData) (*leaderboardmodel.LeaderboardDataModelFullData, error) {
	resp, err := gETDynamic(
		fmt.Sprintf("%s/%s", lbGoServiceInternalAddr, endpointLeaderboardGetFullDataMember),
		convertRawQuery(*req),
		make(map[string]string),
		"",
	)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request fail %v", resp.StatusCode()))
	}

	var res *lbrewardmodel.LeaderboardDataModelFullDataWrap
	errParseJson := json.Unmarshal(resp.Body(), &res)
	if errParseJson != nil {
		return nil, errParseJson
	}
	return res.Data, nil
}
