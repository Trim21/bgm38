package log

import (
	"os"

	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"bgm38/config"
)

func CreateLogger(key string) *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.TimeKey = "@timestamp"
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		zap.CombineWriteSyncers(NewRedisSink(&redis.Options{
			Addr:     config.RedisAddr,
			Password: config.RedisPassword,
			PoolSize: 3,
		}, key), zapcore.AddSync(os.Stdout)),
		zap.InfoLevel,
	))
}

func BindMeta(logger *zap.Logger) *zap.Logger {
	return logger.With(zap.Object("@metadata", &Metadata{
		Beat:    "bgm38-log-v1",
		Version: config.Version,
	}))
}
