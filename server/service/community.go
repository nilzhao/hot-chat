package service

import (
	"gorm.io/gorm"
)

type CommunityService struct {
	db *gorm.DB
}

func NewCommunityService(db *gorm.DB) *CommunityService {
	return &CommunityService{db: db}
}
