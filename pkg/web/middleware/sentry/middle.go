package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"

	"bgm38/config"
	"bgm38/pkg/web/middleware/requestid"
)

func New() func(ctx *fiber.Ctx) {

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
			event.Contexts["request-id"] = requestid.Get(c)
			sentry.CaptureEvent(sentry.NewEvent())
			c.SendString("server internal error")
			c.SendStatus(500)
		}
	}

	return recover.New(recover.Config{
		Handler: logHandler,
	})
}
