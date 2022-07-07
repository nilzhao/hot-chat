package service

import (
	"red-server/model"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AccountLogService struct {
	db *gorm.DB
}

func NewAccountLogService(db *gorm.DB) *AccountLogService {
	return &AccountLogService{db}
}

func createLogNo() string {
	return ksuid.New().Next().String()
}

func createTradeNo() string {
	return ksuid.New().Next().String()
}

// 创建账户时的日志
func (s *AccountLogService) GenerateAccountCreatedLog(account *model.Account) model.AccountLog {
	logNo := createLogNo()
	accountLog := model.AccountLog{
		LogNo:           logNo,
		AccountNo:       account.AccountNo,
		UserId:          account.UserId,
		Username:        account.Username,
		TradeNo:         createTradeNo(),
		TargetAccountNo: account.AccountNo,
		TargetUserId:    account.UserId,
		TargetUsername:  account.Username,
		Amount:          account.Balance,
		Balance:         account.Balance,
		ChangeType:      model.ACCOUNT_CREATED,
		ChangeFlag:      model.FLAG_ACCOUNT_CREATED,
		Decs:            "账户创建",
	}
	return accountLog
}

func (s *AccountLogService) GenerateAccountTransferredLog(
	dto *model.AccountTransferDTO,
	status model.TransferredStatus,
	balance decimal.Decimal,
) *model.AccountLog {
	logNo := createLogNo()
	// 创建流水记录
	accountLog := &model.AccountLog{
		LogNo:           logNo,
		TradeNo:         createTradeNo(),
		AccountNo:       dto.TradeBody.AccountNo,
		UserId:          dto.TradeBody.UserId,
		Username:        dto.TradeBody.Username,
		TargetAccountNo: dto.TradeTarget.AccountNo,
		TargetUserId:    dto.TradeTarget.UserId,
		TargetUsername:  dto.TradeTarget.Username,
		Amount:          dto.Amount,
		Balance:         balance,
		ChangeType:      dto.ChangeType,
		ChangeFlag:      dto.ChangeFlag,
		Status:          status,
		Decs:            dto.Decs,
	}
	return accountLog
}

func (s *AccountLogService) Create(accountLog *model.AccountLog) error {
	accountLogDaoService := NewAccountLogDaoService(s.db)
	return accountLogDaoService.Inert(accountLog)
}
