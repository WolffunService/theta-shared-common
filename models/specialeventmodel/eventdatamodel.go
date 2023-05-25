package specialeventmodel

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"

	"github.com/WolffunService/theta-shared-common/database/mongodb"
	"github.com/WolffunService/theta-shared-common/enums/currencyenum"
	"github.com/WolffunService/theta-shared-common/enums/gamemodeenum"
	"github.com/WolffunService/theta-shared-common/models/currencymodel"

	eventtype "github.com/WolffunService/theta-shared-common/enums/specialeventenum"
)

func (TournamentEventData) CollectionName() string {
	return "ServerSpecialEventData"
}
func (ListTournamentEventData) CollectionName() string {
	return "ServerSpecialEventData"
}

// List
type ListTournamentEventData struct {
	mongodb.DefaultModel      `bson:",inline"`
	mongodb.DateFields        `bson:",inline"`
	EventType                 eventtype.EventType         `json:"eventType"binding:"enum=eventtype.EventType" bson:"eventType"` //pk
	Status                    bool                        `json:"status" bson:"status"`
	CreatorId                 string                      `json:"creatorId" bson:"creatorId"`
	EventId                   int                         `json:"eventId" bson:"eventId"`
	CronExpression            string                      `json:"cronExpression" bson:"cronExpression"`             //cron job list event //ex : "0 0 * * *" at 00h00 UTC
	CurTournamentEventId      int                         `json:"curTournamentEventId" bson:"curTournamentEventId"` //increase when start new event success - for get next configuration tournament event
	StartIn                   int64                       `json:"startIn" bson:"startIn"`
	Duration                  int64                       `json:"duration" bson:"duration"`
	ProcessingTime            int64                       `json:"processingTime" bson:"processingTime"`
	ListDailyEventData        map[int]DailyEventData      `json:"listDailyEventData" bson:"listDailyEventData"`               //maybe size is 1
	ListDailyBattleConfig     map[int]DailyBattleConfig   `json:"listDailyBattleConfig" bson:"listDailyBattleConfig"`         //maybe size is 1
	ListOfRequiredHeroes      map[int][]HeroEventModel    `json:"listOfRequiredHeroes" bson:"listOfRequiredHeroes"`           //maybe size is 1
	ListOfRequiredHeroesDaily map[int][]RequiredHeroes    `json:"listOfRequiredHeroesDaily" bson:"listOfRequiredHeroesDaily"` //maybe size is 1
	ListEventURLConfig        map[int]EventImageURLConfig `json:"listEventURLConfig" bson:"listEventURLConfig"`               //maybe size is 1
	ListEventRewardConfig     map[int]EventRewardConfig   `json:"listEventRewardConfig" bson:"listEventRewardConfig"`         //maybe size is 1
	//dung chung
	RankRewardRate  map[int]float64 `json:"rankRewardRate" bson:"rankRewardRate"`
	TrophyRequired  int             `json:"trophyRequired" bson:"trophyRequired"`
	PriceTicketGTHC float64         `json:"priceTicket" bson:"priceTicket"`
	DailyShareRate  float64         `json:"dailyShareRate" bson:"dailyShareRate"`
	WeeklyShareRate float64         `json:"weeklyShareRate" bson:"weeklyShareRate"`
	RangeShare      map[int]float64 `json:"rangeShare" bson:"rangeShare"`
}

func (l ListTournamentEventData) GetEventId() int {
	return l.EventId
}

func (t ListTournamentEventData) GetEventType() eventtype.EventType {
	return t.EventType
}

func (l ListTournamentEventData) IsActive() bool {
	return l.Status
}

func (l ListTournamentEventData) Ended() bool {
	return !l.IsActive()
}

func (l *ListTournamentEventData) SetEventId(eventId int) {
	l.EventId = eventId
}

