package middlewares

import (
	"errors"
	"fmt"
	"github.com/WolffunService/thetan-shared-common/comm"
	"github.com/WolffunService/thetan-shared-common/pkg/jwtauth"
	"github.com/WolffunService/thetan-shared-common/thetanerr"

	"github.com/WolffunService/thetan-shared-common/common"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
)

const AuthHandlerName = "iris.jwt"

var verifier *jwt.Verifier

func InitVerifier(signingKey string, blocklist jwt.Blocklist) {
	verifier = jwt.NewVerifier(jwt.HS256, []byte(signingKey))
	verifier.Blocklist = blocklist
	verifier.ErrorHandler = func(ctx *context.Context, err error) {
		resp := comm.CodeResponse(thetanerr.ErrTokenInvalid)
		if errors.Is(err, jwt.ErrExpired) {
			resp = comm.CodeResponse(thetanerr.ErrTokenExpired)
		}

		resp.Debug = fmt.Sprintf("err: %v", err)
		_ = ctx.StopWithJSON(iris.StatusUnauthorized, resp)
	}
}

func Authenticate() context.Handler {
	return verifier.Verify(func() any {
		return new(jwtauth.ThetanUserClaims)
	})
}

func RequiredRole(role common.UserRole, isProduction bool) context.Handler {
	return func(ctx *context.Context) {
		claims, ok := jwt.Get(ctx).(*jwtauth.ThetanUserClaims)
		if !ok || isProduction {
			dontHavePermission(ctx)
			return
		}

		if claims.Role >= role {
			ctx.Next()
			return
		}

		dontHavePermission(ctx)
	}
}

var dontHavePermission = func(ctx *context.Context) {
	resp := comm.ErrorResponse(fmt.Errorf("you don't have permission"))
	_ = ctx.StopWithJSON(iris.StatusForbidden, resp)
}
