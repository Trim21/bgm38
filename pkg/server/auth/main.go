package auth

import (
	"github.com/gin-gonic/gin"
)

func Part(app *gin.Engine) {
	var router = app.Group("/auth")
	router.GET("/v1/callback", callback)
	router.GET("/v1/redirect", redirect)
}
