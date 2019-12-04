package config

import (
	"bgm38/pkg/utils"
)

//RedisAddr redis ip:port
var RedisAddr = utils.GetEnv("REDIS_HOST", "127.0.0.1") + ":6379"

//RedisPassword redis password, empty string if not set
var RedisPassword = utils.GetEnv("REDIS_PASSWORD", "")

//AppID bgm.tv oauth app id
var AppID = utils.GetEnv("appid", "")

// PROTOCOL nginx protocol http when dev, https when prod
var PROTOCOL = utils.GetEnv("PROTOCOL", "http")

// VirtualHost http HOST when exposed to external
var VirtualHost = utils.GetEnv("VirtualHost", "127.0.0.1")

//MysqlHost mysql host, ip only
var MysqlHost = utils.GetEnv("MYSQL_HOST", "192.168.1.4")

//MysqlAuth mysql authentication, in format `username:password`
var MysqlAuth = utils.GetEnv("MYSQL_AUTH", "root:password")
