package errors

import (
	"go-clean-architecture-example/internal/domain/enum"
)

type Error struct {
	Status    int    `json:"-"`
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

var (
	ErrInternalServer = &Error{
		Status:    500,
		Code:      enum.INTERNAL,
		ErrorCode: enum.COMMON_CODE,
		Message:   "Internal server error",
	}

	ErrBadRequest = &Error{
		Status:    400,
		Code:      enum.INVALID_ARGUMENT,
		ErrorCode: enum.COMMON_CODE,
		Message:   "Bad request",
	}

	ErrPermissionDenied = &Error{
		Status:    403,
		Code:      enum.PERMISSION_DENIED,
		ErrorCode: enum.COMMON_CODE,
		Message:   "Permission denied",
	}

	ErrNotFound = &Error{
		Status:    404,
		Code:      enum.NOT_FOUND,
		ErrorCode: enum.COMMON_CODE,
		Message:   "Not found",
	}

	ErrAlreadyExists = &Error{
		Status:    409,
		Code:      enum.ALREADY_EXISTS,
		ErrorCode: enum.COMMON_CODE,
		Message:   "Already exists",
	}

	ErrUnauthenticated = &Error{
		Status:    401,
		Code:      enum.UNAUTHENTICATED,
		ErrorCode: enum.COMMON_CODE,
		Message:   "Unauthorized",
	}
)
