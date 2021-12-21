package exception

import (
	"github.com/pauluswi/bigben/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		ctx.Status(400)
		return ctx.JSON(model.WebResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
	}

	ctx.Status(200)
	return ctx.JSON(model.WebResponse{
		Code:    500,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
		Data:    nil,
	})
}
