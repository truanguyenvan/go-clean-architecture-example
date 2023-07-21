package checks

import (
	"fmt"
	"go-clean-architecture-example/pkg/health"
	"runtime"
	"sync"
	"time"
)

type GarbageCollectionMax struct {
	threshold time.Duration
}

func NewGarbageCollectionChecker(threshold time.Duration) *GarbageCollectionMax {
	return &GarbageCollectionMax{
		threshold: threshold,
	}
}

// GCMaxPauseCheck returns a Check that fails if any recent Go garbage
// collection pause exceeds the provided threshold.
func (gc GarbageCollectionMax) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	thresholdNanoseconds := uint64(gc.threshold.Nanoseconds())
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	for _, pause := range stats.PauseNs {
		if pause > thresholdNanoseconds {
			myStatus = false
			result.Status = false
			errorMessage = fmt.Sprintf("recent GC cycle took %s > %s", time.Duration(pause), gc.threshold)
			break
		}
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "garbage collection",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}
}
