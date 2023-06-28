package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"testing"
	"time"
)

func TestUpdateCragCommandHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo crag.Repository
	}
	type args struct {
		command *dto.UpdateCragRequest
		ctx     context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name: "happy path - no errors - should return nil",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					returnedCrag := crag.Crag{
						ID:        mockUUID,
						Name:      "initial",
						Desc:      "initial",
						Country:   "initial",
						CreatedAt: time.Time{},
					}
					updatedCrag := crag.Crag{
						ID:        mockUUID,
						Name:      "updated",
						Desc:      "updated",
						Country:   "updated",
						CreatedAt: time.Time{},
					}
					mp.On("GetByID", mockUUID).Return(&returnedCrag, nil)
					mp.On("Update", updatedCrag).Return(nil)

					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateCragRequest{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
				ctx: context.Background(),
			},
			err: nil,
		},
		{
			name: "get error should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return(&crag.Crag{ID: mockUUID}, errors.New("get error"))

					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateCragRequest{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
				ctx: context.Background(),
			},
			err: errors.New("get error"),
		},
		{
			name: "get returns nil, should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return((*crag.Crag)(nil), nil)
					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateCragRequest{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
				ctx: context.Background(),
			},
			err: fmt.Errorf("the provided crag id does not exist"),
		},
		{
			name: "update error - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					returnedCrag := crag.Crag{
						ID:        mockUUID,
						Name:      "initial",
						Desc:      "initial",
						Country:   "initial",
						CreatedAt: time.Time{},
					}
					updatedCrag := crag.Crag{
						ID:        mockUUID,
						Name:      "updated",
						Desc:      "updated",
						Country:   "updated",
						CreatedAt: time.Time{},
					}
					mp.On("GetByID", mockUUID).Return(&returnedCrag, nil)
					mp.On("Update", updatedCrag).Return(errors.New("update error"))

					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateCragRequest{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
				ctx: context.Background(),
			},
			err: errors.New("update error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := updateCragRequestHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.ctx, tt.args.command)
			assert.Equal(t, tt.err, err)
		})
	}
}
