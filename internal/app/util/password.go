package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password)); err != nil {
		return err
	}
	return nil
}
