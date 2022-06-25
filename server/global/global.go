package global

import (
	"red-server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Config
	Logger *logrus.Logger
)
