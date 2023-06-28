package app

import (
	"go-clean-architecture-example/internal/app/crag/commands"
	"go-clean-architecture-example/internal/app/crag/queries"
	"go-clean-architecture-example/internal/common/metrics"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
	"go-clean-architecture-example/pkg/logger"
	"go-clean-architecture-example/pkg/time"
	"go-clean-architecture-example/pkg/uuid"
)

// Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler queries.GetAllCragsRequestHandler
	GetCragHandler     queries.GetCragRequestHandler
}

// Commands Contains all available command handlers of this app
type Commands struct {
	AddCragHandler    commands.AddCragRequestHandler
	UpdateCragHandler commands.UpdateCragRequestHandler
	DeleteCragHandler commands.DeleteCragRequestHandler
}

type Application struct {
	Queries  Queries
	Commands Commands
}

func NewApplication(cragRepo crag.Repository, ns notification.Service, logger logger.Logger) Application {
	// init base
	metricsClient := metrics.NoOp{}
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	return Application{
		Queries: Queries{
			GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(cragRepo, logger, metricsClient),
			GetCragHandler:     queries.NewGetCragRequestHandler(cragRepo, logger, metricsClient),
		},
		Commands: Commands{
			AddCragHandler:    commands.NewAddCragRequestHandler(up, tp, cragRepo, ns, logger, metricsClient),
			UpdateCragHandler: commands.NewUpdateCragRequestHandler(cragRepo, logger, metricsClient),
			DeleteCragHandler: commands.NewDeleteCragRequestHandler(cragRepo, logger, metricsClient),
		},
	}
}
