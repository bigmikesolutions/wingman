package vault

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type (
	Secrets struct {
		cfg    settings
		vault  *vault.Client
		logger zerolog.Logger
	}
)

var secretMountPath = vault.WithMountPath("secret")

func New(ctx context.Context, opt ...Setting) (*Secrets, error) {
	opts := newSettings()
	for _, opt := range opt {
		opt(&opts)
	}

	client, err := newClient(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Secrets{
		cfg:    opts,
		vault:  client,
		logger: opts.Logger,
	}, nil
}

func (s *Secrets) Read(ctx context.Context, path string, v any) error {
	sv, err := s.vault.Secrets.KvV2Read(ctx, path, secretMountPath)
	if err != nil {
		if vault.IsErrorStatus(err, http.StatusNotFound) {
			// TODO convert this error here
			return nil
		}
		return fmt.Errorf("read secret: %w", err)
	}

	t := s.logger.Trace()
	if t.Enabled() {
		t.Str("path", path).
			Any("secret", sv.Data.Data).
			Msg("read secret")
	}

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

	t := s.logger.Trace()
	if t.Enabled() {
		t.Str("path", path).
			Any("secret", v).
			Msg("write secret")
	}

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
