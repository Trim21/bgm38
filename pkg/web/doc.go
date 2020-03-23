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
	j := docs.OpenAPI()
	app.Get("/swagger/doc.json", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(200).SendString(j)
	})

	app.Get("/swagger", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "text/html")
		ctx.Status(200).SendString(`
<!DOCTYPE html>
<html>
<head>
    <title>Bgm38 Api Server</title>
    <!-- needed for adaptive design -->
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta charset="utf-8" />
	<link rel="shortcut icon" href="https://blog.trim21.cn/favicon.ico">	
</head>
<body>
<redoc spec-url="/swagger/doc.json" hide-hostname="true" suppress-warnings="true" lazy-rendering></redoc>
<script src="https://rebilly.github.io/ReDoc/releases/v1.x.x/redoc.min.js"></script>
</body>
</html>
`)
	})
}
