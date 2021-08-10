package middleware

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-common/auth"
	"github.com/WolffunGame/theta-shared-common/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", common.TokenValid(nil,"",common.TokenInvalid)
	}

	return parts[1], nil
}
//token marketplace
func RequiredAuthUnverified() func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		//claims
		var service = auth.Default()
		claims, err := service.ParseUnverified(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		//
		fmt.Println(claims)
		c.Set(auth.ClaimKeyId, claims[auth.ClaimKeyId])
		c.Next()
	}
}

func RequiredAuthVerified(service auth.Service) func(c *gin.Context) {
	return func(c *gin.Context) {

		//claims
		claims, err := service.TokenValid(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		//
		fmt.Println(claims)
		c.Set(auth.ClaimKeyId, claims[auth.ClaimKeyId])
		c.Next()
	}
}