package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/dto"
	"github.com/wiscaksono/go-plate/internal/app/model"
	"github.com/wiscaksono/go-plate/internal/app/repository"
)

func GetUser(c *fiber.Ctx) error {
	user := new(model.User)

	if count := repository.DB.Where("id = ?", c.Params("id")).Take(&user).RowsAffected; count == 0 {
		return JSON(c, fiber.StatusInternalServerError, "User not found", nil)
	}

	return JSON(c, fiber.StatusOK, "Success getting user data", fiber.Map{
		"user": dto.CreateUserResponse(user),
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []model.User

	if count := repository.DB.Find(&users).RowsAffected; count == 0 {
		return JSON(c, fiber.StatusInternalServerError, "Users not found", fiber.Map{
			"users": []model.User{},
		})
	}

	return JSON(c, fiber.StatusOK, "Success getting users data", fiber.Map{
		"users": dto.CreateUsersResponse(users),
	})
}

func DeleteUser(c *fiber.Ctx) error {
	user := new(model.User)

	if count := repository.DB.Where("id = ?", c.Params("id")).Delete(&user).RowsAffected; count == 0 {
		return JSON(c, fiber.StatusInternalServerError, "User not found", nil)
	}

	return JSON(c, fiber.StatusOK, "User deleted", nil)
}
