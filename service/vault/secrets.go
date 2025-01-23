package vault

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type (
	Secrets struct {
		cfg   Config
		vault *vault.Client
	}
)

func New(ctx context.Context, cfg Config) (*Secrets, error) {
	client, err := newClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Secrets{
		cfg:   cfg,
		vault: client,
	}, nil
}

func (s *Secrets) Read(ctx context.Context, path string, v any) error {
	sv, err := s.vault.Secrets.KvV2Read(ctx, path, vault.WithMountPath("secret"))
	if err != nil {
		if vault.IsErrorStatus(err, http.StatusNotFound) {
			// TODO convert this error here
			return nil
		}
		return fmt.Errorf("read secret: %w", err)
	}

	if err := unmarshall(sv.Data.Data, v); err != nil {
		return fmt.Errorf("read secret: %w", err)
	}

	return nil
}

func (s *Secrets) Write(ctx context.Context, path string, v any) error {
	sv, err := marshall(v)
	if err != nil {
		return err
	}

	_, err = s.vault.Secrets.KvV2Write(
		ctx,
		path,
		schema.KvV2WriteRequest{
			Data: sv,
		},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		return fmt.Errorf("write secret: %w", err)
	}

	return nil
}
