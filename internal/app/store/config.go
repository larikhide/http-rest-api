package store

import "github.com/sirupsen/logrus"

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAdd: ":8080",
		LogLevel: "debug",
	}
}