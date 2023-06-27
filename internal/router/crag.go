package router

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/api"
)

type CragRouter interface {
	Init(root *fiber.Router)
}

type cragRouter struct {
	api api.CragHttpApi
}

func NewCragRouter(api api.CragHttpApi) CragRouter {
	return &cragRouter{api: api}
}

func (mr *cragRouter) Init(root *fiber.Router) {
	cragRouter := (*root).Group("/crag")
	{
		// commands

		cragRouter.Post("", mr.api.AddCrag)
		cragRouter.Put("/:id", mr.api.UpdateCrag)
		cragRouter.Delete("/:id", mr.api.DeleteCrag)
		// queries
		cragRouter.Get("", mr.api.GetCrags)
		cragRouter.Get("/:id", mr.api.GetCrag)
	}

}
