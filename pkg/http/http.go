package http

import (
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"time"
)

const (
	clientTimeout             = 5 * time.Second
	dialContextTimeout        = 5 * time.Second
	clientTLSHandshakeTimeout = 5 * time.Second
	clientRetryWaitTime       = 300 * time.Millisecond
	retryCount                = 3
)

type Client struct {
	client *resty.Client
}

func New(configs ...Config) *resty.Client {
	// init default config
	cfg := DefaultConfig
	for _, config := range configs {
		config.apply(&cfg)
	}

	t := &http.Transport{
		DialContext:         (&net.Dialer{Timeout: dialContextTimeout}).DialContext,
		TLSHandshakeTimeout: clientTLSHandshakeTimeout,
	}

	client := resty.New().
		SetDebug(debugMode).
		SetTimeout(clientTimeout).
		SetRetryCount(retryCount).
		SetRetryWaitTime(clientRetryWaitTime).
		SetTransport(t).
		AddRetryCondition(retryCondition)

	return client
}
