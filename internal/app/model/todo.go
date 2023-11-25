package model

import "github.com/google/uuid"

type Todo struct {
	Base
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    uuid.UUID
}
