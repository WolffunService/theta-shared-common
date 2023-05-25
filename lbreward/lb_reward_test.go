package lbreward

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/WolffunService/thetan-shared-common/enums/currencyenum"
	"github.com/WolffunService/thetan-shared-common/enums/leaderboardenum"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardbiz"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardmodel"
	"github.com/WolffunService/thetan-shared-common/lbreward/lbrewardstorage"
	"github.com/WolffunService/thetan-shared-common/leaderboard"
	"github.com/WolffunService/thetan-shared-common/models/currencymodel"
	"github.com/WolffunService/thetan-shared-common/models/leaderboardmodel"

	"github.com/WolffunService/thetan-shared-common/database/mongodb"
)

//NOTE : forward mongodb first
// forward podium if u wanna track to lb reward
// forward leaderboard-go if u wanna check - claim lb reward

const (
	dbName    = "thetan"
	connUrl   = "mongodb://thetan:3c327f016341878ab21b@localhost:27017/thetan?authSource=thetan&replicaSet=thetan-data-rs&readPreference=primary&directConnection=true&ssl=false"
	podiumUri = ":8881"
)

func init() {
	dbConfig := &mongodb.MongoConfig{
		DbName:            dbName,
		ConnectionUrl:     connUrl,
		MaxConnectionPool: 1000,
	}
	//init mongodb
	_, _, _ = mongodb.ConnectMongoWithConfig(dbConfig, nil)

	go InitLB()
}
func InitLB() {
	t := make(chan bool)
	_ = leaderboard.NewClientConn(podiumUri)
	<-t
}

var (
	//var test
	//userId = "631f401e18fd2d77a6e3d08d"
	userId      = "635a610b0a7a598c66d24f2a" // Tuna
	lbPrefixKey = leaderboardenum.LBKeyPrefix("Example29")
	lbID        = 1
	lbKey       = leaderboardmodel.NewLBKeyModel(lbPrefixKey, lbID)
)

func TestTrackLBReward(t *testing.T) {
	totalScore, err := lbrewardbiz.TrackLBReward(context.Background(), userId, lbKey, 140, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(totalScore)
}

func TestTrackLeaderboardMarketplace(t *testing.T) {
	data := map[string]int64{
		"61a98172003b2a9503350a81": 1,
		"61a91128fbb1a8b08bb1fbba": 2,
		"61a9692efbb1a8b08bb1fbd3": 3,
		"61ab6275003b2a9503350b2c": 4,
		"61a98288003b2a9503350a82": 5,
		"61b13ee4003b2a9503350c9b": 6,
		"627b99200e37a07bd01a5994": 7,
		"61af8604003b2a9503350c28": 8,
		"62971a1427a535109b15e079": 9,
		"61a8ac6cfbb1a8b08bb1fb92": 10,
		"61ab00e1003b2a9503350b03": 11,
		"61a69c46fbb1a8b08bb1faed": 12,
		"61aa0be5003b2a9503350ab0": 13,
		"61acbcdf003b2a9503350b82": 14,
		"619b0d99fe775c92cca3958d": 15,
		"61a8431dfbb1a8b08bb1fb62": 16,
		"61aef97a003b2a9503350bfb": 17,
		"61ac9430003b2a9503350b77": 18,
		"61a651ecfbb1a8b08bb1fad8": 19,
		"61c9a157e89e8850b84cf3fe": 200, // Trung

		"619a0100fe775c92cca3956e": 0,
		"6199c4cefe775c92cca39560": 0,

		"619f068f379ffd1b14a94d5c": 18, // Oanh
	}

	//groupId := 1
	seasonId := 2
	prefix := "Example121"
	lbKey = leaderboardmodel.NewLBKeyModel(leaderboardenum.LBKeyPrefix(prefix), seasonId)
	for userIdVjp, score := range data {
		totalScore, err := lbrewardbiz.TrackLBReward(context.Background(), userIdVjp, lbKey, score, nil)
		if err != nil {
			t.Error(err)
			break
		}
		t.Log(userId, totalScore)
	}
}

func TestCheckLBReward(t *testing.T) {
	req := &lbrewardmodel.ReqClaimLBReward{
		UserID:      "635a610b0a7a598c66d24f2a",
		LBRewardKey: lbPrefixKey,
	}
	scoreModel, err := lbrewardbiz.CheckReward(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	t.Log(scoreModel)
	t.Log(scoreModel.ScoreData)
	t.Log(scoreModel.LBData)
}

// TestGetLBInfo @san them
// func TestGetLBInfo(t *testing.T) {
// 	c := lbrewardapi.NewLeaderboardAPIClient("http://localhost:8080/internal")
// 	reqGetLBData := &leaderboardmodel.RequestGetLBFullData{
// 		UserID:     "", // empty thì mine=nil
// 		LBKey:      lbKey.String(),
// 		PageSize:   50, // nên >= total player
// 		PageNumber: 1,  // cứ lấy bằng 1
// 	}
// 	res, err := c.GetLeaderboardFullData(reqGetLBData)
// 	fmt.Println("res", res)
// 	if res.Entries != nil {
// 		for _, ele := range res.Entries {
// 			fmt.Println("ele", ele)
// 		}
// 	}
// 	fmt.Println("res.MineLeaderboard", res.MineLeaderboard)
// 	fmt.Println("err", err)
// }

func TestClaimLBReward(t *testing.T) {
	req := &lbrewardmodel.ReqClaimLBReward{
		UserID:      userId,
		LBRewardKey: lbPrefixKey,
	}
	rewardHandler := RewardTestHandler{}
	reward, err := lbrewardbiz.ClaimReward[RewardTest](context.Background(), req, rewardHandler)
	if err != nil {
		t.Error(err)
	}
	t.Log(reward)
}

func TestGetLBReward(t *testing.T) {
	ctx := context.Background()
	req := lbrewardmodel.RequestGetLBReward{
		UserID:     "",
		LBKey:      lbKey.String(),
		PageNumber: 1,
		PageSize:   50,
	}

	handler := RewardTestHandler{}

	res, err := lbrewardbiz.GetLeaderboardReward[RewardTest](ctx, req, handler)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Entries[0].ProfileInfo)
}

func TestIndex(t *testing.T) {
	lbrewardstorage.CreateIndexScoreModelLB(context.Background())
}

// model test
type RewardTest struct {
	THC currencymodel.SystemCurrency
}
type RewardTestHandler struct {
}

func (RewardTestHandler) ExcReward(ctx context.Context, data *leaderboardmodel.LeaderboardDataModelFullData) (RewardTest, error) {
	if data == nil {
		panic("something went wrong")
	}
	jsonString, err := json.Marshal(data)
	if err == nil {
		fmt.Println(string(jsonString))
	}
	fmt.Println("execute reward ... ")
	return RewardTest{
		THC: currencymodel.NewSystemCurrency(1, currencyenum.THC),
	}, nil
}

func (RewardTestHandler) AfterClaim(ctx context.Context) error {
	fmt.Println("after claim action")
	return nil
}

func (RewardTestHandler) BeforeClaim(ctx context.Context) error {
	fmt.Println("before claim action")
	return nil
}
