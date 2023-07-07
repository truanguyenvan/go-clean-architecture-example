package redis

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"

	"github.com/redis/go-redis/v9"

	"go-clean-architecture-example/config"
)

type StandaloneClient struct {
	client *redis.Client
}

// Returns new redis client
func (c *StandaloneClient) NewConn(cfg *config.Configuration) error {
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
	})

	c.client = client

	if err := c.Ping(); err != nil {
		return err
	}
	return nil
}

// Get gets the value for the given key.
// It returns ErrNotFound if the storage does not contain the key.
func (c *StandaloneClient) Get(key string) ([]byte, error) {
	result := c.client.Get(context.Background(), key)
	val, err := result.Bytes()
	if redis.Nil == err {
		return val, fiber.ErrNotFound
	}
	return val, err
}

// Set stores the given value for the given key along with a
// time-to-live expiration value, 0 means live for ever
// Empty key or value will be ignored without an error.
func (c *StandaloneClient) Set(key string, val []byte, ttl time.Duration) error {
	result := c.client.Set(context.Background(), key, val, ttl)
	return result.Err()
}

// Delete deletes the value for the given key.
// It returns no error if the storage does not contain the key,
func (c *StandaloneClient) Delete(key string) error {
	result := c.client.Del(context.Background(), key)
	return result.Err()
}

// Reset resets the storage and delete all keys.
func (c *StandaloneClient) Reset() error {
	result := c.client.FlushAll(context.Background())
	return result.Err()
}

// Close closes the storage and will stop any running garbage
// collectors and open connections.
func (c *StandaloneClient) Close() error {
	return c.client.Close()
}

// Ping check connection
func (c *StandaloneClient) Ping() error {
	_, err := c.client.Ping(context.Background()).Result()
	return err
}
