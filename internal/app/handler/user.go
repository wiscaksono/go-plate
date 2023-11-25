package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/model"
	"github.com/wiscaksono/go-plate/internal/app/repository"
)

func GetUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := repository.DB.Take(&user).Error; err != nil {
		return JSON(c, fiber.StatusInternalServerError, "Failed to fetch users", nil)
	}

	return JSON(c, fiber.StatusOK, "Success getting user data", fiber.Map{
		"user": *user.ToUserResponse(),
	})
}

func GetUsers(c *fiber.Ctx) error {
	users := new([]model.User)

	if err := repository.DB.Find(&users).Error; err != nil {
		return JSON(c, fiber.StatusInternalServerError, "Failed to fetch users", nil)
	}

	userResponses := make([]model.UserResponse, len(*users))
	for i, user := range *users {
		userResponses[i] = *user.ToUserResponse()
	}

	return JSON(c, fiber.StatusOK, "Success getting users data", fiber.Map{
		"users": userResponses,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(model.User)

	if repository.DB.Where("id = ?", id).Take(&user).RowsAffected == 0 {
		return JSON(c, fiber.StatusNotFound, "No user found with ID", nil)
	}

	if err := repository.DB.Delete(&user).Error; err != nil {
		return JSON(c, fiber.StatusInternalServerError, "Failed to delete user", nil)
	}

	return JSON(c, fiber.StatusOK, "User deleted successfully", nil)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(model.User)
	payload := new(model.User)

	if err := c.BodyParser(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Failed to parse JSON", nil)
	}

	if repository.DB.Where("id = ?", id).Take(user).RowsAffected == 0 {
		return JSON(c, fiber.StatusNotFound, "No user found with ID", nil)
	}

	if err := payload.Validate(); err != nil {
		return JSON(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if payload.Password != "" {
		hashedPassword, err := user.HashPassword(payload.Password)
		if err != nil {
			return JSON(c, fiber.StatusInternalServerError, "Could not hash password", nil)
		}
		payload.Password = string(hashedPassword)
	}

	if err := repository.DB.Model(user).Updates(payload).Error; err != nil {
		return JSON(c, fiber.StatusInternalServerError, "Failed to update user", nil)
	}

	return JSON(c, fiber.StatusOK, "User updated successfully", fiber.Map{
		"user": *user.ToUserResponse(),
	})
}
