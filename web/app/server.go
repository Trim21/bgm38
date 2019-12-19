package app

import (
	"bgm38/config"
	"bgm38/pkg/utils"
	"bgm38/web/app/auth"
	"bgm38/web/app/bgmtv"
	"bgm38/web/app/bindata"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"strings"
)

//Serve start http web on env `PORT` or 8080
func Serve() error {
	app := newApp()
	if gin.IsDebugging() {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return app.Run(":" + utils.GetEnv("PORT", "8080"))
}

func newApp() *gin.Engine {
	app := gin.Default()
	app.Use(versionMiddleware)
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	app.SetHTMLTemplate(t)
	auth.Part(app)
	bgmtv.Part(app)
	return app
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		file, err := bindata.Asset(name)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(file))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func versionMiddleware(c *gin.Context) {
	c.Header("x-web-version", config.Version)
	c.Next()
}
