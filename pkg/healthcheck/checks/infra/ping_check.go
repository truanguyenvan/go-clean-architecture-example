package checks

import (
	"fmt"
	"go-clean-architecture-example/pkg/healthcheck"
	"io"
	"net/http"
	"sync"
	"time"
)

type Ping struct {
	URL     string
	Method  string
	Timeout int
	client  http.Client
	Body    io.Reader
	Headers map[string]string
}

func NewPingChecker(URL, Method string, Timeout int, Body io.Reader, Headers map[string]string) *Ping {
	if Method == "" {
		Method = "GET"
	}

	if Timeout == 0 {
		Timeout = 500
	}

	pingChecker := Ping{
		URL:     URL,
		Method:  Method,
		Timeout: Timeout,
		Body:    Body,
		Headers: Headers,
	}
	pingChecker.client = http.Client{
		Timeout: time.Duration(Timeout) * time.Millisecond,
	}

	return &pingChecker
}

func (p Ping) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	req, err := http.NewRequest(p.Method, p.URL, p.Body)

	if err != nil {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("can't create new http request: %s -> %s", p.Method, p.URL)
	}

	for key, value := range p.Headers {
		req.Header.Add(key, value)
	}
	resp, err := p.client.Do(req)
	if err != nil || resp.StatusCode >= 300 {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("request failed: %s -> %s", p.Method, p.URL)
	}
	defer resp.Body.Close()

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "ping",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}

}
