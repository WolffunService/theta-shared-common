package botmodel

import (
	"time"

	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/enums/gamemodeenum"
	"github.com/WolffunService/thetan-shared-common/enums/thetanrivalerrorenum/thetanrivalrank"
	"github.com/WolffunService/thetan-shared-common/models/usermodel"
	"github.com/WolffunService/thetan-shared-common/models/userrankingmodel"
)

type Feature int

const (
	FeatureInGame Feature = 1 + iota
	FeatureLobby
	FeatureLB
	FeatureChatShill
)

func (feature Feature) IsValid() bool {
	return feature == FeatureInGame || feature == FeatureLobby || feature == FeatureLB || feature == FeatureChatShill
}

type BotListRequest struct {
	ArrBotRanklevel []BotRanklevelRequest `json:"arrBotRanklevel"`
}

type BotRanklevelRequest struct {
	GameId       gamemodeenum.GameIds `json:"gameId"`
	Skip         int                  `json:"skip"`
	Numbers      int                  `json:"numbers"`
	RankLevel    *int                 `json:"rankLevel"`
	RegionId     *int                 `json:"regionId"`
	ArrCountry   []string             `json:"arrCountry"`
	Feature      Feature              `json:"feature"`
	UseTimeInSec int64                `json:"useTimeInSec"`
	Battle       int                  `json:"battle"`
}

type BotResponse struct {
	Bots []BotResponseInfo `json:"bots"`
}

type BotResponseInfo struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	CountryCode string `json:"countryCode"`
	RankLevel   int    `json:"rankLevel"`
	Trophy      int    `json:"trophy"`
}

type ListUserRankingRivalResponse struct {
	Success bool                 `json:"success"`
	Data    ThetanRivalsUserInfo `json:"data"`
}

type GetListUserInfoRequest struct {
	UserIds []string `json:"userIds"`
}

type ThetanRivalsUserInfo struct {
	ListUserRanking []TRUserRankingWithIdModel `json:"listUserRanking"`
}

type TRUserRankingWithIdModel struct {
	UserId             string `json:"userId"`
	TRUserRankingModel `json:",inline"`
}

type TRUserRankingModel struct {
	mongodb.DefaultModel `json:"-" bson:",inline"`
	mongodb.DateFields   `json:"-" bson:",inline"`
	Trophy               int                  `json:"trophy" bson:"trophy"`
	Rank                 thetanrivalrank.Rank `json:"rank" bson:"-"`
	Division             int                  `json:"division" bson:"-"`
	Milestone            int                  `json:"milestone" bson:"-"`
	TrophyCurRank        int                  `json:"trophyCurRank" bson:"-"`
	TrophyHighest        int                  `json:"trophyHighest" bson:"trophyHighest"`
	SeasonTrophyHighest  int                  `json:"seasonTrophyHighest" bson:"seasonTrophyHighest"`
	RankingLevelHighest  thetanrivalrank.Rank `json:"rankingLevelHighest" bson:"rankingLevelHighest"`
	SeasonId             int32                `json:"seasonID" bson:"seasonID"`
	Rewards              []RewardInfo         `json:"-" bson:"rewards"`
	ArrRewards           []int                `json:"rewards" bson:"-"`
}

type RewardInfo struct {
	RewardID     int       `json:"rewardID" bson:"rewardID"`
	TimeReceived time.Time `json:"timeReceived" bson:"timeReceived"`
}

type UserRankingRivalsRequest struct {
	ListUserStatistic []UserStatisticRivalsInfo `json:"listUserStatistic"`
}

type UserStatisticRivalsInfo struct {
	UserId                    string `json:"userId"`
	Trophy                    int    `json:"trophy"`
	usermodel.StatisticRivals `json:",inline"`
}

type FakeBotDataThetanRivals struct {
	UserId                    string      `json:"userId"`
	Username                  string      `json:"username"`
	Country                   string      `json:"country"`
	GameIds                   []int       `json:"gameIds" bson:"gameIds"`
	MapGameRankingLevel       map[int]int `json:"mapGameRankingLevel" bson:"mapGameRankingLevel"`
	Trophy                    int         `json:"trophy"`
	usermodel.StatisticRivals `json:",inline"`
}

type UserRankingArenaRequest struct {
	ListUserStatistic []UserStatisticArenaInfo `json:"listUserStatistic"`
}

type UserStatisticArenaInfo struct {
	UserId                    string `json:"userId"`
	Trophy                    int    `json:"trophy"`
	usermodel.PlayerStatistic `json:",inline"`
}

type FakeBotDataThetanArena struct {
	UserId                    string      `json:"userId"`
	Username                  string      `json:"username"`
	Country                   string      `json:"country"`
	GameIds                   []int       `json:"gameIds" bson:"gameIds"`
	MapGameRankingLevel       map[int]int `json:"mapGameRankingLevel" bson:"mapGameRankingLevel"`
	Trophy                    int         `json:"trophy"`
	usermodel.PlayerStatistic `json:",inline"`
}

type ArenaUserRankingWithIdModel struct {
	UserId                       string `json:"userId"`
	userrankingmodel.UserRanking `json:",inline"`
}

type ListUserRankingArenaResponse struct {
	Success bool                `json:"success"`
	Data    ThetanArenaUserInfo `json:"data"`
}

type ThetanArenaUserInfo struct {
	ListUserRanking []ArenaUserRankingWithIdModel `json:"listUserRanking"`
}

type BotInfoBattleEndMessage struct {
	GameId      gamemodeenum.GameIds `json:"gameId"`
	ListBotInfo []BotInfoBattleEnd   `json:"listBotInfo"`
}

type BotInfoBattleEnd struct {
	UserId       string `json:"userId"`
	RankingLevel int    `json:"rankingLevel"`
	Trophy       int    `json:"trophy"`
}

type UserBERequest struct {
	ListUserBEInfo []UserBEInfo `json:"listUserBEInfo"`
}

type UserBEInfo struct {
	UserId       string `json:"userId"`
	BattleRank   int    `json:"battleRank,omitempty"`   /* thetan-riavls */
	InGameMode   int    `json:"inGameMode,omitempty"`   /* thetan-arena */
	BattleResult int    `json:"battleResult,omitempty"` /* thetan-arena */
	Rank         int    `json:"rank,omitempty"`         /* thetan-arena */
}

type UserBEResponse struct {
	Success bool               `json:"success"`
	Data    DataUserBEResponse `json:"data"`
}

type DataUserBEResponse struct {
	ListUserRankingInfo []UserRankingInfo `json:"listUserRankingInfo"`
}

type UserRankingInfo struct {
	UserId       string `json:"userId"`
	RankingLevel int    `json:"rankingLevel"`
	Trophy       int    `json:"trophy"`
}

type CreateNameRequest struct {
	Numbers    int      `json:"numbers"`
	RegionId   int      `json:"regionId"`
	ArrCountry []string `json:"arrCountry"`
}

type CreateNameResponse struct {
	AddedNumbers int `json:"addedNumbers"`
	IncNumbers   int `json:"incNumbers"`
}
