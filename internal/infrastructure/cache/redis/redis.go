package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-clean-architecture-example/config"
	"strings"
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

// NewClusterConn returns new redis client
func NewClusterConn(cfg *config.Configuration) (*ClusterClient, error) {

	adds := strings.Split(cfg.RedisCluster.Address, cfg.RedisCluster.Delimiter)

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:           adds,
		ReadOnly:        cfg.RedisCluster.ReadOnly,
		MinIdleConns:    cfg.RedisCluster.MinIdleCons,
		PoolSize:        cfg.RedisCluster.PoolSize,
		PoolTimeout:     time.Duration(cfg.RedisCluster.PoolTimeout) * time.Second,
		Password:        cfg.RedisCluster.Password, // no password set
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
	})

	clusterClient := &ClusterClient{
		client: client,
		ctx:    context.Background(),
	}
	if err := clusterClient.Ping(); err != nil {
		return nil, err
	}

	return clusterClient, nil
}

// NewStandaloneConn returns new redis client
func NewStandaloneConn(cfg *config.Configuration) (*StandaloneClient, error) {
	redisHost := cfg.Redis.Address

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:            redisHost,
		MinIdleConns:    cfg.Redis.MinIdleCons,
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
	})

	standaloneClient := &StandaloneClient{
		client: client,
		ctx:    context.Background(),
	}
	if err := standaloneClient.Ping(); err != nil {
		return nil, err
	}
	return standaloneClient, nil
}
