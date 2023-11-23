package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/handler"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", handler.GetUsers)
	user.Get("/:id", handler.GetUser)
	user.Delete("/:id", handler.DeleteUser)
	user.Put("/:id", handler.UpdateUser)
}
