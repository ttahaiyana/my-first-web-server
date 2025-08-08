package myfirstwebserver

import "github.com/ttahaiyana/my-first-web-server/storage"

type Config struct {
	BindAddr string `toml:"bindaddr"`
	LogLevel string `toml:"log_level"`
	Storage  *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
