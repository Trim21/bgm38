package requestid

import (
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

// New adds an indentifier to the request using the `X-Request-ID` header
func New() func(*fiber.Ctx) {
	// Init config
	generator := func() string {
		return uuid.New().String()
	}
	// Return middleware handler
	return func(c *fiber.Ctx) {
		// Get value from RequestID
		rid := c.Get(fiber.HeaderXRequestID)
		// Create new ID
		if rid == "" {
			rid = generator()
		}
		c.Locals(fiber.HeaderXRequestID, rid)
		c.Next()
		// Set X-Request-ID
		c.Set(fiber.HeaderXRequestID, rid)
	}
}

func Get(c *fiber.Ctx) string {
	rid := c.Locals(fiber.HeaderXRequestID)
	if v, ok := rid.(string); ok {
		return v
	}
	return ""
}
