package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wiscaksono/go-plate/internal/app/dto"
	"github.com/wiscaksono/go-plate/internal/app/model"
	"github.com/wiscaksono/go-plate/internal/app/repository"
	"github.com/wiscaksono/go-plate/internal/app/util"
)

func CreateTodos(c *fiber.Ctx) error {
	payload := new(dto.TodoRequest)

	if err := c.BodyParser(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Failed to parse JSON", nil)
	}

	if err := util.ValidateStruct(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Validation error", err)
	}

	todo := model.Todo{
		Title:     payload.Title,
		Completed: payload.Completed,
		UserID:    c.Locals("userId").(uuid.UUID),
	}

	if err := repository.DB.Create(&todo).Error; err != nil {
		return JSON(c, fiber.StatusInternalServerError, "Failed creating todo", nil)
	}

	return JSON(c, fiber.StatusCreated, "Success creating todo", nil)
}

func GetTodos(c *fiber.Ctx) error {
	todos := new([]model.Todo)

	if err := repository.DB.Where("user_id = ?", c.Locals("userId")).Find(&todos).Error; err != nil {
		return JSON(c, fiber.StatusNotFound, "No todos found", nil)
	}

	return JSON(c, fiber.StatusOK, "Success getting todos data", dto.CreateTodoResponses(todos))
}

func GetTodo(c *fiber.Ctx) error {
	todo := new(model.Todo)

	if err := repository.DB.Where("id = ?", c.Params("id")).First(&todo).Error; err != nil {
		return JSON(c, fiber.StatusNotFound, "No todo found", nil)
	}

	return JSON(c, fiber.StatusOK, "Success getting todo data", dto.CreateTodoResponse(todo))
}

func UpdateTodos(c *fiber.Ctx) error {
	payload := new(dto.TodoUpdateRequest)

	if err := c.BodyParser(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Failed to parse JSON", nil)
	}

	if err := util.ValidateStruct(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Validation error", err)
	}

	todo := new(model.Todo)

	if err := repository.DB.Where("id = ?", c.Params("id")).First(&todo).Error; err != nil {
		return JSON(c, fiber.StatusNotFound, "No todo found", nil)
	}

	todo.Title = payload.Title
	todo.Completed = payload.Completed

	if err := repository.DB.Save(&todo).Error; err != nil {
		return JSON(c, fiber.StatusInternalServerError, "Failed updating todo", nil)
	}

	return JSON(c, fiber.StatusOK, "Success updating todo", nil)
}

func DeleteTodos(c *fiber.Ctx) error {
	todo := new(model.Todo)

	if c.Params("id") == "all" {
		if err := repository.DB.Where("user_id = ?", c.Locals("userId")).Delete(&todo).Error; err != nil {
			return JSON(c, fiber.StatusInternalServerError, "Failed deleting todo", nil)
		}

		return JSON(c, fiber.StatusOK, "Success deleting todo", nil)
	}

	if err := repository.DB.Where("id = ?", c.Params("id")).First(&todo).Error; err != nil {
		return JSON(c, fiber.StatusNotFound, "No todo found", nil)
	}

	if err := repository.DB.Delete(&todo).Error; err != nil {
		return JSON(c, fiber.StatusInternalServerError, "Failed deleting todo", nil)
	}

	return JSON(c, fiber.StatusOK, "Success deleting todo", nil)
}
