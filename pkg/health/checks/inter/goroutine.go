package checks

import (
	"fmt"
	"go-clean-architecture-example/pkg/health"
	"runtime"
	"sync"
	"time"
)

type Goroutine struct {
	threshold int
}

func NewGoroutineChecker(threshold int) *Goroutine {
	return &Goroutine{
		threshold: threshold,
	}
}

// GoroutineCountCheck returns a Check that fails if too many goroutines are
// running (which could indicate a resource leak).
func (gr Goroutine) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)
	count := runtime.NumGoroutine()
	if count > gr.threshold {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("too many goroutines (%d > %d)", count, gr.threshold)
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "goroutine",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}
}
