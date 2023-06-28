package queries

import (
	"context"
	"go-clean-architecture-example/internal/common/decorator"
	"go-clean-architecture-example/internal/common/utils"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/pkg/logger"
)

// GetAllCragsRequestHandler Contains the dependencies of the Handler
type GetAllCragsRequestHandler decorator.QueryHandler[dto.GetAllCragRequest, []dto.GetAllCragsResult]

type getAllCragsRequestHandler struct {
	repo crag.Repository
}

// NewGetAllCragsRequestHandler Handler constructor
func NewGetAllCragsRequestHandler(repo crag.Repository, logger logger.Logger,
	metricsClient decorator.MetricsClient) GetAllCragsRequestHandler {
	return decorator.ApplyQueryDecorators[dto.GetAllCragRequest, []dto.GetAllCragsResult](
		getAllCragsRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handles the query
func (h getAllCragsRequestHandler) Handle(ctx context.Context, _ dto.GetAllCragRequest) ([]dto.GetAllCragsResult, error) {
	res, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dto.GetAllCragsResult
	for _, modelCrag := range res {
		var cragResult dto.GetAllCragsResult
		err = utils.BindingStruct(modelCrag, &cragResult)
		if err != nil {
			return result, err
		}
		result = append(result, cragResult)
	}
	return result, nil
}
