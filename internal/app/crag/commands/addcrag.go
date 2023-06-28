package commands

import (
	"context"
	"go-clean-architecture-example/internal/common/decorator"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
	"go-clean-architecture-example/pkg/logger"
	timePkg "go-clean-architecture-example/pkg/time"
	uuidPkg "go-clean-architecture-example/pkg/uuid"
)

type AddCragRequestHandler decorator.CommandHandler[*dto.AddCragRequest]

type addCragRequestHandler struct {
	uuidProvider        uuidPkg.Provider
	timeProvider        timePkg.Provider
	repo                crag.Repository
	notificationService notification.Service
}

// NewAddCragRequestHandler Initializes an AddCommandHandler
func NewAddCragRequestHandler(
	uuidProvider uuidPkg.Provider,
	timeProvider timePkg.Provider,
	repo crag.Repository,
	notificationService notification.Service,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) AddCragRequestHandler {
	return decorator.ApplyCommandDecorators[*dto.AddCragRequest](
		addCragRequestHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService},
		logger,
		metricsClient,
	)
}

// Handle Handles the AddCragRequest
func (h addCragRequestHandler) Handle(ctx context.Context, req *dto.AddCragRequest) error {
	c := crag.Crag{
		ID:        h.uuidProvider.NewUUID(),
		Name:      req.Name,
		Desc:      req.Desc,
		Country:   req.Country,
		CreatedAt: h.timeProvider.Now(),
	}
	err := h.repo.Add(c)
	if err != nil {
		return err
	}
	n := notification.Notification{
		Subject: "New crag added",
		Message: "A new crag with name '" + c.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
