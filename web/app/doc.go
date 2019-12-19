// Package app

// @title bgm38 api server
// @version 0.0.1
// @description  A Set of http api for bangumi.
// @termsOfService No

// @contact.name Trim21
// @contact.url https://github.com/Trim21
// @contact.email i@trim21.me

// @license.name MIT
// @license.url https://github.com/Trim21/bgm38/blob/dev/LICENSE
// @schemes https
// @host api.bgm38.com
// @BasePath /

package app

//go:generate swag init --generalInfo ./doc.go -o ./docs
//go:generate go-bindata -o ./docs/bindata.go -fs -prefix "./docs/" -pkg docs ./docs/swagger.json ./docs/swagger.yaml
import (
	"bgm38/config"
	"bgm38/web/app/docs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger(r *gin.Engine) {
	logrus.Debugln(docs.AssetNames())
	logrus.Debugln("setup swagger")
	if gin.IsDebugging() {
		docs.SwaggerInfo.Schemes = []string{"http"}
		docs.SwaggerInfo.Host = "localhost:8080"
	}
	docs.SwaggerInfo.Version = config.Version
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/swagger", func(context *gin.Context) {
		context.Header("content-type", "text/html")
		context.String(200, `<!DOCTYPE html>
<html>
<head>
<link type="text/css" rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui.css">
<link rel="shortcut icon" href="https://fastapi.tiangolo.com/img/favicon.png">
<title>Pol server - Swagger UI</title>
</head>
<body>
<div id="swagger-ui">
</div>
<script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
<script>
const ui = SwaggerUIBundle({
	url: '/swagger/doc.json',

	dom_id: '#swagger-ui',
	presets: [
	SwaggerUIBundle.presets.apis,
	SwaggerUIBundle.SwaggerUIStandalonePreset
	],
	layout: "BaseLayout",
	deepLinking: true
})
</script>
</body>
</html>
`)
	})
}
