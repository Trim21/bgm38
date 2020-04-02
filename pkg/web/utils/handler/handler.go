package handler

import (
	"github.com/gofiber/fiber"
	"go.uber.org/zap"

	"bgm38/pkg/utils/log"
	logger2 "bgm38/pkg/web/utils/logger"
)

func LogError(f func(*fiber.Ctx, *zap.Logger) error) func(*fiber.Ctx) {
	var _logger = log.GetLogger()

	return func(ctx *fiber.Ctx) {
		s := logger2.HeaderFields(ctx)
		// rid := ctx.Fasthttp.Response.Header.Len(fiber.HeaderXRequestID)
		logger := _logger.With(s)
		// ctx.Locals()
		err := f(ctx, logger)
		if err != nil {
			logger.Error(err.Error())
		}

	}
}
