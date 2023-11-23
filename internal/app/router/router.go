package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/config"
)

func SetupRoutes(app *fiber.App) {
	SetupAuthRoutes(app)
	SetupUserRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"app":     config.AppName,
			"version": config.AppVersion,
		})
	})
}