// genarate new tournament event 0 ~ 6
func (l ListTournamentEventData) GetNewTournamentEvent() (*TournamentEventData, error) {
	curId := l.CurTournamentEventId % 7 // 0~6

	event := &TournamentEventData{}
	event.EventType = eventtype.TOURNAMENT_EVENT
	event.IsDailyEvent = true
	//Start in
	{
		//first event
		if l.CurTournamentEventId == 0 {
			event.StartIn = l.StartIn
		} else {
			//another event
			event.StartIn = time.Now().UTC().Unix()
		}
	}
	//Duration time
	event.Duration = l.Duration
	//Processing time
	event.ProcessingTime = l.ProcessingTime
	//Daily Event Data
	{
		if l.ListDailyEventData == nil || len(l.ListDailyEventData) <= 0 {
			return nil, fmt.Errorf("l.ListDailyEventData in wrong format")
		}
		if dailyData, ok := l.ListDailyEventData[curId]; ok {
			event.DailyEventData = &dailyData
		} else {
			dailyData = l.ListDailyEventData[0]
			event.DailyEventData = &dailyData
		}
	}

	//Daily battle config
	{
		if l.ListDailyBattleConfig == nil || len(l.ListDailyBattleConfig) <= 0 {
			return nil, fmt.Errorf("l.ListDailyBattleConfig in wrong format")
		}
		if dailyBattleConfig, ok := l.ListDailyBattleConfig[curId]; ok {
			event.DailyBattleConfig = dailyBattleConfig
		} else {
			dailyBattleConfig = l.ListDailyBattleConfig[0]
			event.DailyBattleConfig = dailyBattleConfig
		}
	}
	//Hero Require
	{
		if l.ListOfRequiredHeroes == nil || len(l.ListOfRequiredHeroes) <= 0 {
			return nil, fmt.Errorf("l.ListOfRequiredHeroes in wrong format")
		}
		if heroRequired, ok := l.ListOfRequiredHeroes[curId]; ok {
			event.ListOfRequiredHeroes = heroRequired
		} else {
			event.ListOfRequiredHeroes = l.ListOfRequiredHeroes[0]
		}

	}

	//Hero Require Daily
	{
		if l.ListOfRequiredHeroesDaily == nil || len(l.ListOfRequiredHeroesDaily) <= 0 {
			return nil, fmt.Errorf("l.ListOfRequiredHeroesDaily in wrong format")
		}
		if heroRequiredDaily, ok := l.ListOfRequiredHeroesDaily[curId]; ok {
			event.ListOfRequiredHeroesDaily = heroRequiredDaily
		} else {
			event.ListOfRequiredHeroesDaily = l.ListOfRequiredHeroesDaily[0]
		}
	}

	//Hero Require Daily
	{
		if l.ListEventURLConfig == nil || len(l.ListEventURLConfig) <= 0 {
			return nil, fmt.Errorf("l.ListEventURLConfig in wrong format")
		}
		if urlConfig, ok := l.ListEventURLConfig[curId]; ok {
			event.EventImageURLConfig = urlConfig
		} else {
			event.EventImageURLConfig = l.ListEventURLConfig[0]
		}
	}

	//EventRewardConfig
	{
		if l.ListEventRewardConfig == nil || len(l.ListEventRewardConfig) <= 0 {
			return nil, fmt.Errorf("l.ListEventRewardConfig in wrong format")
		}
		if rewardConfig, ok := l.ListEventRewardConfig[curId]; ok {
			event.EventRewardConfig = rewardConfig
		} else {
			event.EventRewardConfig = l.ListEventRewardConfig[0]
		}
	}

	event.RangeShare = l.RangeShare
	event.RankRewardRate = l.RankRewardRate
	event.TrophyRequired = l.TrophyRequired
	event.PriceTicketGTHC = l.PriceTicketGTHC
	event.DailyShareRate = l.DailyShareRate
	event.WeeklyShareRate = l.WeeklyShareRate
	return event, nil
}

