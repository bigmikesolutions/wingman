package vault

import (
	"context"
	"fmt"
	"log"
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

var secretMountPath = vault.WithMountPath("secret")

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
	// TODO make it a debug
	log.Printf("reading secret with a path %s", path)

	sv, err := s.vault.Secrets.KvV2Read(ctx, path, secretMountPath)
	if err != nil {
		if vault.IsErrorStatus(err, http.StatusNotFound) {
			// TODO convert this error here
			return nil
		}
		return fmt.Errorf("read secret: %w", err)
	}

	log.Printf("read secret: %v", sv.Data.Data)
	if err := unmarshall(sv.Data.Data, &v); err != nil {
		return fmt.Errorf("read secret: %w", err)
	}

	return nil
}

func (s *Secrets) Write(ctx context.Context, path string, v any) error {
	sv, err := marshall(v)
	if err != nil {
		return err
	}

	// TODO make it a debug
	log.Printf("writing secret %+v with a path %s", v, path)

	_, err = s.vault.Secrets.KvV2Write(
		ctx,
		path,
		schema.KvV2WriteRequest{
			Data: sv,
		},
		secretMountPath,
	)
	if err != nil {
		return fmt.Errorf("write secret: %w", err)
	}

	return nil
}
