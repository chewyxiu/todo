package logic

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"github.com/todo/dto"
	"github.com/todo/storage"
	"github.com/todo/types/status"
	"time"
)

func CreateTodo(todo *dto.Todo) (*dto.Todo, error) {
	dueDate, err := checkValidTime(todo.DueDate)
	if err != nil {
		return nil, err
	}
	newTodo := &storage.Todo{
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
		DueDate:     dueDate,
		Priority:    todo.Priority,
	}
	 createdTodo, err := storage.TodoDAO.Create(newTodo)
	if err != nil {
		return nil, err
	}

	return createdTodo.ToDTO(), nil
}

func UpdateTodo(todo *dto.Todo) (*dto.Todo, error) {
	dueDate, err := checkValidTime(todo.DueDate)
	if err != nil {
		return nil, err
	}

	updateTodo := &storage.Todo{
		Model:       gorm.Model{
			ID: todo.ID,
		},
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
		DueDate:     dueDate,
		Priority:    todo.Priority,
	}

	updatedTodo, err := storage.TodoDAO.Update(updateTodo)
	if err != nil {
		return nil, err
	}

	return updatedTodo.ToDTO(), nil
}

func GetTodos(status status.Type, priority uint, from string, limit uint, offset uint)([]*dto.Todo, error) {
	var values []interface{}
	var str bytes.Buffer

	if status != "" {
		str.WriteString( "status = ? AND ")
		values = append(values, status)
	}

	if priority != 0 {
		str.WriteString("priority = ? AND ")
		values = append(values, priority)
	}

	if from != "" {
		_, err := checkValidTime(from)
		if err != nil {
			return nil, err
		}
		str.WriteString("due_date > ? AND ")
		values = append(values, from)
	}

	todos, err := storage.TodoDAO.Get(str.String(), values, limit, offset)
	if err != nil {
		return nil, err
	}

	return todos.ToDTO(), nil
}

func checkValidTime (date string) (time.Time, error){
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}, ErrInvalidTimeFormat
	}
	return t, nil
}