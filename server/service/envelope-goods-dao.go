package service

import (
	"errors"
	"red-server/global"
	"red-server/model"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EnvelopeGoodsDao struct {
	db *gorm.DB
}

func NewEnvelopeGoodsDaoService(db *gorm.DB) *EnvelopeGoodsDao {
	return &EnvelopeGoodsDao{db}
}

func (s *EnvelopeGoodsDao) Insert(envelopeGoods *model.EnvelopeGoods) error {
	result := s.db.Create(envelopeGoods)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("红包创建失败")
	}
	return nil
}

func (s *EnvelopeGoodsDao) GetOne(envelopeNo string) *model.EnvelopeGoods {
	envelopeGoods := &model.EnvelopeGoods{}
	result := s.db.First(envelopeGoods, "envelope_no = ?", envelopeNo)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	return envelopeGoods
}

func (s *EnvelopeGoodsDao) UpdateBalance(envelopeNo string, amount decimal.Decimal) (int64, error) {
	sql := `
		update envelope_goods
		set
		remain_quantity=remain_quantity-1,
		remain_amount=remain_amount-CAST(? AS DECIMAL(30,2))
		where
		envelope_no = ? and 
		remain_quantity>0	and 
		remain_amount>=CAST(? AS DECIMAL(30,6))
	`
	result := s.db.Exec(sql, amount.String(), envelopeNo, amount.String())
	return result.RowsAffected, result.Error
}
