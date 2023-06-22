//go:build wireinject
// +build wireinject

package server

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/config"
	"go-clean-architecture-example/internal/api"
	"go-clean-architecture-example/internal/app"
	"go-clean-architecture-example/internal/infrastructure/notification"
	"go-clean-architecture-example/internal/infrastructure/persistence"

	"go-clean-architecture-example/internal/exception"
	"go-clean-architecture-example/internal/router"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"go-clean-architecture-example/docs"
	"go-clean-architecture-example/pkg/healthcheck"
	interHealcheck "go-clean-architecture-example/pkg/healthcheck/checks/inter"
	"go-clean-architecture-example/pkg/logger"
	"os"
	"time"
)

// Server struct
type Server struct {
	app    *fiber.App
	cfg    *config.Configuration
	logger logger.Logger
}

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		router.Set,
		api.Set,
		app.Set,
		persistence.Set,
		notification.Set,
	)))
}

// @title  My SERVER
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email minkj1992@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
func NewServer(
	cfg *config.Configuration,
	cragRouter router.CragRouter) *Server {

	// init logger
	logger := logger.NewApiLogger(cfg)

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.CustomErrorHandler,
		ReadTimeout:  time.Second * cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * cfg.Server.WriteTimeout,
	})

	app.Use(fiberlog.New(fiberlog.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover.New())

	setSwagger(cfg.Server.BaseURI)
	app.Get("/swagger/*", swagger.HandlerDefault)

	// init healcheck app
	healthcheckApp := healthcheck.NewApplication(cfg.Server.Name, cfg.Server.AppVersion)
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
	app.Get("/liveness", func(c *fiber.Ctx) error {
		result := healthcheckApp.LiveEndpoint()
		if result.Status {
			return c.Status(fiber.StatusOK).JSON(result)
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	})

	app.Get("/readiness", func(c *fiber.Ctx) error {
		result := healthcheckApp.ReadyEndpoint()
		if result.Status {
			return c.Status(fiber.StatusOK).JSON(result)
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	cragRouter.Init(&v1)

	return &Server{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Configuration {
	return serv.cfg
}

func (serv Server) Logger() logger.Logger {
	return serv.logger
}

func setSwagger(baseURI string) {
	docs.SwaggerInfo.Title = "Go Clean Architecture Example ✈️"
	docs.SwaggerInfo.Description = "This is a go clean architecture example."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = baseURI
	docs.SwaggerInfo.BasePath = "/api/v1"
}
