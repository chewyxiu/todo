package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/todo/dto"
	"github.com/todo/types/status"
	"time"
)

type Todos []*Todo

type Todo struct {
	gorm.Model
	Name            string
	Description     string
	IsDeleted       bool
	Status          status.Type
	Priority        uint
	DueDate        time.Time
	CreatedBy       string
	UpdatedBy       string
}

type ITodoDAO interface {
	Create(t *Todo) (*Todo, error)
	Get(condition string, values []interface{}, limit uint, offset uint) (Todos, error)
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

func (t *todoDAOImpl) Get(condition string, args []interface{}, limit uint, offset uint) (Todos, error) {
	var res []*Todo
	condition += "is_deleted = ?"
	args = append(args, false)

	if limit != 0 && offset != 0 {
		if err := DB.Where(condition, args...).Limit(limit).Offset(offset).Find(&res).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
	} else {
		if err := DB.Where(condition, args...).Find(&res).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
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
		Priority: t.Priority,
		DueDate:  timeValue(t.DueDate),
	}
}

func (t Todos) ToDTO() []*dto.Todo {
	todos := make([]*dto.Todo, len(t))
	for index, todo := range t {
		todos[index] = todo.ToDTO()
	}
	return todos
}

// Returns string value of time.Time
func timeValue(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().String()
}
