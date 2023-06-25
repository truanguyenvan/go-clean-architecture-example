package crag

import "github.com/google/uuid"

// DeleteCragRequest Command Model
type DeleteCragRequest struct {
	CragID uuid.UUID
}
