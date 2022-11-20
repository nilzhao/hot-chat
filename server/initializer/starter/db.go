package starter

import (
	"hot-chat/global"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBStarter struct {
	BaseStarter
	gormConfig  *gorm.Config
	mysqlConfig mysql.Config
}

func (s *DBStarter) Name() string {
	return "数据库"
}

func (s *DBStarter) Init() {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	}
	defaultLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	config.Logger = defaultLogger
	s.gormConfig = config
	dsn := global.Config.DB.Dsn()
	mysqlConfig := mysql.Config{
		DriverName:                global.Config.DB.DriverName,
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	s.mysqlConfig = mysqlConfig
}

func (s *DBStarter) Start() {
	db, err := gorm.Open(mysql.New(s.mysqlConfig), s.gormConfig)
	if err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(global.Config.DB.MaxIdleConns)
		sqlDB.SetMaxOpenConns(global.Config.DB.MaxOpenConns)
	}
	global.DB = db
}
