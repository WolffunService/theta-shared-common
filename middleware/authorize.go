package middleware

import (
	"github.com/WolffunGame/theta-shared-common/auth"
	"github.com/WolffunGame/theta-shared-common/common"
	"github.com/WolffunGame/theta-shared-database/user/usermodel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", common.ErrorResponse(common.TokenInvalid, "")
	}
	return parts[1], nil
}

func RequiredAuthVerified(service auth.Service, roles ...usermodel.UserRole) func(c *gin.Context) {
	return func(c *gin.Context) {

		//claims
		claims, err := service.TokenValid(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		if len(roles) == 0 {
			roles = append(roles, usermodel.NONE)
		}
		//
		//fmt.Println(claims)
		c.Set(auth.ClaimKeyId, claims[auth.ClaimKeyId])
		c.Set(auth.ClaimKeySid, claims[auth.ClaimKeySid])
		userRole := claims[auth.ClaimKeyRole].(usermodel.UserRole)
		for _, role := range roles {
			if role == userRole{
				c.Next()
				break
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, "This account does not have this permission")
	}
}