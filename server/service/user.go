package service

import (
	"red-server/model"
	"red-server/utils"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) Create(userDto *model.UserCreateDto) (*model.User, error) {
	password := utils.BcryptHash(userDto.Password)
	user := &model.User{
		Password: password,
		Name:     userDto.Name,
		NickName: userDto.NickName,
		Email:    userDto.Email,
		Phone:    userDto.Phone,
		Gender:   userDto.Gender,
		Avatar:   userDto.Avatar,
		Birthday: userDto.Birthday,
	}
	return user, s.Insert(user)
}
