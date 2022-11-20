package config

const (
	ENV_OPTION = "ENV"
)

type Env = string

const (
	ENV_DEV  Env = "dev"
	ENV_PROD Env = "prod"
)

type Config struct {
	System SystemConfig `json:"system" yaml:"system"`
	DB     DBConfig     `json:"db" yaml:"db"`
	Log    LogConfig    `json:"log" yaml:"log"`
	Jwt    Jwt          `json:"jwt" yaml:"jwt"`
	Attach Attach       `json:"attach" yaml:"attach"`
}
