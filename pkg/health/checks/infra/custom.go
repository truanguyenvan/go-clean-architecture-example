package checks

import (
	"fmt"
	"go-clean-architecture-example/pkg/health"
	"sync"
	"time"
)

type Custom struct {
	handler func() error
}

func NewCustomChecker(handler func() error) *Custom {
	return &Custom{handler: handler}
}

func (c *Custom) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	if c.handler == nil {
		myStatus = false
		result.Status = false
		errorMessage = "handler is nil"
	}

	err := c.handler()
	if c.handler == nil {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprint(err)
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "custom",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}

}
