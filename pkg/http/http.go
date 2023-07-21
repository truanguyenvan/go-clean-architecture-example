package http

import (
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
)

type Client interface {
	MakeRequest() *resty.Request
	RestyClient() *resty.Client
}

type client struct {
	restyClient *resty.Client
}

// New  Init resty client
// example:
//
//	configs := []Config{
//		ClientTimeout(3),
//		DebugMode(true),
//	}
//	New(configs...)
func New(configs ...Config) Client {
	// init default config
	cfg := defaultConfig
	for _, config := range configs {
		config.apply(&cfg)
	}

	t := &http.Transport{
		DialContext:         (&net.Dialer{Timeout: cfg.DialContextTimeout}).DialContext,
		TLSHandshakeTimeout: cfg.ClientTLSHandshakeTimeout,
	}

	cli := resty.New().
		SetDebug(cfg.DebugMode).
		SetTimeout(cfg.ClientTimeout).
		SetRetryCount(cfg.RetryCount).
		SetRetryWaitTime(cfg.ClientRetryWaitTime).
		SetTransport(t).
		AddRetryCondition(cfg.RetryCondition)
	return &client{
		restyClient: cli,
	}
}

func (cli client) MakeRequest() *resty.Request {
	return cli.restyClient.R()
}

func (cli client) RestyClient() *resty.Client {
	return cli.restyClient
}
