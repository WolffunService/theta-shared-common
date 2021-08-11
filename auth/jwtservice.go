package auth

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-common/common"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
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
		if claims[ClaimKeyPlayfabId] == nil || claims[ClaimKeyEmail] == nil {
			return nil, common.ErrorResponse(common.TokenInvalid,err.Error())
		}
		return claims, nil
	} else {
		log.Printf("Invalid JWT Token")
		return nil, common.ErrorResponse(common.TokenInvalid,err.Error())
	}
}


func (s service) parseUnverified(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, common.ErrorResponse(common.TokenInvalid,err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, common.ErrorResponse(common.TokenInvalid,err.Error())
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

func (s service) verifyToken(tokenString string)  (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.signingKey), nil
	})
	if err != nil {
		errorToken := err.(*jwt.ValidationError)
		return nil,common.ErrorResponse(common.TokenInvalid,err.Error()).RootCode(int(errorToken.Errors))
	}
	return token, nil
}