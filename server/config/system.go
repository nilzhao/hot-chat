// 系统配置
package config

type SystemConfig struct {
	Port    int           `json:"port" yaml:"port"`
	Name    string        `json:"name" yaml:"name"`
	Host    string        `json:"host" yaml:"host"`
	Account SystemAccount `json:"account" yaml:"account"`
}

type SystemAccount struct {
	AccountNo string `json:"accountNo" yaml:"accountNo"`
	UserId    int64  `json:"userId" yaml:"userId"`
	Username  string `json:"username" yaml:"username"`
}
