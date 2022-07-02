package controller

import (
	"errors"
	"fmt"
	"red-server/global"
	"red-server/model"
	"red-server/service"
	"red-server/utils"

	"red-server/core"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AccountController struct{}

func NewAccountController() Controller {
	return &AccountController{}
}

// 创建账户
func (c *AccountController) Create(ctx *core.Context) {
	account := &model.Account{}
	if err := ctx.BindJSON(account); err != nil {
		global.Logger.Error(err)
		ctx.ResFail("请求参数错误")
		return
	}
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		accountService := service.NewAccountService(tx)
		accountLogService := service.NewAccountLogService(tx)

		err := accountService.Create(account)
		if err != nil {
			return err
		}
		// 生成日志
		accountLog := accountLogService.GenerateAccountCreatedLog(account)
		err = accountLogService.Create(&accountLog)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		ctx.ResFail("创建失败" + err.Error())
		return
	}
	ctx.ResOk(account)
}

func (c *AccountController) Detail(ctx *core.Context) {
	accountService := service.NewAccountService(global.DB)
	account := accountService.GetByNo(ctx.Param("accountNo"))
	if account == nil {
		ctx.ResNotFound()
		return
	}
	ctx.ResOk(account)
}

func (c *AccountController) Transfer(ctx *core.Context) {
	// 先查找账户
	accountService := service.NewAccountService(global.DB)
	sourceAccountNo := ctx.Param("accountNo")
	targetAccountNo := ctx.Param("targetAccountNo")
	sourceAccount := accountService.GetByNo(sourceAccountNo)
	if sourceAccount == nil {
		ctx.ResNotFound(fmt.Sprintf("账号[%s]不存在", sourceAccountNo))
		return
	}
	targetAccount := accountService.GetByNo(targetAccountNo)
	if targetAccount == nil {
		ctx.ResNotFound(fmt.Sprintf("账号[%s]不存在", targetAccountNo))
		return
	}
	// 组合转账的参数
	transferDto := &model.AccountTransferDTO{}
	err := ctx.BindJSON(transferDto)
	if err != nil {
		ctx.ResFail(err.Error())
		return
	}
	amount, err := decimal.NewFromString(transferDto.AmountStr)
	if err != nil {
		ctx.ResFail(err.Error())
		return
	}
	transferDto.Amount = amount
	transferDto.TradeBody = *sourceAccount
	transferDto.TradeTarget = *targetAccount
	// 验证参数
	err = utils.ValidateStruct(transferDto)
	if err != nil {
		ctx.ResFail(err.Error())
		return
	}
	// 开始转账
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		accountService := service.NewAccountService(tx)
		status, err := accountService.Transfer(transferDto)
		if err != nil {
			return err
		}
		if status != model.TRANSFERRED_STATUS_SUCCESS {
			return errors.New("转账失败")
		}
		return nil
	})
	if err != nil {
		ctx.ResFail(err.Error())
		return
	}
	ctx.ResOk(nil)
}

func (c *AccountController) Name() string {
	return "account"
}

func (c *AccountController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/accounts", core.CreateHandlerFunc(c.Create))
	api.GET("/accounts/:accountNo", core.CreateHandlerFunc(c.Detail))
	api.POST("/accounts/:accountNo/transfer/:targetAccountNo", core.CreateHandlerFunc(c.Transfer))
}