package service

import (
	"errors"
	"hot-chat/model"
)

func (s *AccountLogService) Insert(accountLog *model.AccountLog) error {
	result := s.db.Create(accountLog)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("账户日志插入失败")
	}
	return nil
}
