package api

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/app"
	"go-clean-architecture-example/internal/common/errors"
	"go-clean-architecture-example/internal/common/responses"
	"go-clean-architecture-example/internal/common/validator"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
)

type CragHttpApi interface {
	AddCrag(ctx *fiber.Ctx) error
	GetCrags(ctx *fiber.Ctx) error
}

type cragHttpApi struct {
	cragApp app.Application
}

// NewHandler Constructor
func NewCragHttpApi(cragApp app.Application) CragHttpApi {
	return &cragHttpApi{cragApp: cragApp}
}

// GetCrags GetAll Returns all available crags
// @Summary Get all crags
// @Tags Crag
// @Produce json
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Router /crag [get]
func (cr *cragHttpApi) GetCrags(ctx *fiber.Ctx) error {
	crag, err := cr.cragApp.Queries.GetAllCragsHandler.Handle(ctx.Context(), dto.GetAllCragRequest{})
	if err != nil {
		return err
	}
	resp := responses.DefaultSuccessResponse
	resp.Data = crag
	return resp.JSON(ctx)
}

// AddCrag Add a new crag
// @Summary Add a new crag
// @Tags Crag
// @Accept json
// @Produce json
// @Param crag body dto.AddCragRequest true "The crag data"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Router /crag [post]
func (cr *cragHttpApi) AddCrag(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := new(dto.AddCragRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return errors.ErrBadRequest
	}

	if err := validator.GetValidator().Validate(req); err != nil {
		return errors.ErrBadRequest
	}

	if err := cr.cragApp.Commands.AddCragHandler.Handle(context, req); err != nil {
		return err
	}
	return responses.DefaultSuccessResponse.JSON(ctx)
}
