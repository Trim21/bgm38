package config

import (
	"time"

	"bgm38/pkg/utils"
)

// Version App build version
var Version = "development"

// RedisAddr redis ip:port
var RedisAddr = utils.GetEnv("REDIS_HOST", "192.168.1.3") + ":6379"

// RedisPassword redis password, empty string if not set
var RedisPassword = utils.GetEnv("REDIS_PASSWORD", "mypassword")

// AppID bgmtv.tv oauth app id
var AppID = utils.GetEnv("appid", "")

// PROTOCOL nginx protocol http when dev, https when prod
var PROTOCOL = utils.GetEnv("PROTOCOL", "http")

// VirtualHost http HOST when exposed to external
var VirtualHost = utils.GetEnv("VirtualHost", "127.0.0.1")

// MysqlHost mysql host, ip only
var MysqlHost = utils.GetEnv("MYSQL_HOST", "192.168.1.3")

// MysqlAuth mysql authentication, in format `username:password`
var MysqlAuth = utils.GetEnv("MYSQL_AUTH", "root:password")

// RedisSpiderURLKey redis list to read url
var RedisSpiderURLKey = utils.GetEnv("REDIS_SPIDER_DISPATCH_KEY", "bgm_tv_spider:start_urls")

var TimeZone = time.FixedZone("Asia/Shanghai", 3600*8)

var DSN = utils.GetEnv("SENTRY_DSN", "___DSN___")
