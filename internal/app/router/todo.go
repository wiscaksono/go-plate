package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/handler"
	"github.com/wiscaksono/go-plate/internal/app/middleware"
)

func SetupTodoRoutes(app *fiber.App) {
	todo := app.Group("/todo")
	todo.Use(middleware.Protected)
	todo.Get("/", handler.GetTodos)
	todo.Post("/", handler.CreateTodos)
}
