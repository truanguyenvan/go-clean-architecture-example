package app

import (
	"go-clean-architecture-example/internal/app/crag/commands"
	"go-clean-architecture-example/internal/app/crag/queries"
	"go-clean-architecture-example/internal/app/notification"
	"go-clean-architecture-example/internal/domain/crag"
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

// CragServices Contains the grouped queries and commands of the app layer
type CragServices struct {
	Queries  Queries
	Commands Commands
}

// NewServices Bootstraps Application Layer dependencies
func NewServices(cragRepo crag.Repository, ns notification.Service) CragServices {
	// init base
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	return CragServices{
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
