package dto

import "github.com/todo/types/status"

type Todo struct {
	ID     uint `json:"id"`
	Name   string `json:"name"`
	Description  string `json:"description"`
	Status status.Type `json:"status"`
	Priority uint `json:"priority"`
	DueDate string `json:"dueDate"`
	UserID uint `json:"userID,omitempty"`
	IsPrivate bool `json:"private"`
	IsDeleted bool `json:"deleted,omitempty"`
}

type CreateTodoRequest struct {
	UserID uint  `json:"userID" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Status status.Type `json:"status" validate:"required"`
	Priority uint `json:"priority" validate:"required"`
	DueDate string `json:"dueDate" validate:"required"`
	Private bool `json:"private"`
}

type UpdateTodoRequest struct {
	TodoID     uint `json:"todoID" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Status status.Type `json:"status" validate:"required"`
	Priority uint `json:"priority" validate:"required"`
	DueDate string `json:"dueDate" validate:"required"`
	Private bool `json:"private"`
	Delete  bool `json:"delete"`
}

type GetTodosRequest struct {
	CurrUserID uint `json:"currentUserID" validate:"required"`
	UserID uint `json:"UserID" validate:"required"`
	Status status.Type `json:"status"`
	Priority uint `json:"priority"`
	From string `json:"from"`
	Limit uint `json:"limit"`
	Offset uint `json:"offset"`
}