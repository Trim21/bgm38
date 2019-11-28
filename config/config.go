package config

import (
	"bgm38/pkg/utils"
)

var RedisAddr = utils.GetEnv("REDIS_HOST", "127.0.0.1") + ":6379"
var RedisPassword = utils.GetEnv("REDIS_PASSWORD", "")
