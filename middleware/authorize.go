package middleware

import (
	"github.com/WolffunGame/theta-shared-common/auth"
	"github.com/WolffunGame/theta-shared-common/auth/rbac"
	"github.com/WolffunGame/theta-shared-common/common"
	"github.com/WolffunGame/theta-shared-common/common/thetaerror"
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
			//c.JSON(http.StatusUnauthorized, common.ErrorResponse(common.Error,err.Error()))
			c.JSON(http.StatusUnauthorized, err) //this err was common.ErrorResponse
			c.Abort()
			return
		}
		if len(roles) == 0 {
			roles = append(roles, common.NONE)
		}
		c.Set(auth.ClaimKeyId, claims[auth.ClaimKeyId])
		c.Set(auth.ClaimKeySid, claims[auth.ClaimKeySid])

		userRole := common.NONE
		if claims[auth.ClaimKeyRole] != nil {
			userRole = common.UserRole(claims[auth.ClaimKeyRole].(float64))
		}
		c.Set(auth.ClaimKeyRole, userRole)

		if userRole == common.ROOT {
			c.Next()
			return
		}

		for _, role := range roles {
			if role == userRole {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, common.ErrorResponse(common.Error, "This account does not have this permission"))
		c.Abort()
	}
}

func RequiredAPIKeyVerified(apiKeyService auth.APIKeyService, rbac rbac.AuthorizationService, object string, action string) func(c *gin.Context) {
	return func(c *gin.Context) {
		//claims

		rawAPIKey := c.Request.Header.Get("X-API-KEY")
		isAllow, err := rbac.Enforce(rawAPIKey, object, action)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, common.ErrorResponse(common.Error,err.Error()))
			c.JSON(http.StatusUnauthorized, err) //this err was common.ErrorResponse
			c.Abort()
			return
		}

		if isAllow {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, common.ErrorResponse(thetaerror.ErrorInternal, "This API Key does not have permission to do this action"))
		c.Abort()
	}
}

func ParseValidToken(service auth.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := service.TokenValid(c.Request)
		if err != nil {
			c.Next()
			return
		}

		c.Set(auth.ClaimKeyId, claims[auth.ClaimKeyId])
		c.Set(auth.ClaimKeySid, claims[auth.ClaimKeySid])

		userRole := common.NONE
		if claims[auth.ClaimKeyRole] != nil {
			userRole = common.UserRole(claims[auth.ClaimKeyRole].(float64))
		}
		c.Set(auth.ClaimKeyRole, userRole)
		c.Next()
	}
}
