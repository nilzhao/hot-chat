package accountService

import (
	"red-server/global"
	"red-server/model"

	"red-server/core"

	"github.com/segmentio/ksuid"
)

func GenerateAccountNo() string {
	return ksuid.New().Next().String()
}

func Create(ctx *core.Context) {
	account := model.Account{}
	if err := ctx.BindJSON(&account); err != nil {
		global.Logger.Error(err)
		ctx.ResFail("请求参数错误")
		return
	}
	account.AccountNo = GenerateAccountNo()
	result := global.DB.Create(&account)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		ctx.ResFail(result.Error.Error(), core.ResponseFailInfo{
			Code: core.ForbiddenCode,
		})
		return
	}
	ctx.ResOk(account)
}

func GetByNo(ctx *core.Context, accountNo string) {
	account := model.Account{}
	result := global.DB.Where(map[string]any{"account_no": accountNo}).Find(&account)
	if result.Error != nil {
		ctx.ResFail(result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		ctx.ResNotFound()
		return
	}
	ctx.ResOk(&account)
}
