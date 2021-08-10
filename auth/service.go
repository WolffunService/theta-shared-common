package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

type Service interface {
	TokenValid(r *http.Request) (jwt.MapClaims,error)
	ParseUnverified(tokenString string) (jwt.MapClaims, error)
	StringTokenValid(refreshToken string) (jwt.MapClaims, error)
}
type service struct {
	signingKey      string
	tokenExpiration int
	refreshTokenExpiration int
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration int, refreshTokenExpiration int, logger log.Logger) Service {
	return service{signingKey, tokenExpiration, refreshTokenExpiration, logger}
}
func Default() Service {
	return service{}
}
////generate new token
//func (s service) RefreshToken(identity entity.Identity) (*entity.TokenResBody, *errors.Response)  {
//	return s.generateJWT(identity)
//}

//check token valid - return claims
func (s service) TokenValid(r *http.Request) (jwt.MapClaims, error){
	return s.tokenValid(r)
}
//check token valid - return jwt token
func (s service) StringTokenValid(refreshToken string) (jwt.MapClaims, error) {
	return s.tokenValidString(refreshToken)
}
func (s service) ParseUnverified(tokenString string) (jwt.MapClaims, error) {
	return s.parseUnverified(tokenString)
}