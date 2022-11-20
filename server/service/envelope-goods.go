package service

import (
	"errors"
	"fmt"
	"hot-chat/global"
	"hot-chat/model"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EnvelopeGoodsService struct {
	db *gorm.DB
}

func NewEnvelopeGoodsService(db *gorm.DB) *EnvelopeGoodsService {
	return &EnvelopeGoodsService{db}
}

func createEnvelopeNo() string {
	return ksuid.New().Next().String()
}

func (s *EnvelopeGoodsService) Create(goods *model.EnvelopeGoods) error {
	goods = s.GenerateCreatingGoods(goods)
	return s.Insert(goods)
}

// 生成红包订单
func (s *EnvelopeGoodsService) GenerateCreatingGoods(goods *model.EnvelopeGoods) *model.EnvelopeGoods {
	goods.RemainQuantity = goods.Quantity
	// 普通红包,计算出红包金额
	if goods.Type == model.ENVELOPE_TYPE_GENERAL {
		goods.Amount = goods.AmountOne.Mul(
			decimal.NewFromFloat(float64(goods.Quantity)))
	} else if goods.Type == model.ENVELOPE_TYPE_LUCKY {
		goods.AmountOne = decimal.NewFromFloat(0)
	}
	// 初始化剩余金额
	goods.RemainAmount = goods.Amount
	// 过期时间
	goods.ExpiredAt = time.Now().Add(24 * time.Hour)
	goods.Status = model.ENVELOPE_STATUS_CREATE
	goods.EnvelopeNo = createEnvelopeNo()
	return goods
}

// 发红包
func (s *EnvelopeGoodsService) SendOut(goods *model.EnvelopeGoods) (*model.EnvelopeGoods, error) {
	err := s.Create(goods)
	if err != nil {
		return nil, err
	}
	// 红包金额支付
	// 1. 需要红包中间商的红包资金账户，定义在配置文件中，事先初始化到资金账户表中
	// 2. 从红包发送人的资金账户中扣减红包金额 ，把红包金额从红包发送人的资金账户里扣除
	accountService := NewAccountService(s.db)

	body := accountService.GetOne(goods.AccountNo)
	if body == nil {
		return nil, fmt.Errorf("账户 %s 不存在", goods.AccountNo)
	}
	systemAccountNo := global.Config.System.Account.AccountNo
	target := accountService.GetOne(systemAccountNo)
	if target == nil {
		global.Logger.Error(fmt.Errorf("系统账户 %s 不存在", systemAccountNo))
		return nil, errors.New("内部错误")
	}
	dto := &model.AccountTransferDTO{
		TradeBody:   *body,
		TradeTarget: *target,
		Amount:      goods.Amount,
		ChangeType:  model.ENVELOPE_OUTGOING,
		ChangeFlag:  model.FLAG_TRANSFER_OUT,
		Desc:        "红包支出",
	}
	status, err := accountService.Transfer(dto)
	if status != model.TRANSFERRED_STATUS_SUCCESS || err != nil {
		return nil, err
	}
	return goods, nil
}
