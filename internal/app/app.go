package app

import (
	"go-clean-architecture-example/internal/app/crag/commands"
	"go-clean-architecture-example/internal/app/crag/queries"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
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
	CreateCragHandler commands.CreateCragRequestHandler
	UpdateCragHandler commands.UpdateCragRequestHandler
	DeleteCragHandler commands.DeleteCragRequestHandler
}

type Application struct {
	Queries  Queries
	Commands Commands
}

func NewApplication(cragRepo crag.Repository, ns notification.Service) Application {
	// init base
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	return Application{
		Queries: Queries{
			GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(cragRepo),
			GetCragHandler:     queries.NewGetCragRequestHandler(cragRepo),
		},
		Commands: Commands{
			CreateCragHandler: commands.NewAddCragRequestHandler(up, tp, cragRepo, ns),
			UpdateCragHandler: commands.NewUpdateCragRequestHandler(cragRepo),
			DeleteCragHandler: commands.NewDeleteCragRequestHandler(cragRepo),
		},
	}
}
