package starter

import (
	"flag"
	"fmt"
	"os"
	"red-server/config"
	"red-server/global"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"path/filepath"
)

type ConfigStarter struct {
	BaseStarter
}

func (s *ConfigStarter) Init() {
	fmt.Println("初始化配置...")
	configPath := getConfig()
	parseConfig(configPath)
	fmt.Println("初始化配置成功")
}

func getConfig() (configPath string) {
	flag.StringVar(&configPath, "c", "", "choose config file.")
	flag.Parse()
	if configPath == "" { // 判断命令行参数是否为空
		if configEnv := os.Getenv(config.Env); configEnv == "" { // 判断 config.Env 常量存储的环境变量是否为空
			switch gin.Mode() {
			case gin.DebugMode:
				configPath = config.DevFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, config.DevFile)
			case gin.ReleaseMode:
				configPath = config.ProdFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, config.ProdFile)
			case gin.TestMode:
				configPath = config.DevFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, config.DevFile)
			}
		} else { // config.Env 常量存储的环境变量不为空 将值赋值于config
			configPath = configEnv
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", config.Env, configPath)
		}
	} else { // 命令行参数不为空 将值赋值于 configPath
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", configPath)
	}
	return configPath
}

func parseConfig(configPath string) {
	v := viper.New()
	fmt.Printf("配置文件地址 %s\n", filepath.Join("config/yaml", configPath))
	v.SetConfigFile(filepath.Join("config/yaml", configPath))
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件解析错误: %s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
}
