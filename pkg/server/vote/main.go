package vote

import (
	"github.com/gin-gonic/gin"
)

func Part(app *gin.Engine) {
	indexPart(app)
}

func indexPart(app *gin.Engine) {
	var router = app.Group("/vote")
	router.GET("/", index)
	router.GET("/svg/:id", svg)
}
