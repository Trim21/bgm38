package bgmtv

import (
	"github.com/gofiber/fiber"

	"bgm38/app/web/utils/handler"
)

func Group(app *fiber.App) {
	v1 := app.Group("/bgm.tv/v1")
	v1.Get("/calendar/:user_id", handler.LogError(userCalendar))
}
