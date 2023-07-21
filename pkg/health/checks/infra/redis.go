package checks

import (
	"context"
	"fmt"
	"go-clean-architecture-example/pkg/health"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedisChecker(client *redis.Client) *Redis {
	return &Redis{Client: client}
}

func (r *Redis) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	if r.Client == nil {
		myStatus = false
		result.Status = false
		errorMessage = "connection is nil"
	}

	_, err := r.Client.Ping(context.Background()).Result()
	if err != nil {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("ping failed: %s", err)
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "redis",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}

}
