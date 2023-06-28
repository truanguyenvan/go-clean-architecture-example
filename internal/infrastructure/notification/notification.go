package notification

import (
	"encoding/json"
	"github.com/google/wire"
	"go-clean-architecture-example/internal/domain/entities/notification"
	"go-clean-architecture-example/pkg/logger"
)

var Set = wire.NewSet(
	NewNotificationService,
)

// NotificationService provides a console implementation of the Service
type NotificationService struct {
	logger logger.Logger
}

// NewNotificationService constructor for NotificationService
func NewNotificationService(logger logger.Logger) notification.Service {
	return &NotificationService{
		logger: logger,
	}
}

// Notify prints out the notifications in console
func (r NotificationService) Notify(notification notification.Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	r.logger.Infof("Notification Received: %v", string(jsonNotification))
	return nil
}
