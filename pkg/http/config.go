package http

import (
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	clientTimeout             = 5 * time.Second
	dialContextTimeout        = 5 * time.Second
	clientTLSHandshakeTimeout = 5 * time.Second
	clientRetryWaitTime       = 300 * time.Millisecond
	retryCount                = 3
	debugMode                 = false
)

type config struct {
	ClientTimeout             time.Duration
	DialContextTimeout        time.Duration
	ClientTLSHandshakeTimeout time.Duration
	ClientRetryWaitTime       time.Duration
	RetryCount                int
	RetryCondition            func(r *resty.Response, err error) bool
	DebugMode                 bool
}

var defaultConfig = config{
	ClientTimeout:             clientTimeout,
	DialContextTimeout:        dialContextTimeout,
	ClientTLSHandshakeTimeout: clientTLSHandshakeTimeout,
	ClientRetryWaitTime:       clientRetryWaitTime,
	RetryCount:                retryCount,
	RetryCondition:            defaultRetryCondition,
	DebugMode:                 debugMode,
}

type Config interface {
	apply(*config)
}

type funcConfig struct {
	f func(*config)
}

func (fdo *funcConfig) apply(do *config) {
	fdo.f(do)
}

func newFuncConfig(f func(*config)) *funcConfig {
	return &funcConfig{
		f: f,
	}
}

func ClientTimeout(val time.Duration) Config {
	return newFuncConfig(func(c *config) {
		c.ClientTimeout = val
	})
}

func DialContextTimeout(val time.Duration) Config {
	return newFuncConfig(func(c *config) {
		c.DialContextTimeout = val
	})
}

func ClientTLSHandshakeTimeout(val time.Duration) Config {
	return newFuncConfig(func(c *config) {
		c.ClientTLSHandshakeTimeout = val
	})
}

func ClientRetryWaitTime(val time.Duration) Config {
	return newFuncConfig(func(c *config) {
		c.ClientRetryWaitTime = val
	})
}

func RetryCount(val int) Config {
	return newFuncConfig(func(c *config) {
		c.RetryCount = val
	})
}

func RetryCondition(f func(r *resty.Response, err error) bool) Config {
	return newFuncConfig(func(c *config) {
		c.RetryCondition = f
	})
}

func DebugMode(mode bool) Config {
	return newFuncConfig(func(c *config) {
		c.DebugMode = mode
	})
}
