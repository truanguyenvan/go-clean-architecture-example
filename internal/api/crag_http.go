package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-clean-architecture-example/internal/app"
	"go-clean-architecture-example/internal/common/errors"
	"go-clean-architecture-example/internal/common/responses"
	"go-clean-architecture-example/internal/common/validator"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
)

type CragHttpApi interface {
	AddCrag(ctx *fiber.Ctx) error
	UpdateCrag(ctx *fiber.Ctx) error
	DeleteCrag(ctx *fiber.Ctx) error
	GetCrags(ctx *fiber.Ctx) error
	GetCrag(ctx *fiber.Ctx) error
}

type cragHttpApi struct {
	cragApp app.Application
}

// NewHandler Constructor
func NewCragHttpApi(cragApp app.Application) CragHttpApi {
	return &cragHttpApi{cragApp: cragApp}
}

// GetCrag GetById swagger documentation
// @Summary Get a crag by ID
// @Description Get a crag by ID
// @Tags Crag
// @Accept json
// @Produce json
// @Param id path string true "Crag ID"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /crag/{id} [get]
func (cr *cragHttpApi) GetCrag(ctx *fiber.Ctx) error {
	context := ctx.Context()

	_cragId := ctx.Params("id", "")
	if _cragId == "" {
		return errors.ErrBadRequest
	}

	cragId, err := uuid.Parse(_cragId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := &dto.GetCragRequest{CragID: cragId}

	crag, err := cr.cragApp.Queries.GetCragHandler.Handle(context, req)
	if err != nil {
		return err
	}

	resp := responses.DefaultSuccessResponse
	resp.Data = crag
	return resp.JSON(ctx)
}

// GetCrags GetAll Returns all available crags
// @Summary Get all crags
// @Tags Crag
// @Produce json
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /crag [get]
func (cr *cragHttpApi) GetCrags(ctx *fiber.Ctx) error {
	crags, err := cr.cragApp.Queries.GetAllCragsHandler.Handle(ctx.Context(), dto.GetAllCragRequest{})
	if err != nil {
		return err
	}
	resp := responses.DefaultSuccessResponse
	resp.Data = crags
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
// @Failure 400 {object} responses.General
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

// UpdateCrag swagger documentation
// @Summary Update a crag
// @Description Update a crag by ID
// @Tags Crag
// @Accept json
// @Produce json
// @Param id path string true "Crag ID"
// @Param request body dto.UpdateCragRequest true "UpdateCragRequest object"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /crag/{id} [put]
func (cr *cragHttpApi) UpdateCrag(ctx *fiber.Ctx) error {
	context := ctx.Context()

	_cragId := ctx.Params("id", "")
	if _cragId == "" {
		return errors.ErrBadRequest
	}

	cragId, err := uuid.Parse(_cragId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := new(dto.UpdateCragRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return errors.ErrBadRequest
	}

	if err := validator.GetValidator().Validate(req); err != nil {
		return errors.ErrBadRequest
	}
	req.ID = cragId

	if err := cr.cragApp.Commands.UpdateCragHandler.Handle(context, req); err != nil {
		return err
	}
	return responses.DefaultSuccessResponse.JSON(ctx)
}

// DeleteCrag swagger documentation
// @Summary Delete a crag
// @Description Delete a crag by ID
// @Tags Crag
// @Accept json
// @Produce json
// @Param id path string true "Crag ID"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /crag/{id} [delete]
func (cr *cragHttpApi) DeleteCrag(ctx *fiber.Ctx) error {
	context := ctx.Context()

	_cragId := ctx.Params("id", "")
	if _cragId == "" {
		return errors.ErrBadRequest
	}

	cragId, err := uuid.Parse(_cragId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := &dto.DeleteCragRequest{
		CragID: cragId,
	}

	if err := cr.cragApp.Commands.DeleteCragHandler.Handle(context, req); err != nil {
		return err
	}
	return responses.DefaultSuccessResponse.JSON(ctx)
}