// Tournament
type TournamentEventData struct {
	mongodb.DefaultModel `bson:",inline"`
	mongodb.DateFields   `bson:",inline"`
	CreatorId            string `json:"creatorId" bson:"creatorId"`
	EventId              int    `json:"eventId" bson:"eventId"`

	//admin input
	EventType                 eventtype.EventType `json:"eventType"binding:"enum=eventtype.EventType" bson:"eventType"` //pk
	StartIn                   int64               `json:"startIn" bson:"startIn"`
	Duration                  int64               `json:"duration" bson:"duration"`
	ProcessingTime            int64               `json:"processingTime" bson:"processingTime"`
	DailyEventData            *DailyEventData     `json:"dailyEventData" bson:"dailyEventData"`
	ListOfRequiredHeroes      []HeroEventModel    `json:"listOfRequiredHeroes" bson:"listOfRequiredHeroes"`
	ListOfRequiredHeroesDaily []RequiredHeroes    `json:"listOfRequiredHeroesDaily" bson:"listOfRequiredHeroesDaily"`
	EventImageURLConfig       `json:",inline" bson:",inline"`
	EventRewardConfig         `json:",inline" bson:",inline"`
	DailyBattleConfig         DailyBattleConfig            `json:"dailyBattleConfig" bson:"dailyBattleConfig"`
	DailyRewardPool           currencymodel.SystemCurrency `json:"dailyRewardPool" bson:"dailyRewardPool"`
	TournamentEventPool       currencymodel.SystemCurrency `json:"tournamentEventPool" bson:"tournamentEventPool"`
	TournamentEventTicket     int32                        `json:"tournamentEventTicket" bson:"tournamentEventTicket"`
	DailyEventTicket          map[int]DailyTicket          `json:"-" bson:"dailyEventTicket"`
	BaseRewardRate            map[int]float64              `json:"baseRewardRate" bson:"baseRewardRate"` //not available when isDailyEvent == true
	RankRewardRate            map[int]float64              `json:"rankRewardRate" bson:"rankRewardRate"`
	TrophyRequired            int                          `json:"trophyRequired" bson:"trophyRequired"`
	PriceTicketGTHC           float64                      `json:"priceTicket" bson:"priceTicket"`
	DailyShareRate            float64                      `json:"dailyShareRate" bson:"dailyShareRate"`
	WeeklyShareRate           float64                      `json:"weeklyShareRate" bson:"weeklyShareRate"`
	IsDailyEvent              bool                         `json:"isDailyEvent" bson:"isDailyEvent"`
	RangeShare                map[int]float64              `json:"rangeShare" bson:"rangeShare"`
	Version                   int32                        `json:"version" bson:"version"`
}

type EventImageURLConfig struct {
	BannerEventUrl        string `json:"bannerEventUrl" bson:"bannerEventUrl"`
	BannerMobileEventUrl  string `json:"bannerMobileEventUrl" bson:"bannerMobileEventUrl"`
	UnityBannerEventUrl   string `json:"unityBannerEventUrl" bson:"unityBannerEventUrl"`
	UnityGameModeEventUrl string `json:"unityGameModeEventUrl" bson:"unityGameModeEventUrl"`
}

type EventRewardConfig struct {
	MinDailyPoolGTHG   float64 `json:"minDailyPoolgTHG" bson:"minDailyPoolgTHG"`
	MinWeeklyPoolGTHG  float64 `json:"minWeeklyPoolgTHG" bson:"minWeeklyPoolgTHG"`
	ExchangeRateSample float64 `json:"exchangeRateSample" bson:"exchangeRateSample"`
	BonusGTHGEachUser  float64 `json:"bonusGTHGEachUser" bson:"bonusGTHGEachUser"`
}

type TournamentEventDataClient struct {
	EventId                   int                       `json:"eventId"`
	EventType                 eventtype.EventTypeString `json:"eventType"`
	StartIn                   int64                     `json:"eventStartIn"`
	EndIn                     int64                     `json:"eventEndIn"`
	ProcessingTime            int64                     `json:"processingTime"`
	*DailyEventDataClient     `json:",inline"`
	ListOfRequiredHeroes      []HeroEventModel             `json:"listOfRequiredHeroes"`
	ListOfRequiredHeroesDaily []RequiredHeroes             `json:"listOfRequiredHeroesDaily"`
	BannerEventUrl            string                       `json:"bannerEventUrl"`
	BannerMobileEventUrl      string                       `json:"bannerMobileEventUrl"`
	UnityBannerEventUrl       string                       `json:"unityBannerEventUrl"`
	UnityGameModeEventUrl     string                       `json:"unityGameModeEventUrl"`
	DailyBattleConfigClient   DailyBattleConfigClient      `json:"dailyBattles"`
	IsJoined                  bool                         `json:"isJoined"`
	TrophyRequired            int                          `json:"trophyRequired"`
	PriceTicket               currencymodel.SystemCurrency `json:"priceTicket"`
	IsDailyEvent              bool                         `json:"isDailyEvent"`
}

