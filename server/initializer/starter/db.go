package starter

import (
	"fmt"
	"log"
	"os"
	"server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBStarter struct {
	BaseStarter
	gormConfig  *gorm.Config
	mysqlConfig mysql.Config
}

func (s *DBStarter) Init() {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	defaultLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	config.Logger = defaultLogger
	s.gormConfig = config
	dsn := global.CONFIG.DB.Dsn()

	mysqlConfig := mysql.Config{
		DriverName:                global.CONFIG.DB.DriverName,
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	s.mysqlConfig = mysqlConfig
}

func (s *DBStarter) Start() {
	db, err := gorm.Open(mysql.New(s.mysqlConfig), s.gormConfig)
	if err != nil {
		log.Panic("数据库连接失败")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(global.CONFIG.DB.MaxIdleConns)
		sqlDB.SetMaxOpenConns(global.CONFIG.DB.MaxOpenConns)
		fmt.Println("数据库连接成功")
	}
}
