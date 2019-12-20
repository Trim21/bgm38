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
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func setupSwagger(app *iris.Application) {
	logrus.Debugln(docs.AssetNames())
	logrus.Debugln("setup swagger")
	docs.SwaggerInfo.Version = config.Version

	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(&swagger.Config{
		URL: "/swagger/doc.json", // The url pointing to API definition
	}, swaggerFiles.Handler))

	app.Get("/swagger", func(context iris.Context) {
		context.StatusCode(200)
		context.HTML(`<!DOCTYPE html>
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
