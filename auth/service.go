package auth

import (
	"net/http"

	"github.com/WolffunService/theta-shared-common/auth/entity"
	"github.com/WolffunService/theta-shared-common/log"
	"github.com/golang-jwt/jwt"
)

type Service interface {
	RefreshToken(identity entity.Identity) (*entity.TokenResBody, error)
	TokenValid(r *http.Request) (jwt.MapClaims, error)
	ParseUnverified(tokenString string) (jwt.MapClaims, error)
	StringTokenValid(refreshToken string) (jwt.MapClaims, error)
}

type service struct {
	signingKey             string
	tokenExpiration        int
	refreshTokenExpiration int
	audience               string
	issuer                 string
	logger                 log.Logger
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration, refreshTokenExpiration int, audience, issuer string, logger log.Logger) Service {
	return service{signingKey, tokenExpiration,
		refreshTokenExpiration, audience, issuer, logger}
}
func Default() Service {
	return service{}
}

////generate new token
//func (s service) RefreshToken(identity entity.Identity) (*entity.TokenResBody, *errors.Response)  {
//	return s.generateJWT(iden tity)
//}

// TokenValid check token valid - return claims
func (s service) TokenValid(r *http.Request) (jwt.MapClaims, error) {
	return s.tokenValid(r)
}

// StringTokenValid check token valid - return jwt token
func (s service) StringTokenValid(refreshToken string) (jwt.MapClaims, error) {
	return s.tokenValidString(refreshToken)
}
func (s service) ParseUnverified(tokenString string) (jwt.MapClaims, error) {
	return s.parseUnverified(tokenString)
}

// RefreshToken generate new token
func (s service) RefreshToken(identity entity.Identity) (*entity.TokenResBody, error) {
	return s.generateJWT(identity)
}
