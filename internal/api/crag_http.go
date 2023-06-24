package api

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/app"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
)

type CragHttpApi interface {
	GetCrags(ctx *fiber.Ctx) error
}

type cragHttpApi struct {
	cragApp app.Application
}

// NewHandler Constructor
func NewCragHttpApi(cragApp app.Application) CragHttpApi {
	return &cragHttpApi{cragApp: cragApp}
}

// GetAll Returns all available crags
func (cr *cragHttpApi) GetCrags(ctx *fiber.Ctx) error {
	crag, err := cr.cragApp.Queries.GetAllCragsHandler.Handle(ctx.Context(), dto.GetAllCragRequest{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(nil)
	}
	return ctx.Status(fiber.StatusCreated).JSON(crag)
}
