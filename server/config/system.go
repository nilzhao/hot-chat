// 系统配置
package config

type SystemConfig struct {
	Port int    `json:"port" yaml:"port"`
	Name string `json:"name" yaml:"name"`
	Host string `json:"host" yaml:"host"`
}
