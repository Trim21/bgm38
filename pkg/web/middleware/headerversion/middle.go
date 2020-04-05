package headerversion

import (
	"github.com/gofiber/fiber"

	"bgm38/config"
)

func New() func(*fiber.Ctx) {

	return func(c *fiber.Ctx) {
		c.Set("x-server-version", config.Version)
		c.Next()
	}
}
