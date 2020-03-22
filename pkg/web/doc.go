// @title bgm38 api server
// @version 0.0.1
// @description  A Set of http api for bangumi.

// @schemes https
// @host api.bgm38.com
// @BasePath /

// @contact.name Trim21
// @contact.url https://github.com/Trim21
// @contact.email i@trim21.me

// @license.name MIT
// @license.url https://github.com/Trim21/bgm38/blob/dev/LICENSE

package web

import (
	"github.com/gofiber/fiber"

	"bgm38/pkg/web/docs"
)

func setupSwagger(app *fiber.App) {
	j := docs.OpenApi()
	app.Get("/swagger/doc.json", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(200).SendString(j)
	})

	app.Get("/swagger", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "text/html")
		ctx.Status(200).SendString(`<!DOCTYPE html>	
<html>	
<head>	
<link type="text/css" rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui.css">	
<link rel="shortcut icon" href="https://blog.trim21.cn/favicon.ico">	
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
