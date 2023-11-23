package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wiscaksono/go-plate/config"
	"github.com/wiscaksono/go-plate/internal/app/repository"
	"github.com/wiscaksono/go-plate/internal/app/router"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	if err := repository.InitDatabase(); err != nil {
		panic(err)
	}
	router.SetupRoutes(app)
	app.Listen(config.APP_PORT)
}
