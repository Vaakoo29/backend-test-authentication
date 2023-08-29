package store

// Config ...
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfing ...
func NewConfig() *Config {
	return&Config{}
}