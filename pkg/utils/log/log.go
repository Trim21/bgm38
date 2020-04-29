package log

import (
	"os"

	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"bgm38/config"
)

func CreateLogger() *zap.Logger {
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
		}, "bgm38-log-v2"), zapcore.AddSync(os.Stdout)),
		zap.InfoLevel,
	))
}

func BindMeta(beatName string, logger *zap.Logger) *zap.Logger {
	// var beatName = "bgm38-log-v1"
	// if len(beatNames) > 0 {
	// 	beatName = beatNames[0]
	// }
	return logger.With(zap.Object("@metadata", &Metadata{
		Beat:    beatName,
		Version: config.Version,
	}))
}