type DailyEventData struct {
	DailyEventId        int    `json:"dailyEventId" bson:"dailyEventId"`
	DailyCronExpression string `json:"dailyCronExpression" bson:"dailyCronExpression"`
	DailyStartIn        int64  `json:"dailyStartIn" bson:"dailyStartIn"` //time now in cronjob
	DailyDuration       int64  `json:"dailyDuration" bson:"dailyDuration"`
}

type DailyEventDataClient struct {
	DailyEventId      int                       `json:"dailyEventId"`
	DailyEventType    eventtype.EventTypeString `json:"dailyEventType"`
	DailyEventStartIn int64                     `json:"dailyEventStartIn"`
	DailyEventEndIn   int64                     `json:"dailyEventEndIn"`
}

type RequiredHeroes struct {
	HeroEventModel []HeroEventModel `json:"heroEventModels" bson:"heroEventModels"`
}

type HeroEventModel struct {
	HeroTypeId  int `json:"heroTypeId" bson:"heroTypeId"`
	SkinId      int `json:"skinId" bson:"skinId"`
	HeroLevel   int `json:"heroLevel" bson:"heroLevel"`
	TrophyLevel int `json:"trophyLevel" bson:"trophyLevel"`
}

type DailyBattleConfig struct {
	DMDailyBattleConfig   int `json:"DMDailyBattle" bson:"DMDailyBattle"`
	SSDailyBattleConfig   int `json:"SSDailyBattle" bson:"SSDailyBattle"`
	TWDailyBattleConfig   int `json:"TWDailyBattle" bson:"TWDailyBattle"`
	SoloDailyBattleConfig int `json:"soloDailyBattle" bson:"soloDailyBattle"`
	DualDailyBattleConfig int `json:"dualDailyBattle" bson:"dualDailyBattle"`
}

type DailyTicket struct {
	Total int `json:"total" bson:"total"`
}

func (d *DailyBattleConfig) GetDailyBattleLimitByMode(mode gamemodeenum.InGameMode) int {
	switch mode {
	case gamemodeenum.TEAM_COLLECT_STAR_4_VS_4:
		return d.SSDailyBattleConfig
	case gamemodeenum.DEATH_MATCH:
		return d.DMDailyBattleConfig
	case gamemodeenum.TOWER:
		return d.TWDailyBattleConfig
	case gamemodeenum.SOLO_SURVIVAL:
		return d.SoloDailyBattleConfig
	case gamemodeenum.DUAL_SURVIVAL:
		return d.DualDailyBattleConfig
	}
	return 0
}

type DailyBattleConfigClient struct {
	DailyBattleConfig  `json:",inline"`
	CurDMDailyBattle   int `json:"curDMDailyBattle"`
	CurSSDailyBattle   int `json:"curSSDailyBattle"`
	CurTWDailyBattle   int `json:"curTWDailyBattle"`
	CurSoloDailyBattle int `json:"curSoloDailyBattle"`
	CurDualDailyBattle int `json:"curDualDailyBattle"`
}

func (l TournamentEventData) GetID() interface{} {
	return l.ID
}

func (t TournamentEventData) GetEventId() int {
	return t.EventId
}

func (t TournamentEventData) GetEventType() eventtype.EventType {
	return t.EventType
}

func (s *TournamentEventData) EndIn() int64 {
	return s.StartIn + s.Duration
}

func (d *DailyEventData) EndIn() int64 {
	return d.DailyStartIn + d.DailyDuration
}

func (s *TournamentEventData) IsActive() bool {
	curTimeStamp := time.Now().UTC().Unix()
	return curTimeStamp >= s.StartIn && s.EndIn()-10 > curTimeStamp //10s for cronjob
}

