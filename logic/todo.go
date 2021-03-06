package logic

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"github.com/todo/dto"
	"github.com/todo/storage"
	"github.com/todo/types/status"
	"time"
)

func CreateTodo(todo *dto.Todo, userID uint) (*dto.Todo, error) {
	dueDate, err := checkValidTime(todo.DueDate)
	if err != nil {
		return nil, err
	}

	user, err := storage.UserDAO.Get(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidUser
	}

	newTodo := &storage.Todo{
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
		DueDate:     dueDate,
		Priority:    todo.Priority,
		IsPrivate: todo.IsPrivate,
	}

	createdTodo, err := storage.TodoDAO.Create(newTodo)
	if err != nil {
		return nil, err
	}

	userTodo := &storage.UserTodo{
		UserID:    user.ID,
		TodoID:    createdTodo.ID,
	}

	userTodo, err = storage.UserTodoDAO.Create(userTodo)
	if err != nil {
		return nil, err
	}


	todoDTO := createdTodo.ToDTO()
	todoDTO.UserID = user.ID
	return todoDTO, nil
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
		IsPrivate:   todo.IsPrivate,
		IsDeleted:   todo.IsDeleted,
	}

	updatedTodo, err := storage.TodoDAO.Update(updateTodo)
	if err != nil {
		return nil, err
	}

	todoDTO := updatedTodo.ToDTO()
	return todoDTO, nil
}

func GetTodos(status status.Type, priority uint, from string, limit uint, offset uint, currentUserID uint, viewUserID uint)([]*dto.Todo, error) {
	var values []interface{}

	if from != "" {
		_, err := checkValidTime(from)
		if err != nil {
			return nil, err
		}
	}

	if currentUserID != viewUserID {
		friends, err := storage.FriendDAO.Get(viewUserID)
		if err != nil {
			return nil, err
		}

		isFriend := checkFriend(friends, currentUserID)
		if !isFriend {
			return nil, ErrNotFriend
		}
	}

	// get list of todos belonging to user
	userTodos, err := storage.UserTodoDAO.Get(viewUserID)
	if err != nil {
		return nil, err
	}

	todoMap := make(map[uint]bool)
	for _, val := range userTodos {
		todoMap[val.TodoID] = true
	}


	condition, values := buildQuery(status, priority, from)
	todos, err := storage.TodoDAO.Get(condition, values, limit, offset)
	if err != nil {
		return nil, err
	}

	res := []*dto.Todo{}
	for _, v := range todos {
		if _, ok := todoMap[v.ID]; ok {
			if currentUserID == viewUserID || !v.IsPrivate {
				res = append(res, v.ToDTO())
			}
		}
	}

	return res, nil
}

func checkValidTime (date string) (time.Time, error){
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}, ErrInvalidTimeFormat
	}
	return t, nil
}

func checkFriend(friends []uint, id  uint) bool {
	for _, f := range friends {
		if f == id {
			return true
		}
	}
	return false
}

func buildQuery(status status.Type, priority uint, from string) (string, []interface{}) {
	var str bytes.Buffer
	var values []interface{}

	if status != "" {
		str.WriteString( "status = ? AND ")
		values = append(values, status)
	}

	if priority != 0 {
		str.WriteString("priority = ? AND ")
		values = append(values, priority)
	}

	if from != "" {
		str.WriteString("due_date > ? AND ")
		values = append(values, from)
	}

	return str.String(), values
}