package crag

import (
	"github.com/google/uuid"
	"time"
)

// GetCragRequest Model of the Handler
type GetCragRequest struct {
	CragID uuid.UUID
}

// GetCragResult is the return model of Crag Query Handlers
type GetCragResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
