package probes

import (
	"go-clean-architecture-example/config"
	"go-clean-architecture-example/pkg/health"
	infraHealthcheck "go-clean-architecture-example/pkg/health/checks/infra"
	interHealthcheck "go-clean-architecture-example/pkg/health/checks/inter"
	"time"
)

type HealthCheckApplication interface {
	LiveEndpoint() healthcheck.ApplicationHealthDetailed
	ReadyEndpoint() healthcheck.ApplicationHealthDetailed
}

type healthCheckApplication struct {
	healthcheckApp healthcheck.IHealthCheckApplication
}

func NewHealthChecker(configuration *config.Configuration) HealthCheckApplication {
	// init healcheck app
	healthcheckApp := healthcheck.NewApplication(configuration.Server.Name, configuration.Server.AppVersion)

	// add liveness checker
	grChecker := interHealthcheck.NewGoroutineChecker(configuration.Server.GrRunningThreshold)
	healthcheckApp.AddLivenessCheck("goroutine checker", grChecker)

	pauseTimethreshold := time.Duration(configuration.Server.GcPauseThreshold) * time.Millisecond
	gcChecker := interHealthcheck.NewGarbageCollectionChecker(pauseTimethreshold)
	healthcheckApp.AddLivenessCheck("garbage collection checker", gcChecker)

	//envChecker := interHealthcheck.NewEnvChecker("test env", "")
	//healthcheckApp.AddLivenessCheck("environment variable checker", envChecker)

	// add readieness checker
	gogleChecker := infraHealthcheck.NewPingChecker("https://google.com.vn", "Get", 5, nil, nil)
	healthcheckApp.AddReadinessCheck("google ping checker", gogleChecker)

	return &healthCheckApplication{
		healthcheckApp: healthcheckApp,
	}
}

func (app healthCheckApplication) LiveEndpoint() healthcheck.ApplicationHealthDetailed {
	return app.healthcheckApp.LiveChecker()
}

func (app healthCheckApplication) ReadyEndpoint() healthcheck.ApplicationHealthDetailed {
	return app.healthcheckApp.ReadyChecker()
}
