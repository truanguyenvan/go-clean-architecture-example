package server

import (
	"go-clean-architecture-example/internal/app"
	cragHttp "go-clean-architecture-example/internal/infra/inputports/http/crag"
	"go-clean-architecture-example/internal/infra/interfaceadapters"
	"go-clean-architecture-example/pkg/time"
	"go-clean-architecture-example/pkg/uuid"

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

	// init base
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()

	// Init repositories
	interfaceAdapterServices := interfaceadapters.NewServices()

	// Init useCases
	cragServices := app.NewServices(interfaceAdapterServices.CragRepository, interfaceAdapterServices.NotificationService, up, tp)

	// handlers
	cragHandler := cragHttp.NewHandler(cragServices)

	v1 := g.Group("/api/v1")

	crag := v1.Group("/crag")
	cragHttp.MapCragRoutes(crag, cragHandler)

	return nil
}
