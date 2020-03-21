package log

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"bgm38/config"
)

const redisKey = "bgm38 log"

func getHook() *RedisHook {
	options := &redis.Options{
		Addr:        config.RedisAddr,
		Password:    config.RedisPassword,
		PoolSize:    3,
		DialTimeout: time.Second,
	}
	hook := NewRedisHook(options, redisKey)
	return hook
}

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        config.RedisAddr,
		Password:    config.RedisPassword,
		DialTimeout: time.Second,
	})

}

func TestRedisHook_Fire(t *testing.T) {
	logrus.SetReportCaller(true)
	client := getRedisClient()
	err := client.Del(redisKey).Err()
	assert.Nil(t, err)
	logrus.AddHook(getHook())
	logrus.Info("test")

	re, err := client.LLen(redisKey).Result()
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.True(t, re == 1, "key length should not be zero, got ", re)
}
