package web

import (
	"github.com/gofiber/embed"
	"github.com/gofiber/fiber"
	"github.com/gofiber/requestid"
	"github.com/markbates/pkger"
	"go.uber.org/zap"

	"bgm38/app/web/bgmtv"
	"bgm38/app/web/middleware/headerversion"
	"bgm38/app/web/middleware/sentry"
	"bgm38/app/web/utils/handler"
	"bgm38/app/web/utils/logger"
	"bgm38/pkg/utils"
)

func Start() error {
	var port = utils.GetEnv("PORT", "3000")
	logger.GetLogger().Info("start listen on http://127.0.0.1:" + port)
	return createApp().Listen(port)
}

func createApp() *fiber.App {
	app := fiber.New()
	app.Settings.StrictRouting = true
	setupMiddleware(app)
	setupSwaggerRouter(app)
	assertsRouter(app)

	app.Get("/ping", handler.LogError(func(c *fiber.Ctx, logger *zap.Logger) error {
		logger.Info("hello", zap.Int("key", 8))
		c.Send("Hello, World!")
		return nil
	}))

	bgmtv.Group(app)
	md2bbcRouter(app)

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

func assertsRouter(app *fiber.App) {
	app.Use("/asserts/web", embed.New(embed.Config{
		Root: pkger.Dir("/asserts/web"),
	}))
}