func (s *TournamentEventData) Ended() bool {
	curTimeStamp := time.Now().UTC().Unix()
	return s.EndIn()-10 < curTimeStamp //10s for cronjob
}

func (s *TournamentEventData) InProcesingTime() bool {
	curTimeStamp := time.Now().UTC().Unix()
	return s.EndIn()+s.ProcessingTime >= curTimeStamp
}
func (s *TournamentEventData) SetEventId(eventId int) {
	s.EventId = eventId
}

func (d *DailyEventData) IsActive() bool {
	curTimeStamp := time.Now().UTC().Unix()
	return curTimeStamp >= d.DailyStartIn && d.EndIn()-10 > curTimeStamp
}

func (s *TournamentEventData) GetGTHGRateReward(rank int) float64 {
	keys := sortKeyIntMap(s.BaseRewardRate)
	for _, k := range keys {
		if rank <= k {
			return s.BaseRewardRate[k]
		}
	}
	return 0
}

// get range share rank
func (s *TournamentEventData) GetRangeShare(rank int) float64 {
	keys := sortKeyIntMap(s.RangeShare)
	for _, k := range keys {
		if rank <= k {
			return s.RangeShare[k]
		}
	}
	return 0
}

// get max rank in rank (next rank config minus 1)
func (s *TournamentEventData) GetMaxRiR(rank int) int {
	keys := sortKeyIntMap(s.RangeShare)
	for i, k := range keys {
		if rank <= k && i <= len(keys)-1 { //in range
			if i < 1 {
				return k
			}
			return k - keys[i-1]
		}
	}
	return 1
}

// get rank in rank
func (s *TournamentEventData) GetRiR(rank int) int {
	keys := sortKeyIntMap(s.RangeShare)
	for i, k := range keys {
		if rank <= k {
			if i-1 < 0 {
				return 1
			}
			return rank - (keys[i-1])
		}
	}
	return 1
}

func (s *TournamentEventData) GetBonusByRankingLevel(trophyLevel int) float64 {
	return s.RankRewardRate[trophyLevel]
}

const (
	MIN_DAILY_POOL_GTHG  = 300
	MIN_WEEKLY_POOL_GTHG = 500

	EXCHAGE_RATE_SAMPLE = 0.01
)

