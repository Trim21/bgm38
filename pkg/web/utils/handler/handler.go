package handler

import (
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/gofiber/fiber"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"bgm38/config"
	"bgm38/pkg/web/utils"
	"bgm38/pkg/zapx"
)

var _logger *zap.Logger

func LogError(f func(*fiber.Ctx, *zap.Logger) error) func(*fiber.Ctx) {
	if _logger == nil {
		_logger = getLogger()
	}

	return func(ctx *fiber.Ctx) {
		s := utils.HeaderFields(ctx)
		// rid := ctx.Fasthttp.Response.Header.Len(fiber.HeaderXRequestID)
		logger := _logger.With(s)
		// ctx.Locals()
		err := f(ctx, logger)
		if err != nil {
			logger.Error(err.Error())
		}

	}
}

func getLogger() *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		zap.CombineWriteSyncers(zapx.NewRedisSink(&redis.Options{
			Addr:     config.RedisAddr,
			Password: config.RedisPassword,
			PoolSize: 3,
		}, "bgm38 log v2"), zapcore.AddSync(os.Stdout)),
		zap.InfoLevel,
	))
}
