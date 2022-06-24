package global

import (
	"server/config"

	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Config
)
