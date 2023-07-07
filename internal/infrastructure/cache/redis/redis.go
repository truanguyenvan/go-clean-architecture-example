package redis

import (
	"go-clean-architecture-example/config"
	"time"
)

const (
	maxRetries      = 5
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 500 * time.Millisecond
	dialTimeout     = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 3 * time.Second
)

type IRedisStorage[C any] interface {
	NewConn(cfg *config.Configuration) error
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration)
	Delete(key string) error
	Reset() error
	Close() error
	Ping() error
}

type RedisStorage[C any] struct {
	client IRedisStorage[C]
}

func NewClient[C any](cfg *config.Configuration) (RedisStorage[C], error) {
	redisStorage := RedisStorage[C]{}
	err := redisStorage.client.NewConn(cfg)
	if err != nil {
		return redisStorage, err
	}
	return redisStorage, nil
}
