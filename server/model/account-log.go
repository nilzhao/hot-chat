package model

import (
	"github.com/shopspring/decimal"
)

// 转账的类型：0 = 创建账户, >= 1 进账, <=- 支出
type AccountChangeType int

const (
	// 账户创建
	ACCOUNT_CREATED AccountChangeType = 0
	// 储值
	ACCOUNT_STORE_VALUE AccountChangeType = 1
	// 红包资金的支出
	ENVELOPE_OUTGOING AccountChangeType = -2
	// 红包资金的收入
	ENVELOPE_INCOMING AccountChangeType = 2
	// 红包过期退款
	ENVELOP_EXPIRED_REFUND AccountChangeType = 3
)

type AccountChangeFlag int

const (
	//创建账户=0
	FLAG_ACCOUNT_CREATED AccountChangeFlag = 0
	//支出=-1
	FLAG_TRANSFER_OUT AccountChangeFlag = -1
	//收入=1
	FLAG_TRANSFER_IN AccountChangeFlag = 1
)

// 转账状态
type TransferredStatus int

const (
	// 转账失败
	TRANSFERRED_STATUS_FAILURE TransferredStatus = -1
	// 余额不足
	TRANSFERRED_STATUS_SUFFICIENT_FUNDS TransferredStatus = 0
	// 转账成功
	TRANSFERRED_STATUS_SUCCESS TransferredStatus = 1
)

type AccountLog struct {
	LogBaseModel
	LogNo           string            `gorm:"uniqueIndex"` // 流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string            // 交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string            // 账户编号 账户ID
	UserId          uint              // 用户编号
	Username        string            // 用户名称
	TargetAccountNo string            // 账户编号 账户ID
	TargetUserId    uint              // 目标用户编号
	TargetUsername  string            // 目标用户名称
	Amount          decimal.Decimal   // 交易金额,该交易涉及的金额
	Balance         decimal.Decimal   // 交易后余额,该交易后的余额
	ChangeType      AccountChangeType // 流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      AccountChangeFlag // 交易变化标识：-1 出账 1为进账，枚举
	Status          TransferredStatus // 交易状态：
	Decs            string            // 交易描述
}

// 转账对象
type AccountTransferDTO struct {
	TradeBody   Account           `validate:"required"`         // 交易主体
	TradeTarget Account           `validate:"required"`         // 交易对象
	AmountStr   string            `validate:"required,numeric"` // 交易金额,该交易涉及的金额
	Amount      decimal.Decimal   ``                            // 交易金额,该交易涉及的金额
	ChangeType  AccountChangeType `validate:"required,numeric"` // 流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag  AccountChangeFlag `validate:"required,numeric"` // 交易变化标识：-1 出账 1为进账，枚举
	Decs        string            ``                            // 交易描述
}
