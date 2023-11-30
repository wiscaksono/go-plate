package dto

import "github.com/wiscaksono/go-plate/internal/app/model"

type UserResponse struct {
	model.Base
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Role     model.UserRole `json:"role"`
}

func CreateUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		Base:     user.Base,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}
}

func CreateUsersResponse(users []model.User) []*UserResponse {
	var usersResponse []*UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, CreateUserResponse(&user))
	}

	return usersResponse
}
