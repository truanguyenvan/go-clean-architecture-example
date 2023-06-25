package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go-clean-architecture-example/internal/common/decorator"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
	"time"
)

type AddCragRequestHandler decorator.CommandHandler[*dto.AddCragRequest]

type addCragRequestHandler struct {
	repo                crag.Repository
	notificationService notification.Service
}

// NewAddCragRequestHandler Initializes an AddCommandHandler
func NewAddCragRequestHandler(
	repo crag.Repository,
	notificationService notification.Service,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient) AddCragRequestHandler {
	return decorator.ApplyCommandDecorators[*dto.AddCragRequest](
		addCragRequestHandler{repo: repo, notificationService: notificationService},
		logger,
		metricsClient,
	)
}

// Handle Handles the AddCragRequest
func (h addCragRequestHandler) Handle(ctx context.Context, req *dto.AddCragRequest) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	c := crag.Crag{
		ID:        id,
		Name:      req.Name,
		Desc:      req.Desc,
		Country:   req.Country,
		CreatedAt: time.Now(),
	}
	err = h.repo.Add(c)
	if err != nil {
		return err
	}
	n := notification.Notification{
		Subject: "New crag added",
		Message: "A new crag with name '" + c.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
