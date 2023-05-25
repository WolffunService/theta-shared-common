package thetancontext

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IUserData interface {
		//get set user data
		SetUserData(user UserData)
		GetUserID() string
		GetCountry() string
		GetBattle() int32

		//get set user ranking
		SetUserRanking(userRanking UserRanking)
		GetRankingLevel() int
		GetTrophy() int
		GetRankName() string
	}
)

type (
	// struct
	baseUserData struct {
		userID          string
		userData        UserData
		userRankingData UserRanking
	}

	UserData interface {
		GetUserID() string
		GetCountry() string
		GetBattle() int32
	}

	UserRanking interface {
		RankingLevel() int
		Trophy() int
		GetRankingFullName() string
	}
)

var _b IUserData = &baseUserData{}

func newBase(userID string) *baseUserData {
	return &baseUserData{
		userID: userID,
	}
}

func (b *baseUserData) GetUserID() string {
	return b.userID
}

func (b *baseUserData) SetUserData(user UserData) {
	b.userData = user
}
func (b *baseUserData) SetUserRanking(userRanking UserRanking) {
	b.userRankingData = userRanking
}

func (b *baseUserData) GetCountry() string {
	b.validUser()
	if b.userData == nil {
		return ""
	}
	return b.userData.GetCountry()
}

func (b *baseUserData) GetBattle() int32 {
	b.validUser()
	if b.userData == nil {
		return 0
	}
	return b.userData.GetBattle()
}

func (b *baseUserData) validUser() {
	defer recoverFn()
	if b.userData != nil {
		return
	}
	if b.userID == "" {
		panic(ErrUserNotFound{})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user, err := HookGetUser(ctx, b.userID)
	if err != nil || user == nil {
		panic(ErrUserNotFound{b.userID, err})
	}
	//getuserbyid
	b.userData = user
	return
}

func (b *baseUserData) GetRankingLevel() int {
	b.validUserRanking()
	if b.userRankingData == nil {
		return 0
	}
	return b.userRankingData.RankingLevel()
}

func (b *baseUserData) GetTrophy() int {
	b.validUserRanking()
	if b.userRankingData == nil {
		return 0
	}
	return b.userRankingData.Trophy()
}

func (b *baseUserData) GetRankName() string {
	b.validUserRanking()
	if b.userRankingData == nil {
		return ""
	}
	return b.userRankingData.GetRankingFullName()
}

func (b *baseUserData) validUserRanking() {
	defer recoverFn()

	if b.userRankingData != nil {
		return
	}
	if b.userID == "" {
		panic(ErrUserNotFound{})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userRanking, err := HookGetUserRanking(ctx, b.userID)
	if (err != nil && err != mongo.ErrNoDocuments) || userRanking == nil {
		panic(ErrUserNotFound{b.userID, err})
	}
	//getuserbyid
	b.userRankingData = userRanking
	return
}
