package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerID string `envconfig:"SERVER_ID" required:"true"`
}

func New() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
