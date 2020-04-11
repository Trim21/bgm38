package log

import (
	"github.com/go-redis/redis/v7"
)

type RedisSink struct {
	redisClient *redis.Client
	key         string
	isCluster   bool
}

// Close implement zap.Sink func Close
func (p RedisSink) Close() error {
	return p.redisClient.Close()
}

// Write implement zap.Sink func Write
func (p RedisSink) Write(b []byte) (n int, err error) {
	err = p.redisClient.RPush(p.key, string(b)).Err()
	return len(b), err
}

// Sync implement zap.Sink func Sync
func (p RedisSink) Sync() error {
	return nil
}

func NewRedisSink(options *redis.Options, redisKey string) *RedisSink {

	var redisClient = redis.NewClient(options)
	return &RedisSink{
		redisClient: redisClient,
		key:         redisKey,
		isCluster:   false,
	}

}
