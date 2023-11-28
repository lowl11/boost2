package config

type Config struct {
	envFiles []string
}

var instance *Config

func Get(envFiles ...string) *Config {
	if instance != nil {
		return instance
	}

	instance = &Config{
		envFiles: envFiles,
	}
	instance.load()
	return instance
}
