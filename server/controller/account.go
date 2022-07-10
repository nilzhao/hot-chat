package controller

import (
	"errors"
	"fmt"
	"red-server/global"
	"red-server/model"
	"red-server/service"
	"red-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct{}

func NewAccountController() Controller {
	return &Account{}
}

// 创建账户
func (c *Account) Create(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)
	account := &model.Account{}
	if err := ctx.BindJSON(account); err != nil {
		global.Logger.Error(err)
		utils.ResFailed(ctx, err)
		return
	}
	account.UserId = user.Id
	account.Username = user.Name
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		accountService := service.NewAccountService(tx)
		accountLogService := service.NewAccountLogService(tx)

		err := accountService.Create(account)
		if err != nil {
			return err
		}
		// 生成日志
		accountLog := accountLogService.GenerateAccountCreatingLog(account)
		err = accountLogService.Create(&accountLog)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, account)
}

func (c *Account) Detail(ctx *gin.Context) {
	accountService := service.NewAccountService(global.DB)
	account := accountService.Get(ctx.Param("accountNo"))
	if account == nil {
		utils.ResNotFound(ctx)
		return
	}
	utils.ResOk(ctx, account)
}

func (c *Account) Transfer(ctx *gin.Context) {
	// 先查找账户
	accountService := service.NewAccountService(global.DB)
	sourceAccountNo := ctx.Param("accountNo")
	targetAccountNo := ctx.Param("targetAccountNo")
	sourceAccount := accountService.Get(sourceAccountNo)
	if sourceAccount == nil {
		utils.ResNotFound(ctx, fmt.Sprintf("账号[%s]不存在", sourceAccountNo))
		return
	}
	targetAccount := accountService.Get(targetAccountNo)
	if targetAccount == nil {
		utils.ResNotFound(ctx, fmt.Sprintf("账号[%s]不存在", targetAccountNo))
		return
	}
	// 组合转账的参数
	transferDto := &model.AccountTransferDTO{}
	err := ctx.BindJSON(transferDto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	amount, err := decimal.NewFromString(transferDto.AmountStr)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	transferDto.Amount = amount
	transferDto.TradeBody = *sourceAccount
	transferDto.TradeTarget = *targetAccount
	// 验证参数
	err = utils.ValidateStruct(transferDto)
	if err != nil {
		utils.ResFailed(ctx, err)
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
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, nil)
}

func (c *Account) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/accounts", c.Create)
	api.GET("/accounts/:accountNo", c.Detail)
	api.POST("/accounts/:accountNo/transfer/:targetAccountNo", c.Transfer)
}
