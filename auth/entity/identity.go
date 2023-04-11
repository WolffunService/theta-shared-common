package entity

import "github.com/WolffunService/theta-shared-database/database/mongodb"

// Identity represents an authenticated user identity.
type Identity interface {
	// GetUserId returns the user ID.
	GetUserId() string

	GetUserName() string

	GetAddress() string

	GetRole() int
}

type TokenResBody struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type APIKeyResult struct {
	RawKey string `json:"rawKey"`
	Role   string `json:"role"`
	Owner  string `json:"owner"`
}

type APIKey struct {
	mongodb.DefaultModel `json:",inline" bson:",inline"`
	mongodb.DateFields   `json:",inline" bson:",inline"`
	Prefix               string                    `json:"prefix" bson:"prefix"`
	HashKey              string                    `json:"hashKey" bson:"hashKey"`
	Owner                string                    `json:"owner" bson:"owner"`
	Status               APIKeyStatus              `json:"status" bson:"status"`
	AccessLimit          map[AccessLimitType]int64 `json:"accessLimit" bson:"accessLimit"`
}

type APIKeyStatus int

const (
	APIKeyStatusEnabled  = 1
	APIKeyStatusDisabled = 0
)

func (c APIKey) CollectionName() string {
	return "APIKeys"
}

type AccessLimitType int

const (
	AccessLimitTypeSecond AccessLimitType = 1
	AccessLimitTypeMinute AccessLimitType = 2
	AccessLimitTypeHour   AccessLimitType = 3
	AccessLimitTypeDay    AccessLimitType = 4
)

type AccessLimitInfo struct {
	LimitType  AccessLimitType `json:"limitType" bson:"limitType"`
	LimitCount int64           `json:"limitCount" bson:"limitCount"`
}

type ChangeAccessLimitResult struct {
	UpdatedCount int64 `json:"updatedCount"`
}
