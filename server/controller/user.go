package controller

import (
	"errors"
	"red-server/global"
	"red-server/service"
	"red-server/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// 个人信息
func (c *UserController) GetProfile(ctx *gin.Context) {
	userService := service.NewUserService(global.DB)
	user := utils.GetCurrentUser(ctx)
	user = userService.Get(user.Id)
	if user == nil && !user.DeletedAt.Valid {
		utils.ResFailed(ctx, errors.New("用户不存在"))
		return
	}
	utils.ResOk(ctx, user)
}

func (c *UserController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/user/profile", c.GetProfile)
}
