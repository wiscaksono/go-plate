package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/model"
)

func JSON(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := model.Response{
		Ok:      statusCode <= 400,
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}
