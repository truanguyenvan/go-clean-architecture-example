package http

import (
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
)

type Client interface {
	MakeRequest() *resty.Request
	Resty() *resty.Client
	SetRestyClient(*resty.Client)
}

type client struct {
	*resty.Client
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
		cli,
	}
}

func (cli *client) MakeRequest() *resty.Request {
	return cli.R()
}

func (cli *client) Resty() *resty.Client {
	return cli.Client
}

func (cli *client) SetRestyClient(resClient *resty.Client) {
	cli.Client = resClient
}
