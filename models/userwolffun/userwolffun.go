package userwolffun

import (
	crand "crypto/rand"
	"math/big"
	"time"

	"github.com/WolffunService/theta-shared-common/database/mongodb"
	"github.com/WolffunService/theta-shared-common/enums/gametypeenum"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	defaultExpiredCode = 600
)

type SourceInstallType int

const (
	SourceInstallNone SourceInstallType = iota
	GalaxyStore
)

type UserWolffun struct {
	mongodb.DefaultModel `bson:",inline"`
	IsRegisSuccess       bool                                `bson:"isRegisSuccess"`
	Code                 int                                 `bson:"code,omitempty"`
	Source               string                              `bson:"sourceCreate"`
	LastTimeCall         int64                               `bson:"lastTimeCall"`
	TimeoutCode          int64                               `bson:"timeoutCode"`
	NumInputWrong        int32                               `bson:"numInputWrong"` // số lần input code sai
	HeroesStrikeTheta    *GameData                           `bson:"heroesStrikeTheta,omitempty"`
	MapGameData          map[gametypeenum.TypeGame]*GameData `bson:"gameDatas,omitempty"`

	//account test
	IsTest bool `bson:"isTest,omitempty"`
}

func (u UserWolffun) GetMapGameData() map[gametypeenum.TypeGame]*GameData {
	if u.MapGameData == nil {
		u.MapGameData = make(map[gametypeenum.TypeGame]*GameData)
	}
	if u.HeroesStrikeTheta != nil { //old data
		if _, exist := u.MapGameData[gametypeenum.HeroesStrikeTheta]; !exist {
			u.MapGameData[gametypeenum.HeroesStrikeTheta] = u.HeroesStrikeTheta
		}
	}
	return u.MapGameData
}

func (user *UserWolffun) CheckCode(code int) bool {
	//always valid
	if user.IsTest {
		return true
	}
	return user.Code == code
}

func (user *UserWolffun) CodeIsActive() bool {
	if user.IsTest {
		return true
	}

	now := time.Now().UTC().Unix()
	return user.TimeoutCode > 0 && user.TimeoutCode > now
}

// random new code + expire time code
func (user *UserWolffun) RandomCode() {
	user.Code = randomCode()
	user.TimeoutCode = time.Now().UTC().Unix() + defaultExpiredCode
	user.NumInputWrong = 0
}

// làm mới lại code - và tạm disable tính năng này đến khi user resend code mới ( CodeIsActive == false)
func (user *UserWolffun) ResetCode() {
	user.Code = randomCode()
	user.TimeoutCode = 0
	user.NumInputWrong = 0
}

func (user *UserWolffun) Deleted() {
	user.IsRegisSuccess = false
	user.ResetCode()
}

func (user *UserWolffun) SetData(gameId gametypeenum.TypeGame, data *GameData) {
	if user.MapGameData == nil {
		user.MapGameData = make(map[gametypeenum.TypeGame]*GameData)
	}
	if data != nil {
		data.CreatedAt = time.Now().UTC()
	}
	user.MapGameData[gameId] = data
}

func (user *UserWolffun) GetData(gameId gametypeenum.TypeGame) *GameData {
	mapData := user.GetMapGameData()
	return mapData[gameId]
}

type GameData struct {
	UserId    primitive.ObjectID `bson:"userId,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`

	InstallSource []SourceInstallType `bson:"installSource"`

	GameCountCheckin int32 `bson:"gameCountCheckin"`
}

func NewUserWolffun(email string) *UserWolffun {
	user := &UserWolffun{}
	user.ID = email
	return user
}

func (user UserWolffun) CollectionName() string {
	return "UserWolffun"
}

func randomCode() int {
	code, _ := randomInt(100000, 1000000)
	return code
}

func randomInt(min, max int) (int, error) {
	i, err := random0ToInt(max - min)
	if err != nil {
		return max, nil
	}
	i += min
	return i, nil
}

// Random0ToInt return a number from 0 to max - 1, return 0 if max == 0 and return error if max's negative
func random0ToInt(max int) (int, error) {
	if max == 0 {
		return 0, nil
	}
	preRand, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return -1, err
	}
	return int(preRand.Int64()), nil
}
