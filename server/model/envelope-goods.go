package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type EnvelopeType int

const (
	ENVELOPE_TYPE_GENERAL EnvelopeType = 1
	ENVELOPE_TYPE_LUCKY   EnvelopeType = 2
)

type EnvelopeStatus int

const (
	ENVELOPE_STATUS_CREATE                    EnvelopeStatus = 1 // 创建
	ENVELOPE_STATUS_SENDING                   EnvelopeStatus = 2 // 发布
	ENVELOPE_STATUS_EXPIRED                   EnvelopeStatus = 3 // 过期
	ENVELOPE_STATUS_DISABLED                  EnvelopeStatus = 4 // 失效
	ENVELOPE_STATUS_EXPIRED_REFUND_SUCCESSFUL EnvelopeStatus = 5 // 过期退款成功
	ENVELOPE_STATUS_EXPIRED_REFUND_FAILED     EnvelopeStatus = 6 // 过期退款失败
)

type EnvelopeOrderType int

const (
	ENVELOPE_ORDER_TYPE_SENDING EnvelopeOrderType = 1 // 发布单
	ENVELOPE_ORDER_REFUND       EnvelopeOrderType = 2 // 退款单
)

type EnvelopePayStatus int

const (
	ENVELOPE_PAY_STATUS_PAY_NOTHING EnvelopePayStatus = 1
	ENVELOPE_PAY_STATUS_PAYING      EnvelopePayStatus = 2
	ENVELOPE_PAY_STATUS_PAYED       EnvelopePayStatus = 3
	ENVELOPE_PAY_STATUS_PAY_FAILED  EnvelopePayStatus = 4

	ENVELOPE_REFUND_NOTHING EnvelopePayStatus = 61
	ENVELOPE_REFUNDING      EnvelopePayStatus = 62
	ENVELOPE_REFUNDED       EnvelopePayStatus = 63
	ENVELOPE_REFUND_FAILED  EnvelopePayStatus = 64
)

type EnvelopeGoods struct {
	BaseModel
	EnvelopeNo       string            `json:"envelopeNo"`               // 红包编号
	Type             EnvelopeType      `json:"type" validate:"required"` // 红包类型:普通红包/碰运气红包
	UserId           string            `json:"userId" validate:"required"`
	Username         string            `json:"username" validate:"required"`
	Amount           decimal.Decimal   `json:"amount"`           // 总金额
	AmountOne        decimal.Decimal   `json:"amountOne"`        // 普通红包单个红包金额
	Quantity         int               `json:"quantity"`         // 红包数量
	RemainAmount     decimal.Decimal   `json:"remainAmount"`     // 剩余金额
	RemainQuantity   int               `json:"remainQuantity"`   // 红包剩余数量
	ExpiredAt        time.Time         `json:"expired"`          // 过期时间
	Blessing         string            `json:"blessing"`         // 祝福语
	Status           EnvelopeStatus    `json:"status"`           // 状态: 0红包初始化;1 启动;2 失效
	OrderType        EnvelopeOrderType `json:"orderType"`        // 订单状态: 发布单/退款单
	PayStatus        EnvelopePayStatus `json:"payStatus"`        // 支付状态: 未支付/支付中/已支付/支付失败
	OriginEnvelopeNo string            `json:"originEnvelopeNo"` // 原订单编号: 退款订单
}
