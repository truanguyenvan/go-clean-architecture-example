package redis

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"time"
)

type ClusterClient struct {
	client *redis.ClusterClient
	ctx    context.Context
}

// WithContext for operate
func (c *ClusterClient) WithContext(ctx context.Context) *ClusterClient {
	cp := *c
	cp.ctx = ctx
	return &cp
}

// Get gets the value for the given key.
func (c *ClusterClient) Get(key string) ([]byte, error) {
	result := c.client.Get(c.ctx, key)
	val, err := result.Bytes()
	if redis.Nil == err {
		return val, fiber.ErrNotFound
	}
	return val, err
}

// Set stores the given value for the given key along with a
func (c *ClusterClient) Set(key string, val []byte, ttl time.Duration) error {
	result := c.client.Set(c.ctx, key, val, ttl)
	return result.Err()
}

// Delete deletes the value for the given key.
func (c *ClusterClient) Delete(key string) error {
	result := c.client.Del(c.ctx, key)
	return result.Err()
}

// Reset resets the storage and delete all keys.
func (c *ClusterClient) Reset() error {
	result := c.client.FlushAll(c.ctx)
	return result.Err()
}

// Close closes the storage and will stop any running garbage
func (c *ClusterClient) Close() error {
	return c.client.Close()
}

// Ping check connection
func (c *ClusterClient) Ping() error {
	err := c.client.ForEachShard(c.ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		return err
	}

	err = c.client.ForEachSlave(c.ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		return err
	}

	return nil
}
