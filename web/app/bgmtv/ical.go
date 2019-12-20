package bgmtv

import (
	"bgm38/web/app/bgmtv/viewip"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

// Part bind vote routers to gin app
func Part(app *iris.Application) {
	indexPart(app)
}

func indexPart(app *iris.Application) {
	var router = app.Party("/bgm.tv")
	viewip.Part(router)
	router.Get("/v1/calendar/{user_id:string}", hero.Handler(userCalendar))
}
