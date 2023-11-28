package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/lowl11/boost2/pkg/param"
	"os"
)

func (config *Config) Get(key string) *param.Param {
	return param.New(os.Getenv(key))
}

func (config *Config) Parse(result any) error {
	return envconfig.Process("", &result)
}
