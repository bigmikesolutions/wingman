package secrets

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func vaultAuth(ctx context.Context, client *vault.Client, cfg Config) (string, error) {
	if cfg.ClientCert != "" && cfg.ClientKey != "" {
		resp, err := client.Auth.CertLogin(ctx, schema.CertLoginRequest{
			Name: "my-cert",
		})
		if err != nil {
			return "", fmt.Errorf("vault client-side cert auth: %w", err)
		}
		return resp.Auth.ClientToken, nil
	}

	if cfg.RoleID != "" && cfg.RoleSecretID != "" {
		resp, err := client.Auth.AppRoleLogin(
			ctx,
			schema.AppRoleLoginRequest{
				RoleId:   cfg.RoleID,
				SecretId: cfg.RoleSecretID,
			},
			vault.WithMountPath(cfg.RolePath),
		)
		if err != nil {
			return "", fmt.Errorf("vault role auth: %w", err)
		}
		return resp.Auth.ClientToken, nil
	}

	return cfg.Token, nil
}
