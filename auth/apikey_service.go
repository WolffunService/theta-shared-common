package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/WolffunService/theta-shared-common/auth/entity"
	"github.com/WolffunService/theta-shared-common/auth/rbac"
	"github.com/WolffunService/theta-shared-common/common/thetaerror"
	"github.com/WolffunService/theta-shared-common/thetalog"
	"github.com/WolffunService/theta-shared-database/common/util"
	"github.com/WolffunService/theta-shared-database/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type APIKeyService interface {
	Generate(ctx context.Context, owner string, role string, listAccessLimit []entity.AccessLimitInfo) (*entity.APIKeyResult, error)
	Parse(ctx context.Context, r *http.Request) (*entity.APIKey, error)
	Revoke(key string) error
	ChangeAccessLimit(ctx context.Context, rawAPIKey string, accessLimit []entity.AccessLimitInfo) (*entity.ChangeAccessLimitResult, error)
}

func NewAPIKeyService() APIKeyService {
	err := rbac.InitService(mongodb.GetDB(), "RBACRules")
	if err != nil {
		thetalog.Err(err).Msg("Init API Key service error")
	}
	rbac.Service.LoadPolicy()

	return apiKeyService
}

type apiKeyServiceImplement struct {
}

var apiKeyService apiKeyServiceImplement

func (a apiKeyServiceImplement) Generate(ctx context.Context, owner string, role string, accessLimit []entity.AccessLimitInfo) (*entity.APIKeyResult, error) {
	prefix := randStringBytesMaskImprSrc(7)
	apiKey := randStringBytesMaskImprSrc(64)
	hashKey := HashRawKey(apiKey)

	var mapAccessLimit = make(map[entity.AccessLimitType]int64)
	for _, limit := range accessLimit {
		mapAccessLimit[limit.LimitType] = limit.LimitCount
	}

	key := entity.APIKey{
		Prefix:      prefix,
		HashKey:     hashKey,
		Owner:       owner,
		Status:      entity.APIKeyStatusEnabled,
		AccessLimit: mapAccessLimit,
	}
	key.CreatedAt = time.Now().UTC()

	err := createAPIKey(ctx, &key)

	if err != nil {
		return nil, err
	}

	_, err = rbac.Service.GetEnforce().AddRoleForUser(key.Prefix+"."+key.HashKey, role)
	if err != nil {
		return nil, err
	}

	return &entity.APIKeyResult{
		RawKey: key.Prefix + "." + apiKey,
		Role:   role,
		Owner:  owner,
	}, nil
}

func (a apiKeyServiceImplement) ChangeAccessLimit(ctx context.Context, rawAPIKey string, accessLimit []entity.AccessLimitInfo) (*entity.ChangeAccessLimitResult, error) {

	segments := strings.Split(rawAPIKey, ".")

	if len(segments) < 2 {
		return nil, &thetaerror.Error{
			Code:    thetaerror.ErrorInternal,
			Message: "API Key is not valid",
		}
	}

	prefix := segments[0]
	hashKey := HashRawKey(segments[1])

	var mapAccessLimit = make(map[entity.AccessLimitType]int64)
	for _, limit := range accessLimit {
		mapAccessLimit[limit.LimitType] = limit.LimitCount
	}

	updateCount, err := updateAPIKey(ctx, prefix, hashKey, mapAccessLimit)

	if err != nil {
		return nil, err
	}

	return &entity.ChangeAccessLimitResult{
		UpdatedCount: updateCount,
	}, nil
}

func (a apiKeyServiceImplement) Parse(ctx context.Context, r *http.Request) (*entity.APIKey, error) {
	rawAPIKey := r.Header.Get("X-API-KEY")
	segments := strings.Split(rawAPIKey, ".")

	if len(segments) < 2 {
		return nil, &thetaerror.Error{
			Code:    thetaerror.ErrorInternal,
			Message: "API Key is not valid",
		}
	}

	prefix := segments[0]
	hashKey := HashRawKey(segments[1])

	apiKey, err := GetAPIKey(ctx, prefix, hashKey)

	if err != nil {
		return nil, err
	}

	return apiKey, nil
}

func (a apiKeyServiceImplement) Revoke(rawKey string) error {
	_, err := rbac.Service.GetEnforce().DeleteRolesForUser(rawKey)
	return err
}

func HashRawKey(raw string) string {
	hash := md5.Sum([]byte(raw))
	return hex.EncodeToString(hash[:])
}

func checkKeyHash(rawKey, hashKey string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashKey), []byte(rawKey))
	return err == nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// Hàm random này performance cao mà nhìn ngầu
func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func createAPIKey(ctx context.Context, data *entity.APIKey) error {
	collection := mongodb.Coll(data)
	if err := collection.CreateWithCtx(ctx, data); err != nil {
		return err
	}

	return nil
}

func updateAPIKey(ctx context.Context, prefix string, hashKey string, mapAccessLimit map[entity.AccessLimitType]int64) (int64, error) {
	collection := mongodb.Coll(&entity.APIKey{})
	filter := bson.D{
		{Key: "prefix", Value: prefix},
		{Key: "hashKey", Value: hashKey},
	}
	update := bson.D{}
	update = util.BsonSet(update, "accessLimit", mapAccessLimit)

	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, nil
}

func GetAPIKey(ctx context.Context, prefix string, hashKey string) (*entity.APIKey, error) {
	filter := bson.D{
		{Key: "prefix", Value: prefix},
		{Key: "hashKey", Value: hashKey},
	}

	data := &entity.APIKey{}
	coll := mongodb.CollRead(data)

	err := coll.FirstWithCtx(ctx, filter, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
