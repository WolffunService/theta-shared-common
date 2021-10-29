package middleware

import (
	"github.com/WolffunGame/theta-shared-common/auth"
	"github.com/WolffunGame/theta-shared-common/common"
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

func RequiredAuthVerified(service auth.Service, roles ...common.UserRole) func(c *gin.Context) {
	return func(c *gin.Context) {
		//claims
		claims, err := service.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.ErrorResponse(common.Error,err.Error()))
			c.Abort()
			return
		}
		if len(roles) == 0 {
			roles = append(roles, common.NONE)
		}
		c.Set(auth.ClaimKeyId, claims[auth.ClaimKeyId])
		c.Set(auth.ClaimKeySid, claims[auth.ClaimKeySid])
		c.Set(auth.ClaimKeyRole, claims[auth.ClaimKeyRole])
		userRole := claims[auth.ClaimKeyRole].(float64)
		for _, role := range roles {
			if role == common.UserRole(userRole){
				c.Next()
				break
			}
		}
		c.JSON(http.StatusForbidden, common.ErrorResponse(common.Error, "This account does not have this permission"))
		c.Abort()
	}
}