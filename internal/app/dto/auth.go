package dto

import "github.com/wiscaksono/go-plate/internal/app/model"

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func CreateAuthLoginRequest(email, password string) *AuthLoginRequest {
	return &AuthLoginRequest{
		Email:    email,
		Password: password,
	}
}

type AuthLoginResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

func CreateAuthLoginResponse(token string, user *model.User) *AuthLoginResponse {
	return &AuthLoginResponse{
		Token: token,
		User:  user,
	}
}

type AuthRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func CreateAuthRegisterRequest(username, email, password string) *AuthRegisterRequest {
	return &AuthRegisterRequest{
		Username: username,
		Email:    email,
		Password: password,
	}
}

type AuthRegisterResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

func CreateAuthRegisterResponse(token string, user *model.User) *AuthRegisterResponse {
	return &AuthRegisterResponse{
		Token: token,
		User:  user,
	}
}
