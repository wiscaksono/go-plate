package model

type User struct {
	Base
	Email    string   `json:"email" gorm:"unique;not null"`
	Username string   `json:"username" gorm:"unique;not null"`
	Password string   `json:"password" gorm:"not null"`
	Role     UserRole `json:"role" gorm:"type:user_role;not null;default:USER"`
	Todos    []Todo   `json:"-"`
}

type UserRole string

const (
	ADMIN UserRole = "ADMIN"
	USER  UserRole = "USER"
)
