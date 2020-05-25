package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"` //address where a server is started
}

// NewConfig returns default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
