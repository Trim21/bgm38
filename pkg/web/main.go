package web

import (
	"io"
	"path"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/markbates/pkger"
	"go.uber.org/zap"

	"bgm38/config"
	"bgm38/pkg/utils"
	"bgm38/pkg/utils/log"
	"bgm38/pkg/web/bgmtv"
	"bgm38/pkg/web/middleware/requestid"
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
	app.Get("/asserts/web/*", func(c *fiber.Ctx) {
		filepath := c.Params("*")
		f, err := pkger.Open(path.Join("/asserts/web/", filepath))
		if err != nil {
			c.SendStatus(404)
			return
		}
		defer f.Close()
		c.Type(path.Ext(filepath))
		_, err = io.Copy(c.Fasthttp.Response.BodyWriter(), f)
		if err != nil {
			log.GetLogger().Error(err.Error())
		}
	})

	app.Get("/test", handler.LogError(func(c *fiber.Ctx, logger *zap.Logger) error {
		logger.Info("hello", zap.Int("key", 8))
		c.Send("Hello, World!")
		return nil
	}))

	bgmtv.Group(app)
	rootRouter(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) {
		c.Status(404).
			SendString(`{"message": "not found", "statue": "error"}`)
	})

	return app
}

func setupMiddleware(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(func(c *fiber.Ctx) {
		c.Set("x-server-version", config.Version)
		c.Next()
	})

	var logHandler func(*fiber.Ctx, error)

	if sentry.CurrentHub().Client() == nil {
		logHandler = func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		}
	} else {
		logHandler = func(c *fiber.Ctx, err error) {
			event := sentry.Event{
				Contexts: make(map[string]interface{}),
				Extra:    make(map[string]interface{}),
				Tags:     make(map[string]string, 20),
				Modules:  make(map[string]string),
				Release:  config.Version,
			}

			c.Fasthttp.Request.Header.VisitAll(func(key, value []byte) {
				event.Contexts[string(key)] = value
			})
			event.Contexts["method"] = c.Method()
			event.Contexts["path"] = c.Path()
			event.Contexts["query"] = c.Fasthttp.QueryArgs().QueryString()
			sentry.CaptureEvent(sentry.NewEvent())
			c.SendString("server internal error")
			c.SendStatus(500)
		}
	}

	app.Use(recover.New(recover.Config{
		Handler: logHandler,
	}))
}
