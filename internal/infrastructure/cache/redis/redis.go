package redis

import (
	"errors"
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

type DeploymentType int

const (
	Standalone DeploymentType = iota
	Cluster
	Sentinel
)

type IStorage interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Delete(key string) error
	Reset() error
	Close() error
	Ping() error
}

func NewClient(cfg *config.Configuration, deployType DeploymentType) (IStorage, error) {
	switch deployType {
	case Standalone:
		client, err := NewStandaloneConn(cfg)
		return client, err
	case Cluster:
		client, err := NewClusterConn(cfg)
		return client, err
	}

	return nil, errors.New("deployment type not found")
}
