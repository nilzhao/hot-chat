package model

import (
	"time"

	"gorm.io/gorm"
)

// 日志类的 model,无需更新和删除
type LogBaseModel struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type BaseModel struct {
	LogBaseModel
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
