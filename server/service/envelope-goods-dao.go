package service

import (
	"errors"
	"red-server/global"
	"red-server/model"

	"github.com/shopspring/decimal"
)

func (s *EnvelopeGoodsService) Insert(envelopeGoods *model.EnvelopeGoods) error {
	result := s.db.Create(envelopeGoods)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("红包创建失败")
	}
	return nil
}

func (s *EnvelopeGoodsService) GetOne(envelopeNo string) *model.EnvelopeGoods {
	envelopeGoods := &model.EnvelopeGoods{}
	result := s.db.First(envelopeGoods, "envelope_no = ?", envelopeNo)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		return nil
	}
	return envelopeGoods
}

func (s *EnvelopeGoodsService) UpdateBalance(envelopeNo string, amount decimal.Decimal) (int64, error) {
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

// 随机返回
func (s *EnvelopeGoodsService) RandomGet(userId uint, limits ...int) []*model.EnvelopeGoods {
	limit := 1
	if len(limits) > 0 {
		limit = limits[0]
	}

	sql := `
		SELECT * FROM envelope_goods AS goods1
		JOIN (
			SELECT ROUND(
				RAND() *
				((SELECT MAX(id) FROM envelope_goods)-(SELECT MIN(id) FROM envelope_goods)) +
				(SELECT MIN(id) FROM envelope_goods)
			) AS random_id
		) AS goods2
		WHERE
			goods1.id >= goods2.random_id AND
			goods1.user_id != ? AND
			goods1.remain_amount>0 AND
			goods1.remain_quantity>0
		ORDER BY goods1.id LIMIT ?
	`

	var list []*model.EnvelopeGoods
	result := s.db.Raw(sql, userId, limit).Scan(&list)
	if result.Error != nil {
		global.Logger.Error(result.Error)
	}
	// TODO: 使用 JOIN
	for _, goods := range list {
		user := model.User{}
		s.db.First(&user, goods.UserId)
		goods.User = user
	}
	return list
}
