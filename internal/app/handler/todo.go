package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/dto"
	"github.com/wiscaksono/go-plate/internal/app/model"
	"github.com/wiscaksono/go-plate/internal/app/repository"
	"github.com/wiscaksono/go-plate/internal/app/util"
)

func CreateTodos(c *fiber.Ctx) error {
	payload := new(dto.TodoRequest)
	id := c.Locals("userId")

	fmt.Println(id)
	fmt.Println(payload)

	if err := c.BodyParser(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Failed to parse JSON", nil)
	}

	if err := util.ValidateStruct(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Validation error", err)
	}

	// todoModel := model.Todo{
	// 	Title:     todo.Title,
	// 	Completed: todo.Completed,
	// 	UserID:    id.(uuid.UUID),
	// }
	//
	// if err := repository.DB.Create(&todoModel).Error; err != nil {
	// 	return JSON(c, fiber.StatusBadRequest, "Failed to create todo", err)
	// }

	return JSON(c, fiber.StatusCreated, "Success creating todo", nil)
}

func GetTodos(c *fiber.Ctx) error {
	todos := new([]dto.TodoResponse)
	id := c.Locals("userId")

	var todo model.Todo
	if count := repository.DB.Where("user_id = ?", id).Find(&todo).RowsAffected; count == 0 {
		return JSON(c, fiber.StatusNotFound, "No todos found", nil)
	}

	return JSON(c, fiber.StatusOK, "Success getting todos data", todos)
}
