package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/WolffunGame/theta-shared-common/auth/entity"
	"github.com/WolffunGame/theta-shared-common/common"
	"github.com/dgrijalva/jwt-go"
)

//check token http author
func (s service) tokenValid(r *http.Request) (jwt.MapClaims, error) {
	//get token from Authorization - Header
	tokenString := extractToken(r)
	//check valid
	return s.tokenValidString(tokenString)
}

//check token string
func (s service) tokenValidString(tokenString string) (jwt.MapClaims, error) {
	token, err := s.verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims[ClaimKeyId] == nil {
			return nil, common.ErrorResponse(common.TokenInvalid, err.Error())
		}
		return claims, nil
	} else {
		log.Printf("Invalid JWT Token")
		return nil, common.ErrorResponse(common.TokenInvalid, err.Error())
	}
}

func (s service) parseUnverified(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, common.ErrorResponse(common.TokenInvalid, err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, common.ErrorResponse(common.TokenInvalid, err.Error())
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (s service) verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.signingKey), nil
	})
	if err != nil {
		errorToken := err.(*jwt.ValidationError)
		if errorToken.Errors == jwt.ValidationErrorExpired {
			return nil, common.ErrorResponse(common.TokenExpired, err.Error())
		}
		return nil, common.ErrorResponse(common.TokenInvalid, err.Error()).RootCode(int(errorToken.Errors))
	}
	return token, nil
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity entity.Identity) (*entity.TokenResBody, error) {
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		ClaimKeySid:     identity.GetAddress(),
		ClaimKeySub:     identity.GetUserName(),
		ClaimKeyId:      identity.GetUserId(),
		ClaimKeyRole:    identity.GetRole(),
		ClaimKeyCanMint: false,
		ClaimKeyNbf:     time.Now().Unix(),
		ClaimKeyIss:     s.issuer,
		ClaimKeyAud:     s.audience,
		ClaimKeyExp:     time.Now().Add(time.Second * time.Duration(s.tokenExpiration)).Unix(),
	}).SignedString([]byte(s.signingKey))
	if err != nil {
		return nil, err
	}

	rt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		ClaimKeySid:     identity.GetAddress(),
		ClaimKeySub:     identity.GetUserName(),
		ClaimKeyId:      identity.GetUserId(),
		ClaimKeyCanMint: false,
		ClaimKeyNbf:     time.Now().Unix(),
		ClaimKeyIss:     s.issuer,
		ClaimKeyAud:     s.audience,
		ClaimKeyExp:     time.Now().Add(time.Second * time.Duration(s.refreshTokenExpiration)).Unix(),
	}).SignedString([]byte(s.signingKey))
	if err != nil {
		return nil, err
	}

	return &entity.TokenResBody{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}
