package middlewares

import (
	"io"
	"strings"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/context"
)

func SkipEmptyBody() iris.Handler {
	return func(ctx *context.Context) {
		// Check if the request body is empty
		if ctx.Request().ContentLength == 0 {
			// Replace empty request body with an empty JSON object
			ctx.Request().Body = io.NopCloser(strings.NewReader("{}"))
		}

		ctx.Next()
	}
}
