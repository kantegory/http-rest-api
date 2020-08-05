package apiserver

import "github.com/kantegory/http-rest-api/internal/app/store/sqlstore"

// Config ...
type Config struct {
	BindAddr string `toml: "bind_addr"`
	LogLevel string `toml: "log_level"`
	Store *sqlstore.Config
}


// NewConfig ...
func NewConfig() *Config {
	return &Config {
		BindAddr: ":8080",
		LogLevel: "Debug",
		Store: sqlstore.NewConfig(),
	}
}
