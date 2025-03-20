package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/bigmikesolutions/wingman/server"
	"github.com/bigmikesolutions/wingman/server/vault"
)

type Config struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	HTTP     server.HTTPConfig
	Vault    vault.Config
	Database DatabaseConfig
	A10N     A10NConfig
}

func LoadCfg() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("envconfig load: %w", err)
	}
	return cfg, nil
}
