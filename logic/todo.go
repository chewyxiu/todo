package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/todo/dto"
	"github.com/todo/storage"
)

func CreateTodo(todo *dto.Todo) (*dto.Todo, error) {
	newTodo := &storage.Todo{
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
	}
	 createdTodo, err := storage.TodoDAO.Create(newTodo)
	if err != nil {
		return nil, err
	}

	return createdTodo.ToDTO(), nil
}

func UpdateTodo(todo *dto.Todo) (*dto.Todo, error) {
	updateTodo := &storage.Todo{
		Model:       gorm.Model{
			ID: todo.ID,
		},
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
	}

	updatedTodo, err := storage.TodoDAO.Update(updateTodo)
	if err != nil {
		return nil, err
	}

	return updatedTodo.ToDTO(), nil
}

func GetTodos()([]*dto.Todo, error) {
	todos, err := storage.TodoDAO.Get()
	if err != nil {
		return nil, err
	}

	return todos.ToDTO(), nil
}