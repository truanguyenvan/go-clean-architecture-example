package app

import (
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-example/internal/app/crag/commands"
	"go-clean-architecture-example/internal/app/crag/queries"
	"go-clean-architecture-example/internal/common/metrics"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
	logger2 "go-clean-architecture-example/pkg/logger"
	"go-clean-architecture-example/pkg/time"
	"go-clean-architecture-example/pkg/uuid"
	"testing"
)

func TestNewApp(t *testing.T) {
	mockRepo := crag.MockRepository{}
	notificationService := notification.MockNotificationService{}
	// init base
	logger := logger2.NewApiLogger()
	metricsClient := metrics.NoOp{}
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	type args struct {
		cragRepo            crag.Repository
		notificationService notification.Service
	}
	tests := []struct {
		name string
		args args
		want Application
	}{
		{
			name: "should initialize application layer",
			args: args{
				cragRepo:            mockRepo,
				notificationService: notificationService,
			},
			want: Application{
				Queries: Queries{
					GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(mockRepo, logger, metricsClient),
					GetCragHandler:     queries.NewGetCragRequestHandler(mockRepo, logger, metricsClient),
				},
				Commands: Commands{
					AddCragHandler:    commands.NewAddCragRequestHandler(up, tp, mockRepo, notificationService, logger, metricsClient),
					UpdateCragHandler: commands.NewUpdateCragRequestHandler(mockRepo, logger, metricsClient),
					DeleteCragHandler: commands.NewDeleteCragRequestHandler(mockRepo, logger, metricsClient),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApplication(tt.args.cragRepo, tt.args.notificationService, logger)
			assert.Equal(t, tt.want, got)
		})
	}
}
