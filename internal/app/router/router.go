package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/wiscaksono/go-plate/config"
)

func SetupRoutes(app *fiber.App) {
	SetupAuthRoutes(app)
	SetupUserRoutes(app)
	SetupTodoRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"app":     config.AppName,
			"version": config.AppVersion,
		})
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Go Plate Matrics"}))
}
