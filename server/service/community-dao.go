package service

import (
	"red-server/global"
	"red-server/model"
)

func (s *CommunityService) Insert(community *model.Community) error {
	result := s.db.Create(community)
	return result.Error
}

func (s *CommunityService) GetOne(id int64) *model.Community {
	community := &model.Community{}

	result := s.db.Find(community, id)
	if result.Error != nil {
		global.Logger.Error(result.Error)
	}

	return community
}
