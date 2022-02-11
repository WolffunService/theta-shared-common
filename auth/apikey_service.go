package auth

import (
	"github.com/WolffunGame/theta-shared-common/auth/entity"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/http"
	"time"
)

type APIKeyService interface {
	Generate() (*entity.TokenResBody, error)
	Parse(r *http.Request) (jwt.MapClaims, error)
	Revoke(key string) (jwt.MapClaims, error)
}

func NewAPIKeyService() APIKeyService {
	return apiKeyService
}

type apiKeyServiceImplement struct {
}

var apiKeyService apiKeyServiceImplement

func (a apiKeyServiceImplement) Generate() (*entity.TokenResBody, error) {
	prefix := randStringBytesMaskImprSrc(7)
	apiKey := randStringBytesMaskImprSrc(64)

}

func (a apiKeyServiceImplement) Parse(r *http.Request) (jwt.MapClaims, error) {
	//TODO implement me
	panic("implement me")
}

func (a apiKeyServiceImplement) Revoke(key string) (jwt.MapClaims, error) {
	//TODO implement me
	panic("implement me")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

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
