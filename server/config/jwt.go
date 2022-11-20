package config

type Jwt struct {
	Salt     []byte `json:"salt" yaml:"salt"`
	Period   int64  `json:"period" yaml:"period"`
	Issuer   string `json:"issuer" yaml:"issuer"`
	TokenKey string `json:"key" yaml:"key"`
}
