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
		// queries
		cragRouter.Post("", mr.api.AddCrag)
		// commands
		cragRouter.Get("", mr.api.GetCrags)

	}

}
