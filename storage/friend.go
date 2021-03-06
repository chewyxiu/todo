package storage

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Friend struct {
	gorm.Model
	UserID      uint
	UserFriends string
}

type IFriendDAO interface {
	Get(id uint) ([]uint, error)
}

var FriendDAO IFriendDAO = &friendDAOImpl{}

type friendDAOImpl struct{}


func (t *friendDAOImpl) Get(id uint) ([]uint, error) {
	var res []*Friend
	if err := DB.Where("id = ?", id).Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res[0].ToList(), nil
}


func (f *Friend) ToList() []uint {
	var list []uint
	_ = json.Unmarshal([]byte(f.UserFriends), &list)
	return list
}

