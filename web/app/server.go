package app

//go:generate go-bindata -o ./bindata/templates.go -fs -prefix "../templates" -pkg bindata ../templates/...

import (
	"html/template"
	"strings"
	"time"

	"bgm38/config"
	"bgm38/pkg/utils"
	"bgm38/web/app/auth"
	"bgm38/web/app/bgmtv"
	"bgm38/web/app/bgmtv/viewip"
	"bgm38/web/app/bindata"
	"bgm38/web/app/res"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Serve start http web on env `PORT` or 8080
func Serve() error {
	app := newApp()
	if gin.IsDebugging() {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return app.Run(":" + utils.GetEnv("PORT", "8080"))
}

func newApp() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger(), nice.Recovery(recoveryHandler))
	app.Use(versionMiddleware)
	err := setupMiddleware(app)
	if err != nil {
		logrus.Fatalln(err)
	}
	t, err := loadTemplate()
	if err != nil {
		logrus.Fatalln(err)
	}
	app.SetHTMLTemplate(t)
	setupSwagger(app)
	auth.Part(app)
	bgmtv.Part(app)
	viewip.Part(app)
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

func setupMiddleware(app *gin.Engine) error {
	if gin.IsDebugging() {
		cors.DefaultConfig()
		app.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "PUT", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	return nil
}

func recoveryHandler(c *gin.Context, err interface{}) {
	logrus.Errorln(err)
	c.JSON(500, res.Error{
		Message: "A internal error happened",
		Status:  "error",
	})
}
