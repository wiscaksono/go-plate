package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wiscaksono/go-plate/internal/app/model"
)

func GenerateToken(user *model.User) string {
	claim := &jwt.MapClaims{
		"user": &model.User{
			Base:     user.Base,
			Username: user.Username,
			Email:    user.Email,
		},
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	string, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		panic(err)
	}

	return string
}

func Protected(c *fiber.Ctx) error {
	var tokenString string
	auth := c.Get("Authorization")

	if strings.HasPrefix(auth, "Bearer") {
		tokenString = strings.TrimPrefix(auth, "Bearer ")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
