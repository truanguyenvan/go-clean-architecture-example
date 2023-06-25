package responses

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type General struct {
	Status    int         `json:"-"`
	Code      int         `json:"code"`
	ErrorCode string      `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func (g *General) JSON(c *fiber.Ctx) error {
	return c.Status(g.Status).JSON(g)
}

func BindingGeneral(data interface{}) General {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return DefaultErrorResponse
	}
	var response General
	if err := json.Unmarshal(jsonData, &response); err != nil {
		return DefaultErrorResponse
	}
	return response
}
