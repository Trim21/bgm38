package handler

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/requestid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	loggerUtils "bgm38/app/web/utils/logger"
)

func LogError(f func(*fiber.Ctx, *zap.Logger) error) func(*fiber.Ctx) {
	var _logger = loggerUtils.GetLogger()

	return func(ctx *fiber.Ctx) {
		s := loggerUtils.HeaderFields(ctx)
		logger := _logger.With(s, getRequestID(ctx))

		err := f(ctx, logger)

		if err != nil {
			logger.Error(err.Error())
		}

	}
}

func getRequestID(c *fiber.Ctx) zapcore.Field {
	return zap.String(fiber.HeaderXRequestID, requestid.Get(c))
}
