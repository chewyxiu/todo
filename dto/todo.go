package dto

import "github.com/todo/types/status"

type Todo struct {
	ID     uint `json:"id"`
	Name   string `json:"name"`
	Description  string `json:"description"`
	Status status.Type `json:"status"`
	Priority uint `json:"priority"`
	DueDate string `json:"dueDate"`
}

type CreateTodoRequest struct {
	Name   string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Status status.Type `json:"status" validate:"required"`
	Priority uint `json:"priority" validate:"required"`
	DueDate string `json:"dueDate" validate:"required"`
}

type UpdateTodoRequest struct {
	ID     uint `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Status status.Type `json:"status" validate:"required"`
	Priority uint `json:"priority" validate:"required"`
	DueDate string `json:"dueDate" validate:"required"`
}

type GetTodosRequest struct {
	Status status.Type `json:"status"`
	Priority uint `json:"priority"`
	From string `json:"from"`
	Limit uint `json:"limit"`
	Offset uint `json:"offset"`
}