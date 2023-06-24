package queries

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-clean-architecture-example/internal/commom/decorator"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
)

// GetAllCragsRequestHandler Contains the dependencies of the Handler
type GetAllCragsRequestHandler decorator.QueryHandler[dto.GetAllCragRequest, []dto.GetAllCragsResult]

type getAllCragsRequestHandler struct {
	repo crag.Repository
}

// NewGetAllCragsRequestHandler Handler constructor
func NewGetAllCragsRequestHandler(repo crag.Repository, logger *logrus.Entry,
	metricsClient decorator.MetricsClient) GetAllCragsRequestHandler {
	return decorator.ApplyQueryDecorators[dto.GetAllCragRequest, []dto.GetAllCragsResult](
		getAllCragsRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handles the query
func (h getAllCragsRequestHandler) Handle(ctx context.Context, _ dto.GetAllCragRequest) ([]dto.GetAllCragsResult, error) {
	defer ctx.Done()
	res, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dto.GetAllCragsResult
	for _, crag := range res {
		result = append(result, dto.GetAllCragsResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt})
	}
	return result, nil
}
