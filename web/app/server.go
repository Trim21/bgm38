package app

import (
	"bgm38/pkg/utils"
	"bgm38/web/app/bgmtv"
	"bgm38/web/app/db"
	"bgm38/web/app/res"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func Serve() error {
	db.InitDB()
	app := iris.Default()
	app.Logger().SetLevel("debug")
	bgmtv.Part(app)
	setupSwagger(app)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	return app.Run(iris.Addr(":" + utils.GetEnv("PORT", "8080")))

}

func internalServerError(ctx iris.Context) {
	_, err := ctx.JSON(res.Error{
		Message: "internal error",
		Status:  "error",
	})
	if err != nil {
		logrus.Errorln(err)
	}
}
