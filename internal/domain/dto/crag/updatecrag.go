package dto // UpdateCragRequest Update Model
import "github.com/google/uuid"

type UpdateCragRequest struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Country string    `json:"country"`
}
