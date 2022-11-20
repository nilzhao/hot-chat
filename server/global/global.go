package global

import (
	"hot-chat/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Config
	Logger *logrus.Logger
)
