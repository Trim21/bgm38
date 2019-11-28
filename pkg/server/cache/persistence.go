package cache

import (
	"github.com/gin-contrib/cache/persistence"
	"time"
)

type LocalFileCachePersistence struct {
	persistence.CacheStore
	cacheRoot string
}

func NewRedisCache(cacheRoot string, defaultExpiration time.Duration) *LocalFileCachePersistence {
	return &LocalFileCachePersistence{cacheRoot: cacheRoot}
}

func (p *LocalFileCachePersistence) Get(key string, value interface{}) error {
	return nil
}

func (p *LocalFileCachePersistence) Set(key string, value interface{}, expire time.Duration) error {
	return nil
}
