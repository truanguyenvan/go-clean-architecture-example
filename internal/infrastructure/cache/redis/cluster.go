package redis

import "github.com/redis/go-redis/v9"

type Cluster struct {
	client *redis.ClusterClient
}
