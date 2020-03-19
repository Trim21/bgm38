package log

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

// RedisHook to sends logs to Redis server
type RedisHook struct {
	redisClient *redis.Client
	RedisKey    string
}

// Fire is called when a log event is fired.
func (hook *RedisHook) Fire(entry *logrus.Entry) error {
	var msg, err = fromEntry(entry)
	if err != nil {
		fmt.Printf("error when build logredis to %s", err.Error())
		return err
	}
	b, err := msg.MarshalMsg(nil)
	if err != nil {
		fmt.Printf("error when log to redis %s", err.Error())
		return err
	}

	hook.redisClient.RPush(hook.RedisKey, b)
	return nil
}

// Levels returns the available logging levels.
func (hook *RedisHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func NewRedisHook(options *redis.Options, redisKey string) *RedisHook {
	var redisClient = redis.NewClient(options)
	return &RedisHook{
		redisClient: redisClient,
		RedisKey:    redisKey,
	}

}
