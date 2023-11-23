package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/handler"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", handler.GetUser)
	user.Post("/", handler.CreateUser)
}
