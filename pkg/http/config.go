package http

import "time"

type config struct {
	ClientTimeout             time.Duration
	DialContextTimeout        time.Duration
	ClientTLSHandshakeTimeout time.Duration
	ClientRetryWaitTime       time.Duration
	RetryCount                time.Duration
}

var DefaultConfig = config{
	ClientTimeout:             clientTimeout,
	DialContextTimeout:        dialContextTimeout,
	ClientTLSHandshakeTimeout: clientTLSHandshakeTimeout,
	ClientRetryWaitTime:       clientRetryWaitTime,
	RetryCount:                retryCount,
}

type Config interface {
	apply(*config)
}
