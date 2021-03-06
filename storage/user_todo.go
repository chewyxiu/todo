package storage

import (
	"github.com/jinzhu/gorm"
)
type UserTodos []*UserTodo

type UserTodo struct {
	gorm.Model
    UserID      uint
	TodoID      uint
}

type IUserTodoDAO interface {
	Create(u *UserTodo) (*UserTodo, error)
	Get(id uint)([]*UserTodo, error)
	Update(u *UserTodo) (*UserTodo, error)
}

var UserTodoDAO IUserTodoDAO = &userTodoDAOImpl{}

type userTodoDAOImpl struct{}

func (*userTodoDAOImpl) Create(todo *UserTodo) (*UserTodo, error) {
	if err := DB.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (*userTodoDAOImpl) Get(id uint) ([]*UserTodo, error) {
	var res []*UserTodo
	if err := DB.Where("user_id = ?", id).Find(&res).Error; err != nil {
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

func (*userTodoDAOImpl) Update(userTodo *UserTodo) (*UserTodo, error) {
	if err := DB.Save(userTodo).Error; err != nil {
		return nil, err
	}

	return userTodo, nil
}
