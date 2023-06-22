package exception

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/internal/domain/commom"
	"go-clean-architecture-example/internal/domain/enum"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {

	httpCode := fiber.StatusInternalServerError
	msg := commom.GeneralResponse{
		Code:    enum.UNKNOWN,
		Message: "Internal Server Error",
	}
	//trieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		httpCode = e.Code
		msg.Message = e.Message
		// TODO: handle code by httpcode
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(httpCode).JSON(msg)
}
