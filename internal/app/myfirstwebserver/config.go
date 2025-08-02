package myfirstwebserver

type Config struct {
	BindAddr string `toml:"bindaddr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
