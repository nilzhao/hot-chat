package model

import "github.com/shopspring/decimal"

type EnvelopeGoodsItem struct {
	BaseModel
	ItemNo       string            `json:"itemNo"`
	EnvelopeNo   string            `json:"envelopNo"`
	RecvUsername string            `json:"recvUsername"`
	RecvUserId   uint              `json:"recvUserId"`
	Amount       decimal.Decimal   `json:"amount"`       // 收到的金额
	RemainAmount decimal.Decimal   `json:"remainAmount"` // 剩余金额
	AccountNo    string            `json:"accountNo"`
	PayStatus    EnvelopePayStatus `json:"payStatus"`
	Desc         string            `json:"desc"`
}

type RedEnvelopeReceiveDTO struct {
	EnvelopeNo   string `json:"envelopeNo" validate:"required"` // 红包编号,红包唯一标识
	RecvUserId   uint   `json:"recvUserId"`                     // 红包接收者用户编号
	RecvUsername string `json:"recvUsername"`                   // 红包接收者用户名称
	AccountNo    string `json:"accountNo"`
}
