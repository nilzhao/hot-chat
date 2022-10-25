package service

import (
	"errors"
	"red-server/global"
	"red-server/model"

	"github.com/shopspring/decimal"
)

func (s *AccountService) GetOne(accountNo string) *model.Account {
	account := &model.Account{}
	result := s.db.Where("account_no = ?", accountNo).Find(account)
	if result.RowsAffected == 0 || result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	return account
}

func (s *AccountService) GetByUserId(userId int64, accountType model.AccountType) *model.Account {
	account := &model.Account{}
	result := s.db.Where("user_id = ? and type = ?", userId, accountType).Find(account)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	return account
}

func (s *AccountService) GetByUserIdAndType(
	userId string,
	accountType model.AccountType,
) *model.Account {
	account := &model.Account{}
	result := s.db.Where("user_id = ? and type = ?", userId, accountType).First(account)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	return account
}

func (s *AccountService) Insert(account *model.Account) error {
	result := s.db.Create(account)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("账户已存在")
	}
	return nil
}

// 账户余额的更新
// amount 如果是负数，就是扣减；如果是正数，就是增加
func (s *AccountService) UpdateBalance(
	accountNo string,
	amount decimal.Decimal,
) (int64, error) {
	sql := `
		update account
		 set balance=balance+CAST(? AS DECIMAL(30,6))
		 where account_no=?
		 and balance>=-1*CAST(? AS DECIMAL(30,6))
	`
	amountStr := amount.String()
	result := s.db.Exec(
		sql,
		amountStr,
		accountNo,
		amountStr,
	)
	return result.RowsAffected, result.Error
}

func (s *AccountService) UpdateStatus(accountNo string, status model.AccountStatus) (int64, error) {
	result := s.db.Model(model.Account{}).Where("accountNo = ?", accountNo).Update("status", status)
	return result.RowsAffected, result.Error
}
