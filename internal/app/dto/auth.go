package dto

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginResponse struct {
	Token string        `json:"token"`
	User  *UserResponse `json:"user"`
}

func CreateAuthLoginResponse(token string, user *UserResponse) *AuthLoginResponse {
	return &AuthLoginResponse{
		Token: token,
		User:  user,
	}
}

type AuthRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
