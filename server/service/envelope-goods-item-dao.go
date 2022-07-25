package service

import (
	"errors"
	"red-server/global"
	"red-server/model"
)

func (s *EnvelopeGoodsItemService) Insert(item *model.EnvelopeGoodsItem) error {
	result := s.db.Create(item)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return errors.New("重复领取")
	}
	if result.RowsAffected <= 0 {
		return errors.New("创建红包失败")
	}
	return nil
}

func (s *EnvelopeGoodsItemService) ListByEnvelopeNo(envelopeNo string) []*model.EnvelopeGoodsItem {
	items := []*model.EnvelopeGoodsItem{}
	result := s.db.Where("envelope_no = ?", envelopeNo).Preload("RecvUser").Find(&items)
	if result.Error != nil {
		global.Logger.Error(result.Error)
	}
	return items
}
