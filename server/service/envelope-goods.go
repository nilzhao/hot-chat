package service

import (
	"red-server/model"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EnvelopeGoods struct {
	db *gorm.DB
}

func NewEnvelopeGoodsService(db *gorm.DB) *EnvelopeGoods {
	return &EnvelopeGoods{db}
}

func createEnvelopeNo() string {
	return ksuid.New().Next().String()
}

func (s *EnvelopeGoods) GetByNo(envelopeNo string) *model.EnvelopeGoods {
	envelopeGoodsDao := NewEnvelopeGoodsDaoService(s.db)
	return envelopeGoodsDao.GetOne(envelopeNo)
}

func (s *EnvelopeGoods) Create(envelopeGoods *model.EnvelopeGoods) error {
	envelopeGoodsDao := NewEnvelopeGoodsDaoService(s.db)
	envelopeGoods.EnvelopeNo = createEnvelopeNo()
	return envelopeGoodsDao.Inert(envelopeGoods)
}

func (s *EnvelopeGoods) SendOut(goods model.EnvelopeGoods) {

}

// 生成红包订单
func (c *EnvelopeGoods) GenerateCreatedGoods(goods model.EnvelopeGoods) model.EnvelopeGoods {
	goods.RemainQuantity = goods.Quantity
	goods.Username.Valid = true
	goods.Blessing.Valid = true
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
