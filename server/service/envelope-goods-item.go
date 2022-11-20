package service

import (
	"errors"
	"hot-chat/global"
	"hot-chat/model"
	"hot-chat/utils"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EnvelopeGoodsItemService struct {
	db *gorm.DB
}

func NewEnvelopeGoodsItemService(db *gorm.DB) *EnvelopeGoodsItemService {
	return &EnvelopeGoodsItemService{db}
}

func (s *EnvelopeGoodsItemService) createItemNo() string {
	return ksuid.New().Next().String()
}

func (s *EnvelopeGoodsItemService) Receive(dto *model.RedEnvelopeReceiveDTO) (item *model.EnvelopeGoodsItem, err error) {
	goodsService := NewEnvelopeGoodsService(global.DB)
	// 红包订单的详情
	goods := goodsService.GetOne(dto.EnvelopeNo)
	// 校验剩余数量和金额
	if goods.RemainQuantity <= 0 || goods.RemainAmount.LessThanOrEqual(decimal.NewFromFloat(0)) {
		return nil, errors.New("没有足够的红包和金额")
	}
	// 使用红包算法,算出红包金额
	nextAmount := s.getNextAmount(goods)
	// 更新红包订单的数量和金额
	rowsAffected, err := goodsService.UpdateBalance(goods.EnvelopeNo, nextAmount)
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("没有足够的红包金额了")
	}
	// 单个红包详情
	item = s.preCreatItem(goods, dto, nextAmount)
	item.Amount = nextAmount
	item.PayStatus = model.ENVELOPE_PAY_STATUS_PAYING
	// 保存单个红包的信息
	err = s.Insert(item)
	if err != nil {
		return nil, err
	}
	// 将抢到的红包金额从系统红包中间账户转入当前用户的资金账户
	err = s.transfer(dto, nextAmount)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// 根据已有信息,生成单个红包详情
func (s *EnvelopeGoodsItemService) preCreatItem(
	goods *model.EnvelopeGoods,
	dto *model.RedEnvelopeReceiveDTO,
	nextAmount decimal.Decimal,
) *model.EnvelopeGoodsItem {
	item := &model.EnvelopeGoodsItem{}
	item.ItemNo = s.createItemNo()
	item.AccountNo = dto.AccountNo
	item.EnvelopeNo = dto.EnvelopeNo
	item.RecvUserId = dto.RecvUserId
	item.RecvUsername = dto.RecvUsername
	item.RemainAmount = goods.RemainAmount.Sub(nextAmount)
	item.Desc = goods.Username + "的" + model.EnvelopeTypeDic[goods.Type]
	return item
}

func (s *EnvelopeGoodsItemService) getNextAmount(goods *model.EnvelopeGoods) decimal.Decimal {
	if goods.RemainQuantity == 1 {
		return goods.RemainAmount
	}
	if goods.Type == model.ENVELOPE_TYPE_GENERAL {
		return goods.AmountOne
	}
	if goods.Type == model.ENVELOPE_TYPE_LUCKY {
		multiple := decimal.NewFromFloat(100.0)
		cent := goods.RemainAmount.Mul(multiple).IntPart()
		nextCent := utils.DoubleAverage(int64(goods.RemainQuantity), cent)
		return decimal.NewFromFloat(float64(nextCent)).Div(multiple)
	}
	global.Logger.Error("不支持的红包类型")
	return decimal.NewFromFloat(0)
}

func (s *EnvelopeGoodsItemService) transfer(dto *model.RedEnvelopeReceiveDTO, amount decimal.Decimal) error {
	systemAccount := global.Config.System.Account
	body := model.Account{
		AccountNo: systemAccount.AccountNo,
		UserId:    systemAccount.UserId,
		Username:  systemAccount.Username,
	}
	target := model.Account{
		AccountNo: dto.AccountNo,
		UserId:    dto.RecvUserId,
		Username:  dto.RecvUsername,
	}
	tr := &model.AccountTransferDTO{
		TradeBody:   body,
		TradeTarget: target,
		TradeNo:     dto.EnvelopeNo,
		Amount:      amount,
		ChangeType:  model.ENVELOPE_INCOMING,
		ChangeFlag:  model.FLAG_TRANSFER_IN,
		Desc:        "红包收入：" + dto.EnvelopeNo,
	}
	accountService := NewAccountService(s.db)
	status, err := accountService.Transfer(tr)
	if status == model.TRANSFERRED_STATUS_SUCCESS {
		return nil
	}
	return err
}
