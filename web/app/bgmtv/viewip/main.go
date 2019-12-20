package viewip

import (
	"github.com/gin-gonic/gin"
)

// Part bind vote routers to gin app
func Part(app *gin.Engine) {
	indexPart(app)
}

func indexPart(app *gin.Engine) {
	var router = app.Group("/bgm.tv")
	router.GET("/v0/meta/subject/:subject_id", getSubjectFullInfo)
	router.GET("/v2/view_ip/subject/:subject_id", viewIP)

}
