package controller

import (
	"errors"
	"red-server/global"
	"red-server/model"
	"red-server/service"
	"red-server/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// 用户注册
func (c *AuthController) Register(ctx *gin.Context) {
	userDto := &model.UserCreateDto{}
	err := ctx.ShouldBind(userDto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	err = utils.ValidateStruct(userDto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	userService := service.NewUserService(global.DB)
	user, err := userService.Create(userDto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, user)
}

// 登录
func (c *AuthController) Login(ctx *gin.Context) {
	userService := service.NewUserService(global.DB)
	userDto := &model.UserCreateDto{}
	err := ctx.ShouldBind(userDto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	err = utils.ValidateStruct(userDto)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	user := userService.GetByEmail(userDto.Email)
	// 账号是否存在
	if user == nil {
		utils.ResFailed(ctx, errors.New("账号不存在"))
		return
	}
	// 验证密码是否正确
	if ok := utils.BcryptCheck(user.Password, userDto.Password); !ok {
		utils.ResFailed(ctx, errors.New("账号或密码错误"))
		return
	}
	// 账号状态是否可用
	if user.Status != model.USER_STATUS_NORMAL {
		utils.ResFailed(ctx, errors.New("账号不可用"))
		return
	}
	token, expiredAt := utils.CreateToken(utils.BaseClaims{
		User: *user,
	})
	utils.ResOk(ctx, map[string]interface{}{
		"token":     token,
		"expiredAt": expiredAt,
	})
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	utils.ResOk(ctx, "刷新 token")
}

func (c *AuthController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/register", c.Register)
	api.POST("/login", c.Login)
	api.POST("/token/refresh", c.Login)
}
