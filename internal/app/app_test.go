package app

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-example/internal/app/crag/commands"
	"go-clean-architecture-example/internal/app/crag/queries"
	"go-clean-architecture-example/internal/common/metrics"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
	"testing"
)

func TestNewApp(t *testing.T) {
	mockRepo := crag.MockRepository{}
	notificationService := notification.MockNotificationService{}
	// init base
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

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
					AddCragHandler:    commands.NewAddCragRequestHandler(mockRepo, notificationService, logger, metricsClient),
					UpdateCragHandler: commands.NewUpdateCragRequestHandler(mockRepo, logger, metricsClient),
					DeleteCragHandler: commands.NewDeleteCragRequestHandler(mockRepo, logger, metricsClient),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApplication(tt.args.cragRepo, tt.args.notificationService)
			assert.Equal(t, tt.want, got)
		})
	}
}
