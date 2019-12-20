package auth

import (
	"github.com/kataras/iris/v12"
)

// Part bind auth part to gin app
func Part(app iris.Application) {
	var router = app.Party("/auth")
	router.Get("/v1/bgm.tv/callback", callback)
	router.Get("/v1/bgm.tv/redirect", redirect)

}
