package config

const (
	Env         = "ENV"
	DefaultFile = "config.yaml"
	DevFile     = "config.dev.yaml"
	TestFile    = "config.test.yaml"
	ProdFile    = "config.prod.yaml"
)

type Config struct {
	System SystemConfig `json:"system" yaml:"system"`
	DB     DBConfig     `json:"db" yaml:"db"`
	Log    LogConfig    `json:"log" yaml:"log"`
	Jwt    Jwt          `json:"jwt" yaml:"jwt"`
}
