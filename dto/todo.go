package dto

import "github.com/todo/types/status"

type Todo struct {
	ID     uint `json:"id"`
	Name   string `json:"name"`
	Description  string `json:"description"`
	Status status.Type `json:"status"`
}

type CreateTodoRequest struct {
	Name   string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type UpdateTodoRequest struct {
	ID     uint `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Status status.Type `json:"status" validate:"required"`
}

type GetTodosRequest struct {}