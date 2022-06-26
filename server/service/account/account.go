package accountService

import (
	"red-server/global"
	"red-server/model"
	"red-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func GenerateAccountNo() string {
	return ksuid.New().Next().String()
}

func Create(ctx *gin.Context) {
	account := model.Account{}
	if err := ctx.BindJSON(&account); err != nil {
		global.Logger.Error(err)
		utils.ResFail(ctx, "请求参数错误")
		return
	}
	account.AccountNo = GenerateAccountNo()
	result := global.DB.Create(&account)
	if result.Error != nil {
		global.Logger.Error(result.Error)
		utils.ResFail(ctx, result.Error.Error(), utils.ResponseFailInfo{
			Code: utils.ForbiddenCode,
		})
		return
	}
	utils.ResOk(ctx, account)
}