func (s *TournamentEventData) GetGTHGReward(isDaily bool, rank, trophyLevel, dailyEventId int) (RewardsSE, RewardsAnalytic) {
	isFreeEvent := s.PriceTicketGTHC == 0
	var pool currencymodel.SystemCurrency

	if s.IsDailyEvent {
		isDaily = !isDaily //swap slot daily
		dailyEventId = 0   // TODO check lai khi event daily thay doi
	}

	if isDaily {
		if isFreeEvent {
			pool = s.DailyRewardPool
		} else {
			totalTicket := 0
			if !(s.DailyEventTicket == nil || len(s.DailyEventTicket) == 0) {
				totalTicket = s.DailyEventTicket[dailyEventId].Total
			}
			base := float64(totalTicket) * s.PriceTicketGTHC * s.DailyShareRate * s.ExchangeRateSample
			basePool := math.Max(s.MinDailyPoolGTHG, base)
			pool = currencymodel.ConvertFloatToSystemCurrency(basePool, currencyenum.GTHG)
		}

	} else {
		if s.IsDailyEvent {
			pool = currencymodel.ConvertFloatToSystemCurrency(0, currencyenum.GTHG) //only exist daily pool
		} else {
			if isFreeEvent {
				pool = s.TournamentEventPool
			} else {
				base := float64(s.TournamentEventTicket) * s.PriceTicketGTHC * s.WeeklyShareRate * s.ExchangeRateSample
				basePool := math.Max(s.MinWeeklyPoolGTHG, base)
				pool = currencymodel.ConvertFloatToSystemCurrency(basePool, currencyenum.GTHG)
			}
		}
	}
	reward := RewardsSE{}
	rewardAnalytic := RewardsAnalytic{
		GTHGRateReward: s.GetGTHGRateReward(rank),
		RankBonusRate:  s.GetBonusByRankingLevel(trophyLevel),
		SharedRate:     s.GetRangeShare(rank),
	}

	var gthgBaseReward, bonusTHGEachUser float64
	if s.IsDailyEvent {
		maxRiR := float64(s.GetMaxRiR(rank))
		RR := float64(s.GetRiR(rank))
		//rankShare := s.GetRangeShare(rank)
		temp := (0.8 / maxRiR) + (0.2*(maxRiR-RR+float64(1)))/(maxRiR*(float64(1)+maxRiR)/float64(2))
		gthgBaseReward = temp * rewardAnalytic.SharedRate * pool.GetRealValue()
		if rewardAnalytic.SharedRate <= 0 {
			bonusTHGEachUser = s.BonusGTHGEachUser
		}

	} else {
		gthgBaseReward = pool.GetRealValue() * rewardAnalytic.GTHGRateReward //0*x = 0
	}

	gthgBonusReward := gthgBaseReward * rewardAnalytic.RankBonusRate //0*x = 0

	{
		//hardcode + magic number//only isDaily
		if isDaily && s.Version == -1 {
			if rank >= 31 && rank <= 100 {
				gthgBaseReward = 5
			}
		}
	}

	reward.BaseReward = currencymodel.ConvertFloatToSystemCurrency(round(gthgBaseReward+bonusTHGEachUser, 0.01), pool.Type)
	reward.BonusReward = currencymodel.ConvertFloatToSystemCurrency(round(gthgBonusReward, 0.01), pool.Type)
	return reward, rewardAnalytic
}
func (d *DailyBattleConfig) GetDailyBattleConfigByMode(mode gamemodeenum.InGameMode) int {
	switch mode {
	case gamemodeenum.TEAM_COLLECT_STAR_4_VS_4:
		return d.SSDailyBattleConfig
	case gamemodeenum.DEATH_MATCH:
		return d.DMDailyBattleConfig
	case gamemodeenum.TOWER:
		return d.TWDailyBattleConfig
	case gamemodeenum.SOLO_SURVIVAL:
		return d.SoloDailyBattleConfig
	case gamemodeenum.DUAL_SURVIVAL:
		return d.DualDailyBattleConfig
	}
	return 0
}

func (s *TournamentEventData) ParseToClientData() *TournamentEventDataClient {
	client := &TournamentEventDataClient{}
	client.EventId = s.EventId
	client.EventType = eventtype.TOURNAMENT_EVENT_STRING
	client.StartIn = s.StartIn
	client.EndIn = s.EndIn()
	client.ProcessingTime = s.EndIn() + s.ProcessingTime
	client.DailyEventDataClient = s.DailyEventData.ParseToClientData()
	client.ListOfRequiredHeroes = s.ListOfRequiredHeroes
	client.ListOfRequiredHeroesDaily = s.ListOfRequiredHeroesDaily
	client.BannerMobileEventUrl = s.BannerMobileEventUrl
	client.BannerEventUrl = s.BannerEventUrl
	client.UnityBannerEventUrl = s.UnityBannerEventUrl
	client.UnityGameModeEventUrl = s.UnityGameModeEventUrl
	client.DailyBattleConfigClient = DailyBattleConfigClient{DailyBattleConfig: s.DailyBattleConfig}
	client.TrophyRequired = s.TrophyRequired

	client.PriceTicket = currencymodel.ConvertFloatToSystemCurrency(s.PriceTicketGTHC, currencyenum.GTHC)
	client.IsDailyEvent = s.IsDailyEvent
	return client
}

func (d *DailyEventData) ParseToClientData() *DailyEventDataClient {
	client := &DailyEventDataClient{}
	client.DailyEventId = d.DailyEventId
	client.DailyEventType = eventtype.DAILY_TOURNAMENT_EVENT_STRING
	client.DailyEventStartIn = d.DailyStartIn
	client.DailyEventEndIn = d.EndIn()
	return client
}

// util
func sortKeyIntMap(m interface{}) []int {
	iter := reflect.ValueOf(m).MapRange()
	var keys []int
	for iter.Next() {
		keys = append(keys, int(iter.Key().Int()))
	}
	sort.Ints(keys)
	return keys
}
func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
