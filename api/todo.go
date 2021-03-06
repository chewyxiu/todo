package api

import (
	"github.com/todo/dto"
	"github.com/todo/logic"
	"github.com/todo/types/apiFunc"
	"net/http"
)

var CreateTodo apiFunc.APIFunc = func(validatedRequest interface{})(statusCode int, output interface{}) {
	req := validatedRequest.(*dto.CreateTodoRequest)
	todoDTO := &dto.Todo{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		IsPrivate:  req.Private,
	}
	newTodo, err := logic.CreateTodo(todoDTO, req.UserID)
	if err != nil {
		e, ok := err.(logic.APIError)
		if !ok {
			return http.StatusInternalServerError, nil
		}
		return e.StatusCode(), e.Error()
	}

	return http.StatusCreated, newTodo
}

var UpdateTodo apiFunc.APIFunc = func(validatedRequest interface{})(statusCode int, output interface{}) {
	req := validatedRequest.(*dto.UpdateTodoRequest)
	todoDTO := &dto.Todo{
		ID:          req.TodoID,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		IsPrivate:   req.Private,
	}
	updatedTodo, err := logic.UpdateTodo(todoDTO)
	if err != nil {
		e, ok := err.(logic.APIError)
		if !ok {
			return http.StatusInternalServerError, nil
		}
		return e.StatusCode(), e.Error()
	}

	return http.StatusOK, updatedTodo
}

var GetTodos apiFunc.APIFunc = func(validatedRequest interface{})(statusCode int, output interface{}) {
	req := validatedRequest.(*dto.GetTodosRequest)
	todos, err := logic.GetTodos(req.Status, req.Priority, req.From, req.Limit, req.Offset, req.CurrUserID, req.UserID)
	if err != nil {
		e, ok := err.(logic.APIError)
		if !ok {
			return http.StatusInternalServerError, nil
		}
		return e.StatusCode(), e.Error()
	}

	return http.StatusOK, todos
}