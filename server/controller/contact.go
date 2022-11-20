package controller

import (
	"errors"
	"hot-chat/global"
	"hot-chat/service"
	"hot-chat/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactController struct {
}

func NewContactController() Controller {
	return &ContactController{}
}

// 添加好友
func (c *ContactController) AddFriend(ctx *gin.Context) {
	err := global.DB.Transaction(func(tx *gorm.DB) (err error) {
		contactService := service.NewContactService(tx, ctx)
		friendId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
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

// 加入群聊
func (c *ContactController) JoinCommunity(ctx *gin.Context) {
	contactService := service.NewContactService(global.DB, ctx)
	communityId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ResFailed(ctx, errors.New("群组 ID 格式错误"))
		return
	}
	_, err = contactService.JoinCommunity(communityId)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	utils.ResOk(ctx, "添加成功")
}

// 联系人列表
func (c *ContactController) MyContacts(ctx *gin.Context) {
	contactService := service.NewContactService(global.DB, ctx)
	contacts := contactService.GetMyContacts()
	utils.ResOk(ctx, contacts)
}

func (c *ContactController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/contact/friends/:id", c.AddFriend)
	api.POST("/contact/communities/:id", c.JoinCommunity)
	api.GET("/contacts/me", c.MyContacts)
}
