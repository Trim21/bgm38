package config

import (
	"bgm38/pkg/utils"
)

var RedisAddr = utils.GetEnv("REDIS_HOST", "127.0.0.1") + ":6379"
var RedisPassword = utils.GetEnv("REDIS_PASSWORD", "")
var AppId = utils.GetEnv("appid", "")

var PROTOCOL = utils.GetEnv("PROTOCOL", "http")
var VIRTUAL_HOST = utils.GetEnv("VIRTUAL_HOST", "127.0.0.1")

var MysqlHost = utils.GetEnv("MYSQL_HOST", "192.168.1.4")
var MysqlAuth = utils.GetEnv("MYSQL_AUTH", "root:password")
