package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"` //address where a server is started
	LogLevel string `toml:"log_level"` //because logrus support different levels of logging
}

// NewConfig returns default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
