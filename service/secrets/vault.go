package secrets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type (
	SecretValue map[string]any

	Secrets struct {
		cfg   Config
		vault *vault.Client
	}
)

func New(ctx context.Context, cfg Config) (*Secrets, error) {
	client, err := vault.New(cfg.vaultOptions()...)
	if err != nil {
		return nil, fmt.Errorf("vault client: %w", err)
	}

	token, err := vaultAuth(ctx, client, cfg)
	if err != nil {
		return nil, fmt.Errorf("vault auth: %w", err)
	}

	if err := client.SetToken(token); err != nil {
		return nil, fmt.Errorf("vault token: %w", err)
	}

	return &Secrets{
		cfg:   cfg,
		vault: client,
	}, nil
}

func (s *Secrets) Read(ctx context.Context, path string) (SecretValue, error) {
	v, err := s.vault.Secrets.KvV2Read(ctx, path, vault.WithMountPath("secret"))
	if err != nil {
		if vault.IsErrorStatus(err, http.StatusNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("read secret: %w", err)
	}
	return v.Data.Data, nil
}

func (s *Secrets) Write(ctx context.Context, path string, data SecretValue) error {
	_, err := s.vault.Secrets.KvV2Write(
		ctx,
		path,
		schema.KvV2WriteRequest{
			Data: data,
		},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		return fmt.Errorf("write secret: %w", err)
	}
	return nil
}
