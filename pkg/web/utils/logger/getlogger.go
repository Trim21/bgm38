package logger

import (
	"go.uber.org/zap"

	"bgm38/pkg/utils/log"
)

var _logger *zap.Logger

func GetLogger() *zap.Logger {
	if _logger == nil {
		_logger = getLogger()
	}
	return _logger
}

func getLogger() *zap.Logger {
	return log.BindMeta(log.CreateLogger("bgm38-log-v2"))
}
