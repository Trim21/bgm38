package web

import (
	"github.com/gofiber/fiber"
	"go.uber.org/zap"

	"bgm38/pkg/web/md2bbc"
	"bgm38/pkg/web/utils/handler"
)

func md2bbcRouter(app *fiber.App) {
	app.Post("/v1/md2bbc", handler.LogError(markdownToBBCode))
}

// @ID markdownToBBCodeV1
// @Summary 转换markdown为bbcode
// @Description 转换markdown为bbcode,
// @Description 有部分bbcode不支持的功能不进行转换
// @Description [一个简单的UI](/asserts/web/md2bbc.html)
// @Accept plain
// @Produce plain
// @Param markdown body string true "待转换的markdown"
// @Success 200 {string} string "text/plain"
// @Router /v1/md2bbc [post]
func markdownToBBCode(ctx *fiber.Ctx, logger *zap.Logger) error {
	body := ctx.Fasthttp.Request.Body()
	ctx.Set("characters", "utf-8")
	ctx.Set("content-type", "text/plain")
	content := md2bbc.Render(body)
	ctx.SendBytes(content)
	return nil
}
