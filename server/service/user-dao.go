package service

import (
	"hot-chat/global"
	"hot-chat/model"
)

var (
	userCreateField = []string{"name", "email", "password", "avatar", "phone"}
)

func (s *UserService) Insert(user *model.User) error {
	result := s.db.Select(userCreateField).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) GetOne(id int64) *model.User {
	user := &model.User{}
	result := s.db.First(user, id)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return user
}

func (s *UserService) GetByEmail(email string) *model.User {
	user := &model.User{}
	result := s.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return user
}
