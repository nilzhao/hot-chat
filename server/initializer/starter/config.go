package starter

import (
	"flag"
	"fmt"
	"hot-chat/config"
	"hot-chat/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"path/filepath"
)

type ConfigStarter struct {
	BaseStarter
}

func (s *ConfigStarter) Name() string {
	return "配置"
}

func (s *ConfigStarter) Init() {
	env := getEnv()
	if env == config.ENV_PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	configPath := getConfig(env)
	parseConfig(configPath)
}

func getEnv() (env config.Env) {
	// 从命令行中获取 env
	flag.StringVar(&env, "e", "", "设置运行环境,可选: dev、prod")
	flag.Parse()
	// 从 环境变量中获取 env
	if env == "" {
		env = os.Getenv(config.ENV_OPTION)
	}
	// 默认 env
	if env == "" {
		env = config.ENV_DEV
	}
	return env
}

func getConfig(env config.Env) (configPath string) {
	return fmt.Sprintf("config.%s.yaml", env)
}

func parseConfig(configPath string) {
	v := viper.New()
	configPath, err := filepath.Abs(filepath.Join("config/yaml", configPath))
	if err != nil {
		panic(fmt.Errorf("文件地址错误: %s", err))
	}
	fmt.Println("abs configPath", configPath)
	fmt.Printf("配置文件地址 %s\n", configPath)
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件解析错误: %s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
}
