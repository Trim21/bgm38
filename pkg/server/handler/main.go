package handler

import "github.com/kataras/iris"

type A struct {
	Hello string `json:"hello" xml:"hello"`
}

func Index(ctx iris.Context) {
	ctx.Header("hello", "world")
	err := ctx.View("index.tmpl")
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}
