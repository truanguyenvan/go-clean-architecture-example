package fibercasbin

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/common/jwt"
	"strings"
)

const (
	// RoleKey default
	RoleKey = "role"
	// RoleAnonymous anonymous
	RoleAnonymous = "anonymous"
	// AuthorType Bear Authority
	AuthorType = "Bearer"
)

type RoleAdapter struct {
	Secret []byte
}

// parseRoles interface to string
func parseRoles(role interface{}) string {
	r, ok := role.(string)
	if !ok {
		return ""
	}
	return r
}

func (r *RoleAdapter) GetRoleByToken(reqToken string) (string, error) {
	t, err := jwt.GetValue(reqToken, RoleKey, r.Secret)
	return parseRoles(t), err
}

// GetRole gets the roles name from the request.
func (r *RoleAdapter) GetRole(c *fiber.Ctx) string {
	token := c.Get(fiber.HeaderAuthorization)
	authorization := strings.Split(token, AuthorType)
	if len(authorization) == 2 {
		role, err := r.GetRoleByToken(strings.TrimSpace(authorization[1]))
		if err == nil {
			return role
		}
	}
	return ""
}

// NewRoleAdapter create adapter
func NewRoleAdapter(secret string) *RoleAdapter {
	return &RoleAdapter{
		Secret: []byte(secret),
	}
}
