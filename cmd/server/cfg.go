package main

import (
	"fmt"
	"github.com/bigmikesolutions/wingman/service"
	"github.com/bigmikesolutions/wingman/service/vault"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	HTTP     service.HTTPConfig
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
