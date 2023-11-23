package model

type User struct {
	Base
	Email    string `json:"email" gorm:"unique;not null"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
