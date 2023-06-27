package persistence

import (
	"fmt"
	"github.com/google/uuid"
	"go-clean-architecture-example/internal/common/errors"
	"go-clean-architecture-example/internal/domain/entities/crag"
)

type CragMemRepository struct {
	crags map[string]crag.Crag
}

func NewCragMemRepository() crag.Repository {
	crags := make(map[string]crag.Crag)
	return &CragMemRepository{crags}
}

// GetByID Returns the crag with the provided id
func (r *CragMemRepository) GetByID(id uuid.UUID) (*crag.Crag, error) {
	data, ok := r.crags[id.String()]
	if !ok {
		return nil, errors.ErrNotFound
	}
	return &data, nil
}

// GetAll Returns all stored crags
func (r *CragMemRepository) GetAll() ([]crag.Crag, error) {
	var values []crag.Crag
	for _, value := range r.crags {
		values = append(values, value)
	}
	return values, nil
}

// Add the provided crag
func (r *CragMemRepository) Add(crag crag.Crag) error {
	r.crags[crag.ID.String()] = crag
	return nil
}

// Update the provided crag
func (r *CragMemRepository) Update(crag crag.Crag) error {
	r.crags[crag.ID.String()] = crag
	return nil
}

// Delete the crag with the provided id
func (r *CragMemRepository) Delete(id uuid.UUID) error {
	_, exists := r.crags[id.String()]
	if !exists {
		return fmt.Errorf("id %v not found", id.String())
	}
	delete(r.crags, id.String())
	return nil
}
