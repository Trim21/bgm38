package web

import (
	"github.com/gofiber/fiber"

	"bgm38/pkg/web/md2bbc"
	"bgm38/pkg/web/utils/handler"
)

func rootRouter(app *fiber.App) {
	app.Post("/md2bbc", handler.LogError(func(ctx *fiber.Ctx) error {
		body := ctx.Fasthttp.Request.Body()
		// ctx.Set("characters", "utf-8")
		// ctx.Set("content-type", "text/plain")
		content := md2bbc.Render(body)
		ctx.SendBytes(content)
		return nil
	}))
}
