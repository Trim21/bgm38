package middleware

import (
	"github.com/kataras/iris"
)

func Before(ctx iris.Context) {
	ctx.Next()
}
