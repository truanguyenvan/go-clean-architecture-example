package cache

import (
	"time"

	"github.com/go-redis/redis/v8"

	"go-clean-architecture-example/config"
)

const (
	maxRetries      = 5
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 500 * time.Millisecond
	dialTimeout     = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 3 * time.Second
	minIdleConns    = 20
	poolTimeout     = 6 * time.Second
	idleTimeout     = 12 * time.Second
)

// Returns new redis client
func NewRedisClient(cfg *config.Configuration) *redis.Client {
	redisHost := cfg.Redis.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:            redisHost,
		MinIdleConns:    cfg.Redis.MinIdleConns,
		PoolSize:        cfg.Redis.PoolSize,
		PoolTimeout:     time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:        cfg.Redis.Password, // no password set
		DB:              cfg.Redis.DB,
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		IdleTimeout:     idleTimeout, // use default DB
	})

	return client
}
