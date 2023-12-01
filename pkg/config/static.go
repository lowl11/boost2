package config

import (
	"github.com/lowl11/boost2/internal/infrastructure/config"
	"github.com/lowl11/boost2/log"
	"github.com/lowl11/boost2/pkg/param"
)

func EnvFiles(files ...string) {
	config.Get(files...)
}

func Get(key string) *param.Param {
	return config.Get().Get(key)
}

func Parse(result any) error {
	return config.Get().Parse(result)
}

func Load(filesNames ...string) error {
	return config.Get().Load(filesNames...)
}

func MustParse(result any) {
	if err := config.Get().Parse(result); err != nil {
		log.Fatal("Load config structure error: ", err)
	}
}
