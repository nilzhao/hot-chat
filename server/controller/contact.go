package controller

import (
	"errors"
	"red-server/global"
	"red-server/service"
	"red-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactController struct {
}

func NewContactController() Controller {
	return &ContactController{}
}

func (c *ContactController) AddFriend(ctx *gin.Context) {
	err := global.DB.Transaction(func(tx *gorm.DB) (err error) {
		contactService := service.NewContactService(tx, ctx)
		friendId, err := strconv.ParseInt(ctx.Param("friendId"), 10, 64)
		if err != nil {
			return errors.New("用户 ID 格式错误")
		}
		_, err = contactService.AddFriend(friendId)
		return err
	})
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, "添加成功")
}

func (c *ContactController) AddCommunity(ctx *gin.Context) {
	contactService := service.NewContactService(global.DB, ctx)
	communityId, err := strconv.ParseInt(ctx.Param("communityId"), 10, 64)
	if err != nil {
		utils.ResFailed(ctx, errors.New("群组 ID 格式错误"))
		return
	}
	_, err = contactService.AddCommunity(communityId)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, "添加成功")
}

func (c *ContactController) List(ctx *gin.Context) {
	contactService := service.NewContactService(global.DB, ctx)
	contacts := contactService.GetMyContacts()
	utils.ResOk(ctx, contacts)
}

func (c *ContactController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/contact/friends/:friendId", c.AddFriend)
	api.POST("/contact/communities/:communityId", c.AddCommunity)
	api.GET("/contacts", c.List)
}
