package queries

import (
	"context"
	"errors"
	dto "go-clean-architecture-example/internal/domain/dto/crag"
	"go-clean-architecture-example/internal/domain/entities/crag"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCragsQueryHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo crag.Repository
	}
	type args struct {
		command dto.GetAllCragRequest
		ctx     context.Context
	}
	tests := []struct {
		name   string
		fields fields
		want   []dto.GetAllCragsResult
		err    error
		args   args
	}{
		{
			name: "happy path - no crag with no errors - should return crag",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetAll").Return([]crag.Crag{}, nil)
					return mp
				}(),
			},
			want: []dto.GetAllCragsResult(nil),
			err:  nil,
		},
		{
			name: "happy path - 1 crag with no errors - should return crag",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetAll").Return([]crag.Crag{{ID: mockUUID}}, nil)
					return mp
				}(),
			},
			want: []dto.GetAllCragsResult{{ID: mockUUID}},
			err:  nil,
			args: args{
				command: dto.GetAllCragRequest{},
				ctx:     context.Background(),
			},
		},
		{
			name: "get crags errors - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetAll").Return([]crag.Crag{{ID: mockUUID}}, errors.New("get crags error"))
					return mp
				}(),
			},
			want: []dto.GetAllCragsResult(nil),
			err:  errors.New("get crags error"),
			args: args{
				command: dto.GetAllCragRequest{},
				ctx:     context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getAllCragsRequestHandler{
				repo: tt.fields.repo,
			}
			got, err := h.Handle(tt.args.ctx, tt.args.command)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
