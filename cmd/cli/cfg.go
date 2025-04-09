package main

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/bigmikesolutions/wingman/client/a10n"
)

type (
	Config struct {
		LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
		A10N     A10Config
	}

	A10Config struct {
		KeycloakEndpoint string        `envconfig:"A10N_KEYCLOAK" default:"http://localhost:8080/"`
		Realm            string        `envconfig:"A10N_REALM" default:"wingman"`
		ClientID         string        `envconfig:"A10N_CLIENT"`
		MaxTime          time.Duration `envconfig:"A10N_MAX_TIME" default:"30s"`
	}
)

func LoadCfg() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("envconfig load: %w", err)
	}
	return cfg, nil
}

func (a A10Config) Opts() []a10n.Setting {
	return []a10n.Setting{
		a10n.WithKeycloak(a.KeycloakEndpoint, a.Realm),
		a10n.WithClientID(a.ClientID),
	}
}
