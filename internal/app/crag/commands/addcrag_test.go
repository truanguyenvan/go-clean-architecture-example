package commands

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go-clean-architecture-example/internal/common/metrics"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"

	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddCragCommandHandler_Handle(t *testing.T) {
	mockTime, _ := time.Parse("yyyy-MM-02", "2021-07-30")
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	type fields struct {
		repo                crag.Repository
		notificationService notification.Service
	}
	type args struct {
		request *dto.AddCragRequest
		ctx     context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name: "happy path - should not return error",
			fields: fields{
				repo: func() crag.MockRepository {
					acc := crag.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := crag.MockRepository{}
					mp.On("Add", acc).Return(nil)
					return mp
				}(),
				notificationService: func() notification.MockNotificationService {
					mock := notification.MockNotificationService{}
					n := notification.Notification{
						Subject: "New crag added",
						Message: "A new crag with name 'test' was added in the repository",
					}
					mock.On("Notify", n).Return(nil)
					return mock
				}(),
			},
			args: args{
				request: &dto.AddCragRequest{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
				ctx: context.Background(),
			},
			err: nil,
		},
		{
			name: "memory error - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					acc := crag.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := crag.MockRepository{}
					mp.On("Add", acc).Return(errors.New("test"))
					return mp
				}(),
				notificationService: func() notification.MockNotificationService {
					mock := notification.MockNotificationService{}
					n := notification.Notification{
						Subject: "New crag added",
						Message: "A new crag with name 'test' was added in the repository",
					}
					mock.On("Notify", n).Return(nil)
					return mock
				}(),
			},

			args: args{
				request: &dto.AddCragRequest{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
				ctx: context.Background(),
			},
			err: errors.New("test"),
		},
		{
			name: "happy path - should not return error",
			fields: fields{
				repo: func() crag.MockRepository {
					acc := crag.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := crag.MockRepository{}
					mp.On("Add", acc).Return(nil)
					return mp
				}(),
				notificationService: func() notification.MockNotificationService {
					mock := notification.MockNotificationService{}
					n := notification.Notification{
						Subject: "New crag added",
						Message: "A new crag with name 'test' was added in the repository",
					}
					mock.On("Notify", n).Return(errors.New("notification error"))
					return mock
				}(),
			},
			args: args{
				request: &dto.AddCragRequest{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
				ctx: context.Background(),
			},
			err: errors.New("notification error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := addCragRequestHandler{
				repo:                tt.fields.repo,
				notificationService: tt.fields.notificationService,
			}

			err := h.Handle(tt.args.ctx, tt.args.request)
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestNewAddCragCommandHandler(t *testing.T) {
	type args struct {
		repo                crag.Repository
		notificationService notification.Service
	}
	tests := []struct {
		name string
		args args
		want AddCragRequestHandler
	}{
		{
			name: "should create a request handler",
			args: args{
				repo:                crag.MockRepository{},
				notificationService: notification.MockNotificationService{},
			},
			want: addCragRequestHandler{
				repo:                crag.MockRepository{},
				notificationService: notification.MockNotificationService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := logrus.NewEntry(logrus.StandardLogger())
			metricsClient := metrics.NoOp{}
			got := NewAddCragRequestHandler(tt.args.repo, tt.args.notificationService, logger, metricsClient)
			assert.Equal(t, got, tt.want)
		})
	}
}
