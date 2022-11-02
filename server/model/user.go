package model

import "gorm.io/gorm"

type UserGender int

const (
	GENDER_UNKNOWN UserGender = iota
	GENDER_MALE
	GENDER_FEMALE
)

type UserStatus int

const (
	USER_STATUS_NORMAL UserStatus = iota
	USER_STATUS_DISABLED
)

type User struct {
	BaseModel
	Name     string         `json:"name"`
	NickName string         `json:"nickName"`
	Password string         `json:"-"`
	Avatar   string         `json:"avatar"`
	Email    string         `json:"email" gorm:"uniqueIndex"`
	Gender   UserGender     `json:"gender"`
	Phone    string         `json:"phone"`
	Birthday gorm.DeletedAt `json:"birthday"`
	Status   UserStatus     `json:"status"`
}

type UserDto struct {
	Name     string         `json:"name"`
	NickName string         `json:"nickName"`
	Password string         `json:"password" validate:"required,gte=6"`
	Avatar   string         `json:"avatar"`
	Email    string         `json:"email" validate:"required"`
	Gender   UserGender     `json:"gender"`
	Phone    string         `json:"phone"`
	Birthday gorm.DeletedAt `json:"birthday"`
}

type Users []User
