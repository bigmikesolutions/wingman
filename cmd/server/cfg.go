package main

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
		HTTP     HTTPConfig
		Vault    VaultConfig
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

	VaultConfig struct {
		Address string        `envconfig:"VAULT_ADDRESS" default:""`
		Timeout time.Duration `envconfig:"VAULT_TIMEOUT" default:"10s"`

		// dev settings
		Token string `envconfig:"VAULT_TOKEN" default:""`

		// role auth
		RoleID       string `envconfig:"VAULT_ROLE_ID" default:""`
		RoleSecretID string `envconfig:"VAULT_ROLE_SECRET_ID" default:""`
		RolePath     string `envconfig:"VAULT_ROLE_PATH" default:"approle"`

		// TLS settings
		ServerCert string `envconfig:"VAULT_SERVER_CERT" default:""`

		// TLS with client-side cert auth
		ClientCert string `envconfig:"VAULT_CLIENT_CERT" default:""`
		ClientKey  string `envconfig:"VAULT_CLIENT_KEY" default:""`
	}
)

func LoadCfg() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("envconfig load: %w", err)
	}
	return cfg, nil
}
