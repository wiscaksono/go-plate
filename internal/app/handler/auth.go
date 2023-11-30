package handler

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/go-plate/internal/app/dto"
	"github.com/wiscaksono/go-plate/internal/app/middleware"
	"github.com/wiscaksono/go-plate/internal/app/model"
	"github.com/wiscaksono/go-plate/internal/app/repository"
	"github.com/wiscaksono/go-plate/internal/app/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	payload := new(dto.AuthLoginRequest)

	if err := c.BodyParser(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}

	if err := util.ValidateStruct(payload); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Validation error", err)
	}

	log.Println(payload.Email)

	user := new(model.User)
	if repository.DB.Where("email = ?", payload.Email).Take(&user).Error != nil {
		return JSON(c, fiber.StatusNotFound, "User not found", nil)
	}

	comparePassword := util.ComparePassword(payload.Password)
	if comparePassword != nil && bcrypt.ErrMismatchedHashAndPassword == comparePassword {
		return JSON(c, fiber.StatusUnauthorized, "Password is incorrect", nil)
	}

	return JSON(c, fiber.StatusOK, "Login success", dto.CreateAuthLoginResponse(middleware.GenerateToken(user), dto.CreateUserResponse(user)))
}

func Register(c *fiber.Ctx) error {
	user := new(dto.AuthRegisterRequest)

	if err := c.BodyParser(user); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}

	if err := util.ValidateStruct(user); err != nil {
		return JSON(c, fiber.StatusBadRequest, "Validation error", err)
	}

	if count := repository.DB.Where("email = ?", user.Email).Or("username = ?", user.Username).Find(&model.User{}).RowsAffected; count > 0 {
		return JSON(c, fiber.StatusConflict, "User already exists", nil)
	}

	newUser := model.User{
		Username: strings.ToLower(user.Username),
		Email:    strings.ToLower(user.Email),
		Password: util.HashPassword(user.Password),
	}

	if err := repository.DB.Create(&newUser).Error; err != nil {
		return err
	}

	return JSON(c, fiber.StatusCreated, "User created", nil)
}
