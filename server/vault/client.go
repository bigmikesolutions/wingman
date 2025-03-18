package vault

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func newClient(ctx context.Context, cfg Config) (*vault.Client, error) {
	client, err := vault.New(cfg.vaultOptions()...)
	if err != nil {
		return nil, fmt.Errorf("vault client: %w", err)
	}
	_ = client.SetRequestCallbacks(func(req *http.Request) {
		// TODO make this debug only
		log.Println("request:", *req)
	})
	_ = client.SetResponseCallbacks(func(req *http.Request, resp *http.Response) {
		// TODO make this debug only
		log.Println("response:", *resp)
	})

	token, err := vaultAuth(ctx, client, cfg)
	if err != nil {
		return nil, fmt.Errorf("vault auth: %w", err)
	}

	if err := client.SetToken(token); err != nil {
		return nil, fmt.Errorf("vault token: %w", err)
	}

	return client, nil
}

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
