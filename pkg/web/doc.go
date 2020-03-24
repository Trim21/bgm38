// @title bgm38 api server
// @version 0.0.1
// @description  A Set of http api for bangumi.

// @schemes https
// @host api.bgm38.com
// @BasePath /

// @contact.name Trim21
// @contact.url https://github.com/Trim21/bgm38/issues
// @contact.email i@trim21.me

// @license.name MIT
// @license.url https://github.com/Trim21/bgm38/blob/dev/LICENSE

package web

import (
	"io/ioutil"

	"github.com/gofiber/fiber"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"

	"bgm38/pkg/web/docs"
)

func setupSwaggerRouter(app *fiber.App) {
	f, err := pkger.Open("/asserts/web/redoc.html")
	if err != nil {
		logrus.Fatalln("missing redoc html")
	}

	content, err := ioutil.ReadAll(f)

	if err != nil {
		logrus.Fatalln("can't read redoc html")
	}

	j := docs.OpenAPI()
	app.Get("/swagger/doc.json", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(200).SendString(j)
	})

	app.Get("/swagger", func(ctx *fiber.Ctx) {
		ctx.Set("Content-Type", "text/html")
		ctx.Status(200).SendBytes(content)
	})
}
