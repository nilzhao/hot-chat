// 日志
package starter

import (
	"hot-chat/global"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type LogStarter struct {
	BaseStarter
}

func (s *LogStarter) Name() string {
	return "日志"
}

func (s *LogStarter) Init() {
	logger := logrus.StandardLogger()
	// 定义日志格式
	formatter := global.CONFIG.Log.GetFormatter()
	logger.SetFormatter(formatter)

	// 日志级别
	level := global.CONFIG.Log.GetLevel()
	logger.SetLevel(level)

	// 日志级别分隔
	writerMap := lfshook.WriterMap{}
	for _, level := range logrus.AllLevels {
		writerMap[level] = global.CONFIG.Log.GetOutput(level)
	}
	lfHook := lfshook.NewHook(writerMap, formatter)
	logger.AddHook(lfHook)
	global.Logger = logger
}
