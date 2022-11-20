package controller

import (
	"errors"
	"hot-chat/global"
	"hot-chat/model"
	"hot-chat/service"
	"hot-chat/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EnvelopeGoodsItemController struct {
}

func NewEnvelopeGoodsItemController() *EnvelopeGoodsItemController {
	return &EnvelopeGoodsItemController{}
}

func (c *EnvelopeGoodsItemController) Receive(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	dto := &model.RedEnvelopeReceiveDTO{}
	err := ctx.ShouldBind(dto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	err = utils.ValidateStruct(dto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	dto.RecvUserId = user.Id
	dto.RecvUsername = user.Name
	accountService := service.NewAccountService(global.DB)
	account := accountService.GetByUserId(user.Id, model.ACCOUNT_TYPE_ENVELOPE)
	if account == nil {
		utils.ResFailed(ctx, errors.New("用户未开通红包账户"))
		return
	}
	dto.AccountNo = account.AccountNo
	var item *model.EnvelopeGoodsItem
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		goodsItemService := service.NewEnvelopeGoodsItemService(tx)
		item, err = goodsItemService.Receive(dto)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, item)

}

func (c *EnvelopeGoodsItemController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/goods/receive", c.Receive)
}
