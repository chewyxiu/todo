package storage

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Todos           UserTodos
}
type IUserDAO interface {
	Get(id uint) (*User, error)
}

var UserDAO IUserDAO = &userDAOImpl{}

type userDAOImpl struct{}

func (*userDAOImpl) Get(id uint) (*User, error) {
	var res []*User
	if err := DB.Where("id = ?", id).Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res[0], nil
}
