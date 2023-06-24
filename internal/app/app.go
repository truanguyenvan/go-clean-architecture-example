package app

import (
	"github.com/sirupsen/logrus"
	"go-clean-architecture-example/internal/app/crag/commands"
	"go-clean-architecture-example/internal/app/crag/queries"
	"go-clean-architecture-example/internal/commom/metrics"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
)

// Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler queries.GetAllCragsRequestHandler
	GetCragHandler     queries.GetCragRequestHandler
}

// Commands Contains all available command handlers of this app
type Commands struct {
	CreateCragHandler commands.AddCragRequestHandler
	UpdateCragHandler commands.UpdateCragRequestHandler
	DeleteCragHandler commands.DeleteCragRequestHandler
}

type Application struct {
	Queries  Queries
	Commands Commands
}

func NewApplication(cragRepo crag.Repository, ns notification.Service) Application {
	// init base
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}
	return Application{
		Queries: Queries{
			GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(cragRepo, logger, metricsClient),
			GetCragHandler:     queries.NewGetCragRequestHandler(cragRepo),
		},
		Commands: Commands{
			CreateCragHandler: commands.NewAddCragRequestHandler(cragRepo, ns, logger, metricsClient),
			UpdateCragHandler: commands.NewUpdateCragRequestHandler(cragRepo),
			DeleteCragHandler: commands.NewDeleteCragRequestHandler(cragRepo),
		},
	}
}
