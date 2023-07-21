package errors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/common/responses"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	msg := responses.DefaultErrorResponse

	// trieve the custom status code if it's an fiber.*Error
	var e *fiber.Error
	if errors.As(err, &e) {
		msg.Status = e.Code
		msg.Code = e.Code
		msg.Message = e.Message
		// TODO: handle code by httpcode
	}
	var customErr *Error
	if errors.As(err, &customErr) {
		msg = responses.BindingGeneral(customErr)
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return msg.JSON(ctx)
}
