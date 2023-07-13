package redis

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"

	"github.com/redis/go-redis/v9"
)

type StandaloneClient struct {
	client *redis.Client
	ctx    context.Context
}

// WithContext for operate
func (c *StandaloneClient) WithContext(ctx context.Context) *StandaloneClient {
	cp := *c
	cp.ctx = ctx
	return &cp
}

// Get gets the value for the given key.
func (c *StandaloneClient) Get(key string) ([]byte, error) {
	result := c.client.Get(c.ctx, key)
	val, err := result.Bytes()
	if redis.Nil == err {
		return val, fiber.ErrNotFound
	}
	return val, err
}

// Set stores the given value for the given key along with a
func (c *StandaloneClient) Set(key string, val []byte, ttl time.Duration) error {
	result := c.client.Set(c.ctx, key, val, ttl)
	return result.Err()
}

// Delete deletes the value for the given key.
func (c *StandaloneClient) Delete(key string) error {
	result := c.client.Del(c.ctx, key)
	return result.Err()
}

// Reset resets the storage and delete all keys.
func (c *StandaloneClient) Reset() error {
	result := c.client.FlushAll(c.ctx)
	return result.Err()
}

// Close closes the storage and will stop any running garbage
func (c *StandaloneClient) Close() error {
	return c.client.Close()
}

// Ping check connection
func (c *StandaloneClient) Ping() error {
	_, err := c.client.Ping(c.ctx).Result()
	return err
}
