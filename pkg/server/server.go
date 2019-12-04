package server

import (
	"bgm38/pkg/server/auth"
	"bgm38/pkg/server/bindata"
	"bgm38/pkg/server/vote"
	"bgm38/pkg/utils"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

//Serve start http server on env `PORT` or 8080
func Serve() error {
	app := newApp()
	return app.Run(":" + utils.GetEnv("PORT", "8080"))
}

func newApp() *gin.Engine {
	app := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	app.SetHTMLTemplate(t)
	vote.Part(app)
	auth.Part(app)
	return app
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		file, err := bindata.Asset(name)
		t, err = t.New(name).Parse(string(file))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
