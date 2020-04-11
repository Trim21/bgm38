package requestid

import (
	"net/http"
	"strconv"
	"testing"
	"unsafe"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func byteToString(haystack []byte) string {
	return *(*string)(unsafe.Pointer(&haystack))
}

func extractRequestID(ctx *fiber.Ctx) string {
	return byteToString(ctx.Fasthttp.Response.Header.Peek(fiber.HeaderXRequestID))
}

func BenchmarkRequestIDInHeader(b *testing.B) {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) {
		// Get value from RequestID
		rid := c.Get(fiber.HeaderXRequestID)
		// Create new ID
		if rid == "" {
			rid = uuid.New().String()
		}
		// Set X-Request-ID
		c.Set(fiber.HeaderXRequestID, rid)
		c.Next()
	})
	app.Get("/", func(ctx *fiber.Ctx) {
		common(ctx)
		ctx.SendString(extractRequestID(ctx))
	})
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.Test(req, -1)
	}
	b.StopTimer()
}

func BenchmarkRequestIDInLocals(b *testing.B) {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) {
		// Get value from RequestID
		rid := c.Get(fiber.HeaderXRequestID)
		// Create new ID
		if rid == "" {
			rid = uuid.New().String()
		}
		c.Locals(fiber.HeaderXRequestID, rid)
		c.Next()
		// Set X-Request-ID
		c.Set(fiber.HeaderXRequestID, rid)
	})
	app.Get("/", func(ctx *fiber.Ctx) {
		common(ctx)
		ctx.SendString(Get(ctx))
	})
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.Test(req, -1)
	}
	b.StopTimer()
}

func common(ctx *fiber.Ctx) {
	for i := 0; i <= 100; i++ {
		ctx.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
}
