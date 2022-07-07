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
	userDaoService := NewUserDaoService(s.db)
	User := &model.User{
		Password: password,
		Name:     userDto.Name,
		NickName: userDto.NickName,
		Email:    userDto.Email,
		Phone:    userDto.Phone,
		Gender:   userDto.Gender,
		Avatar:   userDto.Avatar,
		Birthday: userDto.Birthday,
	}
	return User, userDaoService.Inert(User)
}

func (s *UserService) GetOne(id int64) *model.User {
	userDaoService := NewUserDaoService(s.db)
	user := userDaoService.GetOne(id)
	return user
}

func (s *UserService) GetByEmail(email string) *model.User {
	userDaoService := NewUserDaoService(s.db)
	user := userDaoService.GetByEmail(email)
	return user
}
