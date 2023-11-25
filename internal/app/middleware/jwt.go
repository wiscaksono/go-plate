package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/wiscaksono/go-plate/config"
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
	string, err := token.SignedString([]byte(config.APP_SECRET))
	if err != nil {
		panic(err)
	}

	return string
}

func GetUserIDFromToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.APP_SECRET), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("invalid claims type")
	}

	userID, ok := claims["user"].(map[string]interface{})["id"].(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("user ID not found or not a string")
	}

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return parsedUUID, nil
}
