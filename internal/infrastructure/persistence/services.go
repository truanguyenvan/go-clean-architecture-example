package persistence

import (
	"go-clean-architecture-example/internal/app/notification"
	"go-clean-architecture-example/internal/domain/entities/crag"
	notiInfra "go-clean-architecture-example/internal/infrastructure/notification"
	cragRepo "go-clean-architecture-example/internal/infrastructure/persistence/crag/memory"
)

// Services contains the exposed services of interface adapters
type Services struct {
	NotificationService notification.Service
	CragRepository      crag.Repository
}

// NewServices Instantiates the interface adapter services
func NewServices() Services {
	return Services{
		NotificationService: notiInfra.NewNotificationService(),
		CragRepository:      cragRepo.NewRepo(),
	}
}
