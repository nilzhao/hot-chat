// 日志配置
package config

import (
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type LogConfig struct {
	Level string `json:"level" yaml:"level"`
	// output
	Dir        string `json:"dir" yaml:"dir"`
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"maxSize" yaml:"maxSize"`
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"`
	LocalTime  bool   `json:"localTime" yaml:"localTime"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

// 日志格式
func (config *LogConfig) GetFormatter() (formatter *prefixed.TextFormatter) {
	formatter = &prefixed.TextFormatter{
		ForceColors:     true, // 控制台高亮
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2000-01-01 13:04:05.000000",
		ForceFormatting: true,
	}
	return formatter
}

// 日志等级
// panic fatal error warn info debug trace
func (config *LogConfig) GetLevel() logrus.Level {
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		panic(err)
	}
	return level
}

type LevelFileNameMap map[logrus.Level]string

var levelFileNameMap = LevelFileNameMap{
	logrus.DebugLevel: "debug",
	logrus.InfoLevel:  "info",
	logrus.WarnLevel:  "warn",
	logrus.ErrorLevel: "error",
	logrus.FatalLevel: "fatal",
	logrus.PanicLevel: "panic",
}

func (config *LogConfig) GetOutput(level logrus.Level) *lumberjack.Logger {
	fileName := levelFileNameMap[level]
	return &lumberjack.Logger{
		// 日志名称
		Filename: path.Join(config.Dir, fileName),
		// 日志大小限制，单位MB
		MaxSize: 100,
		// 历史日志文件保留天数
		MaxAge: 30,
		// 最大保留历史日志数量
		MaxBackups: 30,
		// 本地时区
		LocalTime: true,
		// 历史日志文件压缩标识
		Compress: false,
	}
}
