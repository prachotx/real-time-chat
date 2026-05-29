package response

import "github.com/gofiber/fiber/v3"

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Send(c fiber.Ctx, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(BaseResponse{
		Message: message,
		Data:    data,
	})
}
