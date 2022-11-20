package service

import (
	"errors"
	"hot-chat/model"
	"hot-chat/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactService struct {
	db  *gorm.DB
	ctx *gin.Context
}

func NewContactService(db *gorm.DB, ctx *gin.Context) *ContactService {
	return &ContactService{ctx: ctx, db: db}
}

func (s *ContactService) AddFriend(friendId int64) (contact *model.Contact, err error) {
	currentUser := utils.GetCurrentUser(s.ctx)
	if currentUser.Id == friendId {
		return nil, errors.New("不能添加自己为好友")
	}
	isFriend := s.IsMyFriend(friendId)
	if isFriend {
		return nil, errors.New("已是好友,不可重复添加")
	}

	err = s.Insert(&model.Contact{
		OwnerId:  currentUser.Id,
		TargetId: friendId,
		Type:     model.CONTACT_TYPE_FRIEND,
	})
	if err != nil {
		return nil, err
	}
	err = s.Insert(&model.Contact{
		OwnerId:  friendId,
		TargetId: currentUser.Id,
		Type:     model.CONTACT_TYPE_FRIEND,
	})
	if err != nil {
		return nil, err
	}
	return contact, err
}

func (s *ContactService) JoinCommunity(communityId int64) (contact *model.Contact, err error) {
	isIn := s.IsMyCommunity(communityId)
	if isIn {
		err = errors.New("已在群组之中,不可重复添加")
		return nil, err
	}
	currentUser := utils.GetCurrentUser(s.ctx)

	err = s.Insert(&model.Contact{
		OwnerId:  currentUser.Id,
		TargetId: communityId,
		Type:     model.CONTACT_TYPE_COMMUNITY,
	})
	if err != nil {
		return nil, err
	}
	return contact, err
}

func (s *ContactService) GetMyContacts() []*model.Contact {
	currentUser := utils.GetCurrentUser(s.ctx)
	return s.GetByOwnerId(currentUser.Id)
}
