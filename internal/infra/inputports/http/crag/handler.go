package crag

import (
	"go-clean-architecture-example/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler Crag http request handler
type Handler struct {
	cragServices app.CragServices
}

// NewHandler Constructor
func NewHandler(app app.CragServices) *Handler {
	return &Handler{cragServices: app}
}

// GetAll Returns all available crags
func (h Handler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		crag, err := h.cragServices.Queries.GetAllCragsHandler.Handle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
		c.JSON(http.StatusCreated, crag)
	}
}
