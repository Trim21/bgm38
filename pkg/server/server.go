package server

import (
	"bgm38/pkg/server/auth"
	"bgm38/pkg/server/vote"
	"bgm38/pkg/utils"
	"github.com/gin-gonic/gin"
)

//Serve start http server on env `PORT` or 8080
func Serve() error {
	app := newApp()
	return app.Run(":" + utils.GetEnv("PORT", "8080"))
}

func newApp() *gin.Engine {
	app := gin.Default()
	app.LoadHTMLGlob("./bindata/templates/**/*.tmpl")
	vote.Part(app)
	auth.Part(app)
	return app
}
