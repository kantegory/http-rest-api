package apiserver

import "github.com/kantegory/http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml: "bind_addr"`
	LogLevel string `toml: "log_level"`
	Store *store.Config
}


// NewConfig ...
func NewConfig() *Config {
	return &Config {
		BindAddr: ":8080",
		LogLevel: "Debug",
		Store: store.NewConfig(),
	}
}
