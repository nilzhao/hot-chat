package service

import (
	"errors"
	"red-server/model"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *Account {
	return &Account{db}
}

func (s *Account) GenerateAccountNo() string {
	return ksuid.New().Next().String()
}

func (a *Account) Create(account *model.Account) error {
	account.AccountNo = a.GenerateAccountNo()
	accountDaoService := NewAccountDaoService(a.db)
	return accountDaoService.Inert(account)
}

// 修正 amount
func fixTransferAmount(dto *model.AccountTransferDTO) decimal.Decimal {
	amount := dto.Amount
	if dto.ChangeFlag == model.FLAG_TRANSFER_OUT && dto.Amount.GreaterThan(decimal.NewFromFloat(0)) {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}
	return amount
}

func (a *Account) Transfer(dto *model.AccountTransferDTO) (status model.TransferredStatus, err error) {
	// 先转账给别人
	status, err = a.TransferTo(dto)
	if status != model.TRANSFERRED_STATUS_SUCCESS || err != nil {
		return status, err
	}
	// 转账成功后,被转账的人,接收
	return a.MayTransferBack(dto)
}

// 转账
func (a *Account) TransferTo(
	dto *model.AccountTransferDTO,
) (status model.TransferredStatus, err error) {
	accountDaoService := NewAccountDaoService(a.db)
	accountLogService := NewAccountLogService(a.db)
	statusFailure := model.TRANSFERRED_STATUS_FAILURE
	amount := fixTransferAmount(dto)
	rowsAffected, err := accountDaoService.UpdateBalance(dto.TradeBody.AccountNo, amount)
	if err != nil {
		return statusFailure, err
	}
	if rowsAffected <= 0 && dto.ChangeFlag == model.FLAG_TRANSFER_OUT {
		return model.TRANSFERRED_STATUS_SUFFICIENT_FUNDS, errors.New("余额不足")
	}
	account := accountDaoService.GetOne(dto.TradeBody.AccountNo)
	if account == nil {
		return statusFailure, errors.New("账户不存在")
	}
	accountLog := accountLogService.GenerateAccountTransferredLog(dto, status, account.Balance)
	err = accountLogService.Create(accountLog)
	if err != nil {
		return statusFailure, err
	}

	return model.TRANSFERRED_STATUS_SUCCESS, err
}

func (a *Account) MayTransferBack(
	dto *model.AccountTransferDTO,
) (model.TransferredStatus, error) {
	// 转账成功后,并且交易主体和交易目标不是同一个人,而且交易类型不是储值,则进行反向操作
	if dto.TradeBody.AccountNo != dto.TradeTarget.AccountNo &&
		dto.ChangeType != model.ACCOUNT_STORE_VALUE {
		backDto := *dto
		backDto.TradeBody = dto.TradeTarget
		backDto.TradeTarget = dto.TradeBody
		backDto.ChangeType = -dto.ChangeType
		backDto.ChangeFlag = -dto.ChangeFlag
		return a.TransferTo(&backDto)
	}
	return model.TRANSFERRED_STATUS_SUCCESS, nil
}

func (a *Account) GetByNo(accountNo string) *model.Account {
	accountDaoService := NewAccountDaoService(a.db)
	return accountDaoService.GetOne(accountNo)
}
