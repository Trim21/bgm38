package handler

import (
	"github.com/gofiber/fiber"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"bgm38/pkg/utils/log"
	"bgm38/pkg/web/middleware/requestid"
	loggerUtils "bgm38/pkg/web/utils/logger"
)

func LogError(f func(*fiber.Ctx, *zap.Logger) error) func(*fiber.Ctx) {
	var _logger = log.GetLogger()

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
