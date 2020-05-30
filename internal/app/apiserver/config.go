package apiserver

import "github.com/larikhide/http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"` //address where a server is started
	LogLevel string `toml:"log_level"` //because logrus support different levels of logging
	Store *store.Config //config struct for store
}

// NewConfig returns default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
