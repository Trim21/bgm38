package web

import (
	"io"
	"path"

	"github.com/gofiber/fiber"
	"github.com/gofiber/requestid"
	"github.com/markbates/pkger"
	"go.uber.org/zap"

	"bgm38/pkg/utils"
	"bgm38/pkg/utils/log"
	"bgm38/pkg/web/bgmtv"
	"bgm38/pkg/web/middleware/headerversion"
	"bgm38/pkg/web/middleware/sentry"
	"bgm38/pkg/web/utils/handler"
)

func Start() error {
	var port = utils.GetEnv("PORT", "3000")
	log.GetLogger().Info("start listen on http://127.0.0.1:" + port)
	return CreateApp().Listen(port)
}

func CreateApp() *fiber.App {

	app := fiber.New()
	app.Settings.StrictRouting = true
	setupMiddleware(app)
	setupSwaggerRouter(app)
	app.Get("/", func(c *fiber.Ctx) {
		c.Redirect("https://api.bgm38.com/swagger")
	})
	app.Get("/asserts/web/*", handler.LogError(func(c *fiber.Ctx, logger *zap.Logger) error {
		filepath := c.Params("*")
		f, err := pkger.Open(path.Join("/asserts/web/", filepath))
		if err != nil {
			c.SendStatus(404)
			return nil
		}
		defer f.Close()
		c.Type(path.Ext(filepath))
		_, err = io.Copy(c.Fasthttp.Response.BodyWriter(), f)
		return err
	}))

	app.Get("/test", handler.LogError(func(c *fiber.Ctx, logger *zap.Logger) error {
		logger.Info("hello", zap.Int("key", 8))
		c.Send("Hello, World!")
		return nil
	}))

	bgmtv.Group(app)
	rootRouter(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) {
		c.Set("content-type", "application/json")
		c.Status(404).
			SendString(`{"message": "not found", "statue": "error"}`)
	})

	return app
}

func setupMiddleware(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(headerversion.New())
	app.Use(sentry.New())
}
