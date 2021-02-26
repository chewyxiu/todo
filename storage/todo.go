package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/todo/dto"
	"github.com/todo/types/status"
)

type Todos []*Todo

type Todo struct {
	gorm.Model
	Name            string
	Description     string
	Status          status.Type
	CreatedBy       string
	UpdatedBy       string
}

type ITodoDAO interface {
	Create(t *Todo) (*Todo, error)
	Get() (Todos, error)
	Update(t *Todo) (*Todo, error)
}

var TodoDAO ITodoDAO = &todoDAOImpl{}

type todoDAOImpl struct{}

func (t *todoDAOImpl) Create(todo *Todo) (*Todo, error) {
	if err := DB.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoDAOImpl) Update(todo *Todo) (*Todo, error) {
	if err := DB.Save(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoDAOImpl) Get() (Todos, error) {
	var res []*Todo
	if err := DB.Where("status <> ?", "deleted").Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res, nil
}


func (t *Todo) ToDTO() *dto.Todo {
	return &dto.Todo{
		ID:        	t.ID,
		Name: 		t.Name,
		Description: t.Description,
		Status:   t.Status,
	}
}

func (t Todos) ToDTO() []*dto.Todo {
	todos := make([]*dto.Todo, len(t))
	for index, todo := range t {
		todos[index] = todo.ToDTO()
	}
	return todos
}

