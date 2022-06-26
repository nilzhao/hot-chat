package model

import (
	"github.com/shopspring/decimal"
)

type Account struct {
	BaseModel
	AccountNo    string          `json:"accountNo" db:"account_no,uni"`   //账户编号,账户唯一标识
	AccountName  string          `json:"accountName" db:"account_name"`   //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  int             `json:"accountType" db:"account_type"`   //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string          `json:"currencyCode" db:"currency_code"` //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	UserId       string          `json:"userId" db:"user_id"`             //用户编号, 账户所属用户
	Username     string          `json:"username" db:"username"`          //用户名称
	Balance      decimal.Decimal `json:"balance" db:"balance"`            //账户可用余额
	Status       AccountStatus   `json:"status" db:"status"`              //账户状态，账户状态：0账户初始化，1启用，2停用
}

// 账户状态
type AccountStatus int

const (
	AccountStatusInit AccountStatus = iota
	AccountStatusEnabled
	AccountStatusDisabled
)
