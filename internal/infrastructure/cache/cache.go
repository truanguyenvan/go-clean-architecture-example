package cache

import (
	"github.com/google/wire"
	"go-clean-architecture-example/config"
	"go-clean-architecture-example/pkg/cache/redis"
	"time"
)

var RedisSet = wire.NewSet(
	NewRedisCache,
)

var MemCacheSet = wire.NewSet(
	NewMemCache,
)

type Engine interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Delete(key string) error
	Reset() error
	Close() error
	Ping() error
}

func NewRedisCache(configuration *config.Configuration) (Engine, error) {
	switch configuration.Server.CacheDeploymentType {
	case 1:
		client, err := redis.NewStandaloneConn(configuration)
		return client, err
	case 2:
		client, err := redis.NewClusterConn(configuration)
		return client, err
	default:
		client, err := redis.NewStandaloneConn(configuration)
		return client, err
	}
}

func NewMemCache(configuration config.Configuration) (Engine, error) {
	return nil, nil
}
