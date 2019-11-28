package server

import (
	"bgm38/pkg/server/vote"
	"bgm38/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Serve() error {
	app := newApp()
	vote.Part(app)
	return app.Run(":" + utils.GetEnv("PORT", "8080"))
}

func newApp() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("./bindata/templates/**/*.tmpl")
	//router.LoadHTMLGlob("./bindata/templates/*.tmpl")

	//	router.GET("/someGet", getting)
	//	router.POST("/somePost", posting)
	//	router.PUT("/somePut", putting)
	//	router.DELETE("/someDelete", deleting)
	//	router.PATCH("/somePatch", patching)
	//	router.HEAD("/someHead", head)
	//	router.OPTIONS("/someOptions", options)
	//
	//	By default it serves on :8080 unless a
	//	PORT environment variable was defined.
	//router.Run()
	//router.Run(":3000") for a hard coded port
	//
	return router
}

//
// loadTemplate loads templates embedded by go-assets-builder
//func loadTemplate() (*template.Template, error) {
//	t := template.New("")
//	for name := range bindata.AssetNames() {
//		file, _ := bindata.Asset(name)
//		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
//			continue
//		}
//		h, err := ioutil.ReadAll(file)
//		if err != nil {
//			return nil, err
//		}
//		t, err = t.New(name).Parse(string(h))
//		if err != nil {
//			return nil, err
//		}
//	}
//	return t, nil
//}
