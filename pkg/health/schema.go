package healthcheck

import "sync"

type IHealthCheckApplication interface {
	AddLivenessCheck(name string, check ICheckHandler)

	AddReadinessCheck(name string, check ICheckHandler)

	LiveChecker() ApplicationHealthDetailed

	ReadyChecker() ApplicationHealthDetailed
}

type ICheckHandler interface {
	Check(name string, result *ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan Integration)
}

// Integration is the type result for requests
type Integration struct {
	Name         string  `json:"name"`
	Kind         string  `json:"kind"`
	Status       bool    `json:"status"`
	ResponseTime float64 `json:"response_time"` //in seconds
	URL          string  `json:"url,omitempty"`
	Error        string  `json:"error,omitempty"` // error.Error()
}

// ApplicationHealthDetailed used to check all application integrations and return status of each of then
type ApplicationHealthDetailed struct {
	Name         string        `json:"name,omitempty"`
	Status       bool          `json:"status"`
	Version      string        `json:"version,omitempty"`
	Date         string        `json:"date"`
	Duration     float64       `json:"duration"`
	Integrations []Integration `json:"integrations,omitempty"`
}

// ApplicationConfig is a config contract to init health caller
type HealthCheckApplication struct {
	Name            string
	Version         string
	checksMutex     sync.RWMutex
	livenessChecks  map[string]ICheckHandler
	readinessChecks map[string]ICheckHandler
}
