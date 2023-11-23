package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/handler"
)

func SetupAuthRoutes(app *fiber.App) {
	user := app.Group("/auth")
	user.Post("/register", handler.Register)
}
