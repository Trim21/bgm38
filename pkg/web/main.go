package web

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/gofiber/requestid"

	"bgm38/pkg/web/bgmtv"
)

func Start() error {
	return CreateApp().Listen(3000)
}

func CreateApp() *fiber.App {
	app := fiber.New()
	app.Use(requestid.New())

	app.Use(recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			fmt.Println(err)
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}))

	setupSwagger(app)
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
	bgmtv.Group(app)
	rootRouter(app)
	app.Use(func(c *fiber.Ctx) {
		c.Status(404).SendString(`{}`)
		// => 404 "Not Found"
	})

	return app
}
