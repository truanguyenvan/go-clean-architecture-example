package checks

import (
	"io"
	"net/http"
	"time"
)

type PingCheck struct {
	URL     string
	Method  string
	Timeout int
	client  http.Client
	Body    io.Reader
	Headers map[string]string
}

func NewPingCheck(URL, Method string, Timeout int, Body io.Reader, Headers map[string]string) PingCheck {
	if Method == "" {
		Method = "GET"
	}

	if Timeout == 0 {
		Timeout = 500
	}

	pingCheck := PingCheck{
		URL:     URL,
		Method:  Method,
		Timeout: Timeout,
		Body:    Body,
		Headers: Headers,
	}
	pingCheck.client = http.Client{
		Timeout: time.Duration(Timeout) * time.Millisecond,
	}

	return pingCheck
}

func (p PingCheck) Pass() bool {
	req, err := http.NewRequest(p.Method, p.URL, p.Body)

	if err != nil {
		return false
	}

	for key, value := range p.Headers {
		req.Header.Add(key, value)
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode <= 299
}

func (p PingCheck) Name() string {
	return "ping-" + p.URL
}
