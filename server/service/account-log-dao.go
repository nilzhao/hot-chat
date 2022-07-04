package service

import (
	"errors"
	"red-server/model"

	"gorm.io/gorm"
)

type AccountLogDao struct {
	db *gorm.DB
}

func NewAccountLogDaoService(db *gorm.DB) *AccountLogDao {
	return &AccountLogDao{db}
}

func (s *AccountLogDao) Inert(accountLog *model.AccountLog) error {
	result := s.db.Create(accountLog)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("账户日志插入失败")
	}
	return nil
}
