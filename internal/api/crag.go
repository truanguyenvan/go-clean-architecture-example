package api

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/app"
)

type CragApi interface {
	GetCrags(ctx *fiber.Ctx) error
}

type cragApi struct {
	cragApp app.Application
}

// NewHandler Constructor
func NewCragApi(cragApp app.Application) CragApi {
	return &cragApi{cragApp: cragApp}
}

// GetAll Returns all available crags
func (cr *cragApi) GetCrags(ctx *fiber.Ctx) error {
	crag, err := cr.cragApp.Queries.GetAllCragsHandler.Handle()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(nil)
	}
	return ctx.Status(fiber.StatusCreated).JSON(crag)
}
