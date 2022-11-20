package starter

import (
	"flag"
	"fmt"
	"hot-chat/config"
	"hot-chat/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"path/filepath"
)

type Env = string

const (
	ENV_DEV  Env = "dev"
	ENV_PROD Env = "prod"
)

type ConfigStarter struct {
	BaseStarter
}

func (s *ConfigStarter) Name() string {
	return "配置"
}

func (s *ConfigStarter) Init() {
	configPath := getConfig()
	parseConfig(configPath)
}

func getEnv() (env Env) {
	// 从命令行中获取 env
	flag.StringVar(&env, "env", "", "设置运行环境,可选: dev、prod")
	flag.Parse()
	// 从 环境变量中获取 env
	if env == "" {
		env = os.Getenv(config.Env)
	}
	// 默认 env
	if env == "" {
		env = ENV_DEV
	}
	return env
}

func getConfig() (configPath string) {
	env := getEnv()
	configPath = fmt.Sprintf("config.%s.yaml", env)
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
