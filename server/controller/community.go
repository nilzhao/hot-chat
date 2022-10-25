package controller

import (
	"errors"
	"red-server/global"
	"red-server/model"
	"red-server/service"
	"red-server/utils"

	"github.com/gin-gonic/gin"
)

type CommunityController struct{}

func NewCommunityController() *CommunityController {
	return &CommunityController{}
}

func (c *CommunityController) Create(ctx *gin.Context) {
	communityService := service.NewCommunityService(global.DB)
	community := &model.Community{}
	err := ctx.ShouldBind(community)
	if err != nil {
		global.Logger.Error(err)
		utils.ResFailed(ctx, errors.New("格式错误"))
		return
	}
	err = communityService.Insert(community)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, community)
}

func (c *CommunityController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/communities", c.Create)
}
