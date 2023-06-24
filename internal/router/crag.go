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
	memoRouter := (*root).Group("/crag")
	{
		memoRouter.Get("", mr.api.GetCrags)
	}

}
