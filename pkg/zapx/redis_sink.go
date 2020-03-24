package zapx

import (
	"bytes"

	"github.com/go-redis/redis/v7"
)

var (
	redisSinkInsts = map[string]RedisSink{}
)

const (
	redisPubSubType      = "channel"
	redisDefaultType     = "list"
	redisClusterNodesKey = "nodes"
	redisDefaultPwd      = ""
	redisDefaultKey      = "just_a_test_key"
	kafkaDefaultTopic    = "just_a_test_topic"
	kafkaAsyncKey        = "isAsync"

	cachedLogLineEnding = ","
	cachedLogMessageKey = ""
	cachedLogSinkURL    = "stderr"
	globalKeyPrefix     = "__"
)

// MultiError multiple error
type MultiError []error

func (p MultiError) Error() string {
	var errBuf bytes.Buffer
	for _, err := range p {
		errBuf.WriteString(err.Error())
		errBuf.WriteByte('\n')
	}
	return errBuf.String()
}

type RedisSink struct {
	redisClient *redis.Client
	key         string
	typee       string
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
		typee:       redisDefaultType,
		isCluster:   false,
	}

}
