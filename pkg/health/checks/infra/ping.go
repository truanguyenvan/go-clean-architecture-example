package checks

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"go-clean-architecture-example/pkg/health"
	"sync"
	"time"
)

type Ping struct {
	URL     string
	Method  string
	Timeout int
	client  *fasthttp.Client
	Body    interface{}
	Headers map[string]string
}

// NewPingChecker : time - millisecond
func NewPingChecker(URL, method string, timeout int, body interface{}, headers map[string]string) *Ping {
	if method == "" {
		method = "GET"
	}

	if timeout == 0 {
		timeout = 500
	}

	pingChecker := Ping{
		URL:     URL,
		Method:  method,
		Timeout: timeout,
		Body:    body,
		Headers: headers,
	}
	pingChecker.client = &fasthttp.Client{}

	return &pingChecker
}

func (p Ping) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(p.URL)

	// set header
	req.Header.SetMethod(p.Method)
	for key, value := range p.Headers {
		req.Header.Add(key, value)
	}
	if p.Method != "GET" && p.Method != "DELETE" {
		byteBody, err := json.Marshal(p.Body)
		if err != nil {
			myStatus = false
			result.Status = false
			errorMessage = fmt.Sprintf("request failed: %s -> %s with error: %s", p.Method, p.URL, err)
		}
		// Set the request body
		req.SetBody(byteBody)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := p.client.Do(req, resp)

	if err != nil || resp.StatusCode() >= 500 {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("request failed: %s -> %s. code: %d. error: %s", p.Method, p.URL, resp.StatusCode(), err)
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "ping",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}

}
