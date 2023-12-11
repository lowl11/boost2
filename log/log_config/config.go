package log_config

import "github.com/lowl11/boost2/internal/infrastructure/logger/log_config"

type Config struct {
	JsonMode bool
	Level    string
}

func SetConfig(cfg *Config) {
	log_config.Config = cfg
}
