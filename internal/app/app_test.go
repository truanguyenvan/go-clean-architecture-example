package app

import (
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-example/internal/app/crag/commands"
	"go-clean-architecture-example/internal/app/crag/queries"
	"go-clean-architecture-example/internal/app/notification"
	"go-clean-architecture-example/internal/domain/crag"
	"go-clean-architecture-example/internal/pkg/time"
	"go-clean-architecture-example/internal/pkg/uuid"
	"testing"
)

func TestNewApp(t *testing.T) {
	mockRepo := crag.MockRepository{}
	UUIDProvider := &uuid.MockProvider{}
	timeProvider := &time.MockProvider{}
	notificationService := notification.MockNotificationService{}

	type args struct {
		up                  uuid.Provider
		tp                  time.Provider
		cragRepo            crag.Repository
		notificationService notification.Service
	}
	tests := []struct {
		name string
		args args
		want CragServices
	}{
		{
			name: "should initialize application layer",
			args: args{
				cragRepo:            mockRepo,
				notificationService: notificationService,
				up:                  UUIDProvider,
				tp:                  timeProvider,
			},
			want: CragServices{
				Queries: Queries{
					GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(mockRepo),
					GetCragHandler:     queries.NewGetCragRequestHandler(mockRepo),
				},
				Commands: Commands{
					CreateCragHandler: commands.NewAddCragRequestHandler(UUIDProvider, timeProvider, mockRepo, notificationService),
					UpdateCragHandler: commands.NewUpdateCragRequestHandler(mockRepo),
					DeleteCragHandler: commands.NewDeleteCragRequestHandler(mockRepo),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewServices(tt.args.cragRepo, tt.args.notificationService, tt.args.up, tt.args.tp)
			assert.Equal(t, tt.want, got.CragServices)
		})
	}
}