package db

import (
	"github.com/go-redis/redis/v7"

	"bgm38/config"
)

// Redis redis client
var Redis = redis.NewClient(&redis.Options{
	Addr:     config.RedisAddr,
	Password: config.RedisPassword, // no password set
	// Mysql:       0,            // use default Mysql
})
