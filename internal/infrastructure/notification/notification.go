package notification

import (
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"go-clean-architecture-example/internal/domain/entities/notification"
)

var Set = wire.NewSet(
	NewNotificationService,
)

// NotificationService provides a console implementation of the Service
type NotificationService struct{}

// NewNotificationService constructor for NotificationService
func NewNotificationService() notification.Service {
	return &NotificationService{}
}

// Notify prints out the notifications in console
func (r NotificationService) Notify(notification notification.Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	fmt.Printf("Notification Received: %v", string(jsonNotification))
	return nil
}
