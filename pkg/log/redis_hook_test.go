package log

import (
	"testing"
	"time"

	"bgm38/config"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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
	client := getRedisClient()
	_, err := client.Del(redisKey).Result()
	assert.Nil(t, err)
	logrus.AddHook(getHook())
	logrus.Info("test")

	re, err := client.LLen(redisKey).Result()
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.True(t, re == 1, "key length should not be zero, got ", re)
}
