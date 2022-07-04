package controller

import (
	"red-server/core"
	"red-server/model"
	"red-server/utils"

	"github.com/gin-gonic/gin"
)

type EnvelopeGoods struct{}

func (c *EnvelopeGoods) SendOut(ctx *core.Context) {
	goods := &model.EnvelopeGoods{}
	err := ctx.ShouldBind(goods)
	if err != nil {
		ctx.ResFailed(err)
		return
	}
	err = utils.ValidateStruct(goods)
	if err != nil {
		ctx.ResFailed(err)
		return
	}

}

func (c *EnvelopeGoods) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/users/:userId/")
}

func (c *EnvelopeGoods) TableName() string {
	return "envelopeGoods"
}
