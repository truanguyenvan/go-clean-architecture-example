package fibercasbin

import (
	"github.com/casbin/casbin/v2/persist"
	fiberCasbin "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/common/errors"
)

type Config struct {
	// ModelFilePath is path to model file for Casbin.
	ModelFilePath string
	// PolicyAdapter is an interface for different persistent providers.
	PolicyAdapter persist.Adapter
	// Secret is a secret key for validation
	Secret string
}

// NewFiberCasbin create fiber casbin middleware
func NewFiberCasbin(config Config) *fiberCasbin.Middleware {
	return fiberCasbin.New(fiberCasbin.Config{
		ModelFilePath: config.ModelFilePath,
		PolicyAdapter: config.PolicyAdapter,
		Lookup:        NewRoleAdapter(config.Secret).GetRole,
		Unauthorized: func(ctx *fiber.Ctx) error {
			return errors.ErrUnauthenticated
		},
		Forbidden: func(ctx *fiber.Ctx) error {
			return errors.ErrPermissionDenied
		},
	})
}
