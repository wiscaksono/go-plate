package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/middleware"
	"github.com/wiscaksono/go-plate/internal/app/model"
	"github.com/wiscaksono/go-plate/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	payload := new(model.User)
	if err := c.BodyParser(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}

	user := new(model.User)
	if repository.DB.Where("email = ?", payload.Email).Or("username = ?", payload.Username).First(&user).Error != nil {
		return JSON(c, fiber.StatusNotFound, "User not found", nil)
	}

	comparePassword := user.ComparePassword(payload.Password)
	if comparePassword != nil && bcrypt.ErrMismatchedHashAndPassword == comparePassword {
		return JSON(c, fiber.StatusUnauthorized, "Password is incorrect", nil)
	}

	return JSON(c, fiber.StatusOK, "Login success", fiber.Map{
		"user":  user.ToUserResponse(),
		"token": middleware.GenerateToken(user),
	})
}

func Register(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if count := repository.DB.Where("email = ?", user.Email).Or("username = ?", user.Username).Find(&model.User{}).RowsAffected; count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email or username already exists",
		})
	}

	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not hash password",
		})
	}

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
	user.Password = hashedPassword

	if err := repository.DB.Create(&user).Error; err != nil {
		return err
	}

	return JSON(c, fiber.StatusCreated, "User created successfully", fiber.Map{
		"user": user.ToUserResponse(),
	})
}
