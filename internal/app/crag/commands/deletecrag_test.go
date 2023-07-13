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
)

func TestDeleteCragCommandHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo crag.Repository
	}
	type args struct {
		command *dto.DeleteCragRequest
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
					mp.On("GetByID", mockUUID).Return(&crag.Crag{ID: mockUUID}, nil)
					mp.On("Delete", mockUUID).Return(nil)
					return mp
				}(),
			},
			args: args{
				command: &dto.DeleteCragRequest{
					CragID: mockUUID,
				},
				ctx: context.Background(),
			},
			err: nil,
		},
		{
			name: "get crag returns error - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return(&crag.Crag{ID: mockUUID}, errors.New("get error"))
					return mp
				}(),
			},
			args: args{
				command: &dto.DeleteCragRequest{
					CragID: mockUUID,
				},
				ctx: context.Background(),
			},
			err: errors.New("get error"),
		},
		{
			name: "get crag returns nil - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return((*crag.Crag)(nil), nil)
					return mp
				}(),
			},
			args: args{
				command: &dto.DeleteCragRequest{
					CragID: mockUUID,
				},
				ctx: context.Background(),
			},
			err: fmt.Errorf("the provided crag id does not exist"),
		},
		{
			name: "delete crag returns error - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return(&crag.Crag{ID: mockUUID}, nil)
					mp.On("Delete", mockUUID).Return(errors.New("delete error"))
					return mp
				}(),
			},
			args: args{
				command: &dto.DeleteCragRequest{
					CragID: mockUUID,
				},
				ctx: context.Background(),
			},
			err: errors.New("delete error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := deleteCragRequestHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.ctx, tt.args.command)
			assert.Equal(t, tt.err, err)
		})
	}
}
