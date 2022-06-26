package controller

import (
	accountService "red-server/service/account"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func NewAccountController() Controller {
	return &AccountController{}
}

// 创建账户
func (c *AccountController) Create(ctx *gin.Context) {
	accountService.Create(ctx)
}

func (c *AccountController) Name() string {
	return "account"
}

func (c *AccountController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/accounts", c.Create)
}
