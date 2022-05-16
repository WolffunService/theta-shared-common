package middleware

import (
	"context"
	"fmt"
	"github.com/WolffunGame/theta-shared-common/auth"
	"github.com/WolffunGame/theta-shared-common/auth/entity"
	"github.com/WolffunGame/theta-shared-common/auth/rbac"
	"github.com/WolffunGame/theta-shared-common/common"
	"github.com/WolffunGame/theta-shared-common/common/thetaerror"
	"github.com/WolffunGame/theta-shared-database/database/mredis"
	"github.com/gin-gonic/gin"
	goredislib "github.com/go-redis/redis/v8"
	"net/http"
	"strings"
	"time"
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
			//c.Next()
			c.JSON(http.StatusForbidden, common.ErrorResponse(common.Error, "This account does not have this permission"))
			c.Abort()
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
		ctx := context.Background()

		rawAPIKey := c.Request.Header.Get("X-API-KEY")
		segments := strings.Split(rawAPIKey, ".")
		if len(segments) < 2 {
			c.JSON(http.StatusUnauthorized, &thetaerror.Error{
				Code:    thetaerror.ErrorInternal,
				Message: "API Key is not valid",
			}) //this err was common.ErrorResponse
			c.Abort()
			return
		}

		prefix := segments[0]
		hashKey := auth.HashRawKey(segments[1])

		isAllow, err := rbac.Enforce(prefix+"."+hashKey, object, action)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, common.ErrorResponse(common.Error,err.Error()))
			c.JSON(http.StatusUnauthorized, err) //this err was common.ErrorResponse
			c.Abort()
			return
		}

		if isAllow {
			apiKey, err := auth.GetAPIKey(ctx, prefix, hashKey)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err)
				c.Abort()
				return
			}

			isValidAccess, err := IsValidAccess(ctx, rawAPIKey, apiKey.AccessLimit, time.Now())
			if err != nil {
				c.JSON(http.StatusUnauthorized, err)
				c.Abort()
				return
			}

			if !isValidAccess {
				c.JSON(http.StatusForbidden, common.ErrorResponse(thetaerror.ErrorInternal, "This API Key has limited access"))
				c.Abort()
				return
			}

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

func IsValidAccess(ctx context.Context, rawAPIKey string, accessLimit map[entity.AccessLimitType]int64, dtnow time.Time) (bool, error) {

	if accessLimit == nil || len(accessLimit) <= 0 {
		return true, nil
	}
	client := mredis.GetClient()
	key := ""
	timeDuration := time.Second
	for limitType, limitCount := range accessLimit {
		switch limitType {
		case entity.AccessLimitTypeSecond:
			key = fmt.Sprintf("%v_%v", rawAPIKey, dtnow.Format("20060102150405"))
			timeDuration = time.Second
			break
		case entity.AccessLimitTypeMinute:
			key = fmt.Sprintf("%v_%v", rawAPIKey, dtnow.Format("200601021504"))
			timeDuration = time.Minute
			break
		case entity.AccessLimitTypeHour:
			key = fmt.Sprintf("%v_%v", rawAPIKey, dtnow.Format("2006010215"))
			timeDuration = time.Hour
			break
		case entity.AccessLimitTypeDay:
			key = fmt.Sprintf("%v_%v", rawAPIKey, dtnow.Format("20060102"))
			timeDuration = time.Hour * 24
			break
		default:
			return false, nil
			break
		}
		accessCount, errRedis := client.Get(ctx, key).Int64()
		if errRedis == goredislib.Nil {
			client.Set(ctx, key, 1, timeDuration)
		} else if errRedis != nil {
			return false, errRedis
		} else {
			if accessCount >= limitCount {
				return false, nil
			}
			client.Incr(ctx, key)
		}
	}
	return true, nil
}
