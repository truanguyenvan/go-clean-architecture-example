package healthcheck

import (
	"sync"
	"time"
)

func NewApplication(name, version string) IHealthCheckApplication {
	app := &HealthCheckApplication{
		livenessChecks:  make(map[string]ICheckHandler),
		readinessChecks: make(map[string]ICheckHandler),
		Name:            name,
		Version:         version,
	}
	return app
}

func (app *HealthCheckApplication) collectChecks(checks map[string]ICheckHandler) ApplicationHealthDetailed {
	var (
		start     = time.Now()
		wg        sync.WaitGroup
		checklist = make(chan Integration, len(checks))
		result    = ApplicationHealthDetailed{
			Name:         app.Name,
			Version:      app.Version,
			Status:       true,
			Date:         start.Format(time.RFC3339),
			Duration:     0,
			Integrations: []Integration{},
		}
	)
	wg.Add(len(checks))
	for name, handler := range checks {
		go handler.Check(name, &result, &wg, checklist)
	}

	go func() {
		wg.Wait()
		close(checklist)
		result.Duration = time.Since(start).Seconds()
	}()

	for chk := range checklist {
		result.Integrations = append(result.Integrations, chk)
	}
	return result
}
func (app *HealthCheckApplication) LiveChecker() ApplicationHealthDetailed {
	return app.collectChecks(app.livenessChecks)
}

func (app *HealthCheckApplication) ReadyChecker() ApplicationHealthDetailed {
	return app.collectChecks(app.readinessChecks)
}

func (app *HealthCheckApplication) AddLivenessCheck(name string, check ICheckHandler) {
	app.checksMutex.Lock()
	defer app.checksMutex.Unlock()
	app.livenessChecks[name] = check
}

func (app *HealthCheckApplication) AddReadinessCheck(name string, check ICheckHandler) {
	app.checksMutex.Lock()
	defer app.checksMutex.Unlock()
	app.readinessChecks[name] = check
}
