package model

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Email    string `json:"email" gorm:"unique;not null"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
}

type UserResponse struct {
	Base
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) ToUserLogin() *UserLogin {
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	}
}

func (u *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		Base:     u.Base,
		Email:    u.Email,
		Username: u.Username,
	}
}

func (u *User) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if err := validateEmail(u.Email); err != nil {
		return err
	}

	if err := validatePassword(u.Password); err != nil {
		return err
	}

	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func validateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}
