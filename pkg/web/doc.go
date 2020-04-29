// @title bgm38 api server
// @version dev
// @description  A Set of http api for bangumi.

// @schemes https
// @host api.bgm38.com
// @BasePath /

// @contact.name github.com/Trim21/bgm38
// @contact.url https://github.com/Trim21/bgm38

// @license.name MIT License
// @license.url https://github.com/Trim21/bgm38/blob/master/LICENSE

package web

import (
	"io/ioutil"

	"github.com/gofiber/fiber"
	"github.com/markbates/pkger"

	"bgm38/pkg/web/docs"
	"bgm38/pkg/web/utils/logger"
)

func setupSwaggerRouter(app *fiber.App) {
	f, err := pkger.Open("/asserts/web/doc.html")
	if err != nil {
		logger.GetLogger().Fatal("missing doc html")
	}

	content, err := ioutil.ReadAll(f)

	if err != nil {
		logger.GetLogger().Fatal("can't read doc html")
	}

	j := docs.OpenAPI()
	app.Get("/swagger.json", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(200).SendString(j)
	})

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "text/html")
		ctx.Status(200).SendBytes(content)
	})
}
