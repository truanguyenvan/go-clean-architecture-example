package dto // UpdateCragRequest Update Model
import "github.com/google/uuid"

type UpdateCragRequest struct {
	ID      uuid.UUID
	Name    string
	Desc    string
	Country string
}
