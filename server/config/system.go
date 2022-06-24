package config

type SystemConfig struct {
	Port int    `json:"port" yaml:"port,omitempty"`
	Name string `json:"name" yaml:"name,omitempty"`
}
