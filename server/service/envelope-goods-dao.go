package service

import (
	"errors"
	"red-server/global"
	"red-server/model"

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
