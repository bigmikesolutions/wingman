package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/bigmikesolutions/wingman/server/token"
)

type A10NConfig struct {
	PrivateKeyPath string        `envconfig:"A10N_PRIVATE_KEY_PATH"`
	PublicKeyPath  string        `envconfig:"A10N_PUBLIC_KEY_PATH"`
	TokenDuration  time.Duration `envconfig:"A10N_TOKEN_DURATION" default:"15m"`
}

func newA10N(cfg A10NConfig) (*token.JWT, error) {
	privateKey, err := os.Open(cfg.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load private key: %w", err)
	}

	publicKey, err := os.Open(cfg.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load public key: %w", err)
	}

	return token.New(privateKey, publicKey, token.Settings{
		SigningMethod: jwt.SigningMethodRS256,
		ExpTime:       cfg.TokenDuration,
	})
}
