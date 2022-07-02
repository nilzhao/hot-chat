package service

import (
	"errors"
	"red-server/global"
	"red-server/model"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDaoService(db *gorm.DB) *AccountDao {
	return &AccountDao{db}
}

func (a *AccountDao) GetOne(accountNo string) *model.Account {
	account := &model.Account{}
	result := a.db.Where("account_no = ?", accountNo).Find(account)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	return account
}

func (a *AccountDao) GetByUserIdAndType(
	userId string,
	accountType model.AccountType,
) *model.Account {
	account := &model.Account{}
	result := a.db.Where("user_id = ? and account_type = ?", userId, accountType).First(account)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	return account
}

func (a *AccountDao) Inert(account *model.Account) error {
	result := a.db.Create(account)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("账户创建失败")
	}

	return nil
}

// 账户余额的更新
// amount 如果是负数，就是扣减；如果是正数，就是增加
func (a *AccountDao) UpdateBalance(
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
	result := a.db.Exec(
		sql,
		amountStr,
		accountNo,
		amountStr,
	)
	return result.RowsAffected, result.Error
}

func (a *AccountDao) UpdateStatus(accountNo string, status model.AccountStatus) (int64, error) {
	result := a.db.Model(model.Account{}).Where("accountNo = ?", accountNo).Update("status", status)
	return result.RowsAffected, result.Error
}
