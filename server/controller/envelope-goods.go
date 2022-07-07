package controller

import (
	"red-server/model"
	"red-server/utils"

	"github.com/gin-gonic/gin"
)

type EnvelopeGoods struct{}

func (c *EnvelopeGoods) SendOut(ctx *gin.Context) {
	goods := &model.EnvelopeGoods{}
	err := ctx.ShouldBind(goods)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	err = utils.ValidateStruct(goods)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}

}

func (c *EnvelopeGoods) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/users/:userId/")
}
