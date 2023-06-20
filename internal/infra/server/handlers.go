package server

import (
	"go-clean-architecture-example/internal/app"
	cragHttp "go-clean-architecture-example/internal/infra/inputports/http/crag"
	"go-clean-architecture-example/internal/infra/interfaceadapters"
	"go-clean-architecture-example/pkg/healthcheck"
	interHealcheck "go-clean-architecture-example/pkg/healthcheck/checks/inter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Map Server Handlers
func (s *Server) MapHandlers(g *gin.Engine) error {
	s.gin.Use(
		gin.Recovery(),
	)
	if s.cfg.Server.Mode == "Development" {
		s.gin.Use(
			gin.Logger(),
		)
	}

	// init healcheck app
	healthcheckApp := healthcheck.NewApplication(s.cfg.Server.Name, s.cfg.Server.AppVersion)
	// add liveness checker
	const GR_RUNING_THRESHOLD = 100 //  threshold for goroutines are running (which could indicate a resource leak).
	grHealthChecker := interHealcheck.NewGoroutineChecker(GR_RUNING_THRESHOLD)
	healthcheckApp.AddLivenessCheck("goroutine checker", grHealthChecker)

	const GC_PAUSE_TIME_THRESHOLD = time.Millisecond * 10 //  threshold threshold garbage collection pause exceeds.
	gcHealthChecker := interHealcheck.NewGarbageCollectionChecker(GC_PAUSE_TIME_THRESHOLD)
	healthcheckApp.AddLivenessCheck("garbage collection checker", gcHealthChecker)

	envHeathChecker := interHealcheck.NewEnvChecker("123", "")
	healthcheckApp.AddLivenessCheck("environment variable checker", envHeathChecker)

	// health check endpoint
	g.GET("/liveness", func(c *gin.Context) {
		result := healthcheckApp.LiveEndpoint()
		if result.Status {
			c.JSON(http.StatusOK, result)
			return
		}
		c.JSON(http.StatusServiceUnavailable, result)
	})

	g.GET("/readiness", func(c *gin.Context) {
		result := healthcheckApp.ReadyEndpoint()
		if result.Status {
			c.JSON(http.StatusOK, result)
			return
		}
		c.JSON(http.StatusServiceUnavailable, result)
	})

	// Init repositories
	interfaceAdapterServices := interfaceadapters.NewServices()

	// Init useCases
	cragServices := app.NewServices(interfaceAdapterServices.CragRepository, interfaceAdapterServices.NotificationService)

	// handlers
	cragHandler := cragHttp.NewHandler(cragServices)

	v1 := g.Group("/api/v1")

	crag := v1.Group("/crag")
	cragHttp.MapCragRoutes(crag, cragHandler)

	return nil
}
