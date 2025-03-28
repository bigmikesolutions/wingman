package main

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/bigmikesolutions/wingman/server/vault"
)

type (
	Config struct {
		LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
		HTTP     HTTPConfig
		Vault    vault.Config
		Database DatabaseConfig
		A10N     A10NConfig
	}

	HTTPConfig struct {
		Address       string        `envconfig:"HTTP_ADDRESS" default:"0.0.0.0:8080"`
		WriteTimeout  time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"15s"`
		ReadTimeout   time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"15s"`
		ShutdownTime  time.Duration `envconfig:"HTTP_SHUTDOWN_TIME" default:"30s"`
		PprofEnabled  bool          `envconfig:"HTTP_PPROF_ENABLED" default:"false"`
		CompressLevel int           `envconfig:"HTTP_COMPRESS_LEVEL" default:"5"`
	}
)

func LoadCfg() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("envconfig load: %w", err)
	}
	return cfg, nil
}
