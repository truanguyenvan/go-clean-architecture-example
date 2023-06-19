package interfaceadapters

import (
	"go-clean-architecture-example/internal/app/notification"
	"go-clean-architecture-example/internal/domain/crag"
	"go-clean-architecture-example/internal/infra/interfaceadapters/notification/console"
	"go-clean-architecture-example/internal/infra/interfaceadapters/storage/memory"
)

// Services contains the exposed services of interface adapters
type Services struct {
	NotificationService notification.Service
	CragRepository      crag.Repository
}

// NewServices Instantiates the interface adapter services
func NewServices() Services {
	return Services{
		NotificationService: console.NewNotificationService(),
		CragRepository:      memory.NewRepo(),
	}
}
