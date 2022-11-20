package service

import (
	"errors"
	"hot-chat/model"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{db}
}

func (s *AccountService) GenerateAccountNo() string {
	return ksuid.New().Next().String()
}

func (s *AccountService) Create(account *model.Account) error {
	account.AccountNo = s.GenerateAccountNo()
	return s.Insert(account)
}

// 修正 amount
func fixTransferAmount(dto *model.AccountTransferDTO) decimal.Decimal {
	amount := dto.Amount
	if dto.ChangeFlag == model.FLAG_TRANSFER_OUT && dto.Amount.GreaterThan(decimal.NewFromFloat(0)) {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}
	return amount
}

func (s *AccountService) Transfer(
	dto *model.AccountTransferDTO,
) (status model.TransferredStatus, err error) {
	// 先转账给别人
	status, err = s.TransferTo(dto)
	if status != model.TRANSFERRED_STATUS_SUCCESS || err != nil {
		return status, err
	}
	// 转账成功后,被转账的人,接收
	return s.MayTransferBack(dto)
}

// 转账
func (s *AccountService) TransferTo(
	dto *model.AccountTransferDTO,
) (status model.TransferredStatus, err error) {
	accountLogService := NewAccountLogService(s.db)
	statusFailure := model.TRANSFERRED_STATUS_FAILURE
	amount := fixTransferAmount(dto)
	rowsAffected, err := s.UpdateBalance(dto.TradeBody.AccountNo, amount)
	if err != nil {
		return statusFailure, err
	}
	if rowsAffected <= 0 && dto.ChangeFlag == model.FLAG_TRANSFER_OUT {
		return model.TRANSFERRED_STATUS_SUFFICIENT_FUNDS, errors.New("余额不足")
	}
	account := s.GetOne(dto.TradeBody.AccountNo)
	if account == nil {
		return statusFailure, errors.New("账户不存在")
	}
	accountLog := accountLogService.GenerateAccountTransferringLog(dto, status, account.Balance)
	err = accountLogService.Insert(accountLog)
	if err != nil {
		return statusFailure, err
	}

	return model.TRANSFERRED_STATUS_SUCCESS, err
}

func (s *AccountService) MayTransferBack(
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
		return s.TransferTo(&backDto)
	}
	return model.TRANSFERRED_STATUS_SUCCESS, nil
}
