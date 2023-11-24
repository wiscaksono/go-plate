package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/model"
	"github.com/wiscaksono/go-plate/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not hash password",
		})
	}

	newUser := model.User{
		Email:    strings.ToLower(user.Email),
		Username: user.Username,
		Password: string(hashedPassword),
	}

	if err := repository.DB.Create(&newUser).Error; err != nil {
		return err
	}

	return JSON(c, fiber.StatusCreated, "User created successfully", fiber.Map{
		"user": newUser.ToUserResponse(),
	})
}
