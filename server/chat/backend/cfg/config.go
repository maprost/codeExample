package cfg

import (
	"time"
)

type ServerConfig struct {
	Port string
	Host string
}

type Config struct {
	Server         ServerConfig
	DateTimeFormat string
	DateFormat     string
	Now            func() time.Time
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: "8080",
			Host: "",
		},
		DateTimeFormat: "2006-01-02 15:04:05",
		DateFormat:     "2006-01-02",
		Now:            Now,
	}
}

func Now() time.Time {
	return time.Now().Local().Truncate(time.Second)
}
