package dto

import "github.com/wiscaksono/go-plate/internal/app/model"

type TodoResponse struct {
	model.Base
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed" validate:"required"`
}

func CreateTodoResponse(todo *model.Todo) *TodoResponse {
	return &TodoResponse{
		Base:      todo.Base,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
}

func CreateTodoResponses(todos *[]model.Todo) *[]TodoResponse {
	var todoResponses []TodoResponse

	for _, todo := range *todos {
		todoResponses = append(todoResponses, *CreateTodoResponse(&todo))
	}

	return &todoResponses
}

type TodoRequest struct {
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed" validate:"required"`
}

func CreateTodoRequest(todo *model.Todo) *TodoRequest {
	return &TodoRequest{
		Title:     todo.Title,
		Completed: todo.Completed,
	}
}

type TodoUpdateRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
