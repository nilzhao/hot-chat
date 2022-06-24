package config

import (
	"server/utils"
	"strings"
	"time"
)

type DBConfig struct {
	DriverName      string          `json:"driverName" yaml:"driverName"`
	Host            string          `json:"host" yaml:"host"`
	Port            string          `json:"port" yaml:"port"`
	Database        string          `json:"database" yaml:"database"`
	User            string          `json:"user" yaml:"user"`
	Password        string          `json:"password" yaml:"password"`
	ConnMaxLifetime time.Duration   `json:"connMaxLifetime" yaml:"connMaxLifetime"`
	MaxIdleConns    int             `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns    int             `json:"maxOpenConns" yaml:"maxOpenConns"`
	LoggingEnabled  bool            `json:"loggingEnabled" yaml:"loggingEnabled"`
	Options         DBConfigOptions `json:"options" yaml:"options"`
}

type DBConfigOptions struct {
	Charset   string `json:"charset" yaml:"charset"`
	ParseTime string `json:"parseTime" yaml:"parseTime"`
	Loc       string `json:"loc" yaml:"loc"`
}

func (config *DBConfig) getOptions() string {
	arr := make([]string, 0)
	optionMap := utils.StructToMap(config.Options)
	for k, v := range optionMap {
		arr = append(arr, k+string(v.(string)))
	}
	return strings.Join(arr, "&")
}

func (config *DBConfig) Dsn() string {
	return config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?" + config.getOptions()
}
