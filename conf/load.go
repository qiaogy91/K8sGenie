package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v11"
)

var config *Config

func LoadFromEnv() {
	ins := NewConfig()
	if err := env.Parse(ins); err != nil {
		panic("环境变量加载错误:: " + err.Error())
	}
	config = ins
}

func LoadFromToml(f string) {
	ins := NewConfig()
	if _, err := toml.DecodeFile(f, ins); err != nil {
		panic("配置文件加载错误::" + err.Error())
	}
	config = ins
}

func C() *Config {
	if config == nil {
		panic("请先加载配置文件")
	}
	return config
}
