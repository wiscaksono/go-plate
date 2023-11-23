package handler

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Create User",
	})
}
