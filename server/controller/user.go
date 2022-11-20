package controller

import (
	"errors"
	"hot-chat/global"
	"hot-chat/service"
	"hot-chat/utils"
	"strconv"

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
	user = userService.GetOne(user.Id)
	if user == nil && !user.DeletedAt.Valid {
		utils.ResFailed(ctx, errors.New("用户不存在"))
		return
	}
	utils.ResOk(ctx, user)
}

// 搜索用户
func (c *UserController) SearchUser(ctx *gin.Context) {
	userService := service.NewUserService(global.DB)
	users := userService.Search(ctx.Query("keyword"))
	utils.ResOk(ctx, users)
}

func (c *UserController) Detail(ctx *gin.Context) {
	userService := service.NewUserService(global.DB)
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id <= 0 || err != nil {
		utils.ResFailed(ctx, errors.New("ID 格式错误"))
		return
	}
	user := userService.GetOne(id)
	if user == nil {
		utils.ResNotFound(ctx)
		return
	}
	utils.ResOk(ctx, user)
}

func (c *UserController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/user/profile", c.GetProfile)
	api.GET("/users/search", c.SearchUser)
	api.GET("/users/:id", c.Detail)
}
