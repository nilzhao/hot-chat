package controller

import (
	"red-server/global"
	"red-server/model"
	"red-server/service"
	"red-server/utils"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EnvelopeGoodsController struct{}

func NewEnvelopeGoodsController() *EnvelopeGoodsController {
	return &EnvelopeGoodsController{}
}

func (c *EnvelopeGoodsController) Detail(ctx *gin.Context) {
	envelopeNo := ctx.Param("envelopeNo")
	envelopeService := service.NewEnvelopeGoodsService(global.DB)
	goods := envelopeService.GetOne(envelopeNo)
	if goods == nil {
		utils.ResNotFound(ctx)
		return
	}
	utils.ResOk(ctx, goods)
}

func (c *EnvelopeGoodsController) SendOut(ctx *gin.Context) {
	goods := &model.EnvelopeGoods{}
	user := utils.GetCurrentUser(ctx)
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
	goods.UserId = user.Id
	goods.Username = user.Name
	// 查找账户
	accountService := service.NewAccountService(global.DB)
	account := accountService.GetByUserId(goods.UserId, model.ACCOUNT_TYPE_ENVELOPE)
	if account == nil {
		utils.ResFailed(ctx, errors.New("未开通红包账户"))
		return
	}
	goods.AccountNo = account.AccountNo
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		goodsService := service.NewEnvelopeGoodsService(tx)
		goods, err = goodsService.SendOut(goods)
		if goods != nil || err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, goods)
}

func (c EnvelopeGoodsController) Find(ctx *gin.Context) {
	goodsService := service.NewEnvelopeGoodsService(global.DB)
	user := utils.GetCurrentUser(ctx)
	list := goodsService.RandomGet(user.Id)
	if len(list) == 0 {
		utils.ResFailed(ctx, errors.New("空空如也"))
		return
	}
	utils.ResOk(ctx, list[0])
}

func (c EnvelopeGoodsController) Items(ctx *gin.Context) {
	itemService := service.NewEnvelopeGoodsItemService(global.DB)
	items := itemService.ListByEnvelopeNo(ctx.Param("envelopeNo"))
	utils.ResOk(ctx, items)
}

func (c *EnvelopeGoodsController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/goods/sendOut", c.SendOut)
	api.GET("/goods/find", c.Find)
	api.GET("/goods/:envelopeNo/items", c.Items)
	api.GET("/goods/:envelopeNo", c.Detail)
}
