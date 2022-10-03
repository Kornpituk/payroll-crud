package config

import (
	"os"
)

type Config struct {
	Env  string
	Port string

	ContextPath string
}

func New() (*Config, error) {

	var config Config

	config.Env = os.Getenv("GO_ENV")
	config.Port = os.Getenv("PORT")
	config.ContextPath = os.Getenv("CONTEXT_PATH")

	return &config, nil
}
