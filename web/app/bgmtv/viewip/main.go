package viewip

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

// Part bind vote routers to gin app
func Part(app iris.Party) {
	var router = app.Party("/view_ip")
	router.Get("/meta/v0/subject/{subject_id:int}", hero.Handler(getSubjectFullInfo))
	router.Get("/v2/subject/{subject_id:int}", hero.Handler(viewIP))
}
