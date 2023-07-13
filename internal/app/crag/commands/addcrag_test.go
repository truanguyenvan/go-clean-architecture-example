package commands

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"go-clean-architecture-example/internal/domain/entities/notification"
	timeUtil "go-clean-architecture-example/pkg/time"
	uuidUtil "go-clean-architecture-example/pkg/uuid"
	"testing"
	"time"
)

func TestAddCragCommandHandler_Handle(t *testing.T) {
	mockTime, _ := time.Parse("yyyy-MM-02", "2021-07-30")
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	type fields struct {
		uuidProvider        uuidUtil.Provider
		timeProvider        timeUtil.Provider
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
				uuidProvider: func() uuidUtil.MockProvider {
					id := mockUUID
					mp := uuidUtil.MockProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeUtil.Provider {
					mp := timeUtil.MockProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
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
			},
			err: nil,
		},
		{
			name: "memory error - should return error",
			fields: fields{
				uuidProvider: func() uuidUtil.MockProvider {
					id := mockUUID
					mp := uuidUtil.MockProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeUtil.Provider {
					mp := timeUtil.MockProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
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
			},
			err: errors.New("test"),
		},
		{
			name: "happy path - should not return error",
			fields: fields{
				uuidProvider: func() uuidUtil.MockProvider {
					id := mockUUID
					mp := uuidUtil.MockProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeUtil.Provider {
					mp := timeUtil.MockProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
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
			},
			err: errors.New("notification error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := addCragRequestHandler{
				uuidProvider:        tt.fields.uuidProvider,
				timeProvider:        tt.fields.timeProvider,
				repo:                tt.fields.repo,
				notificationService: tt.fields.notificationService,
			}

			err := h.Handle(tt.args.ctx, tt.args.request)
			assert.Equal(t, err, tt.err)

		})
	}
}
