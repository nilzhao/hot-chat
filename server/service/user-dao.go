package service

import (
	"red-server/global"
	"red-server/model"

	"gorm.io/gorm"
)

var (
	userCreateField = []string{"name", "email", "password", "avatar", "phone"}
)

type UserDaoService struct {
	db *gorm.DB
}

func NewUserDaoService(db *gorm.DB) *UserDaoService {
	return &UserDaoService{db}
}

func (s *UserDaoService) Inert(user *model.User) error {
	result := s.db.Select(userCreateField).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserDaoService) GetOne(id int64) *model.User {
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

func (s *UserDaoService) GetByEmail(email string) *model.User {
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
