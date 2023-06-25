package dto

import (
	"github.com/google/uuid"
	"time"
)

// GetAllCragsResult is the result of the GetAllCragsRequest Query
type GetAllCragsResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
type GetAllCragRequest struct{}
