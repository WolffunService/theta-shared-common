package middlewares

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"
)

func CORS() context.Handler {
	return cors.AllowAll()
}
