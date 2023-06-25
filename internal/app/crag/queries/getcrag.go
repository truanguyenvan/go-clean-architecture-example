package queries

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-clean-architecture-example/internal/common/decorator"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
)

type GetCragRequestHandler decorator.QueryHandler[dto.GetCragRequest, dto.GetCragResult]

type getCragRequestHandler struct {
	repo crag.Repository
}

// NewGetCragRequestHandler Handler Constructor
func NewGetCragRequestHandler(
	repo crag.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient) GetCragRequestHandler {
	return decorator.ApplyQueryDecorators[dto.GetCragRequest, dto.GetCragResult](
		getCragRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handlers the GetCragRequest query
func (h getCragRequestHandler) Handle(ctx context.Context, query dto.GetCragRequest) (dto.GetCragResult, error) {
	crag, err := h.repo.GetByID(query.CragID)
	var result dto.GetCragResult
	if crag != nil && err == nil {
		result = dto.GetCragResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt}
	}
	return result, err
}
