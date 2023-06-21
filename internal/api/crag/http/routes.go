package crag

import "github.com/gin-gonic/gin"

// Map crag routes
func MapCragRoutes(group *gin.RouterGroup, h *Handler) {
	group.GET("", h.GetAll())
}
