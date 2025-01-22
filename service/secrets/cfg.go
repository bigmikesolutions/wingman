package secrets

import (
	"time"

	"github.com/hashicorp/vault-client-go"
)

type Config struct {
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

func (c Config) vaultOptions() []vault.ClientOption {
	opts := make([]vault.ClientOption, 0)
	opts = append(opts, vault.WithAddress(c.Address))
	opts = append(opts, vault.WithRequestTimeout(c.Timeout))

	if c.ServerCert != "" {
		tls := vault.TLSConfiguration{}
		tls.ServerCertificate.FromFile = c.ServerCert

		if c.ClientCert != "" {
			tls.ClientCertificate.FromFile = c.ClientCert
		}
		if c.ClientKey != "" {
			tls.ClientCertificateKey.FromFile = c.ClientKey
		}

		opts = append(opts, vault.WithTLS(tls))
	}

	return opts
}
