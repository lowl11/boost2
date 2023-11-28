package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func (config *Config) load() {
	envFileName := ".env"
	if len(config.envFiles) > 0 {
		envFileName = config.envFiles[0]
	}

	_, err := os.Stat(envFileName)
	if os.IsNotExist(err) {
		return
	}

	if err = godotenv.Load(config.envFiles...); err != nil {
		log.Fatal("Load configuration error: ", err)
	}
}
