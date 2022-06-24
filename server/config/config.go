package config

const (
	Env         = "CONFIG"
	DefaultFile = "config.yaml"
	DevFile     = "config.dev.yaml"
	TestFile    = "config.test.yaml"
	ProdFile    = "config.prod.yaml"
)

type Config struct {
	System SystemConfig `json:"system" yaml:"system,omitempty"`
	DB     DBConfig     `json:"db" yaml:"db,omitempty"`
}
