package controller

import (
	accountService "red-server/service/account"

	"red-server/core"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func NewAccountController() Controller {
	return &AccountController{}
}

// 创建账户
func (c *AccountController) Create(ctx *core.Context) {
	accountService.Create(ctx)
}

func (c *AccountController) Detail(ctx *core.Context) {
	accountService.GetByNo(ctx, ctx.Param("accountNo"))
}

func (c *AccountController) Name() string {
	return "account"
}

func (c *AccountController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/accounts", core.CreateHandlerFunc(c.Create))
	api.GET("/accounts/:accountNo", core.CreateHandlerFunc(c.Detail))
}
