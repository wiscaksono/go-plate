package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
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
