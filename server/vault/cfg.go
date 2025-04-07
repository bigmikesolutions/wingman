package vault

import (
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type (
	settings struct {
		Address string
		Timeout time.Duration

		// dev settings
		Token string

		// role auth
		RoleID       string
		RoleSecretID string
		RolePath     string

		// TLS settings
		ServerCert string

		// TLS with client-side cert auth
		ClientCert string
		ClientKey  string

		Logger zerolog.Logger
	}

	Setting func(config *settings)
)

func newSettings() settings {
	return settings{
		Timeout: 10 * time.Second,
		Logger:  log.Logger,
	}
}

func (c settings) vaultOptions() []vault.ClientOption {
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

func WithAddress(v string) Setting {
	return func(s *settings) {
		s.Address = v
	}
}

func WithToken(v string) Setting {
	return func(s *settings) {
		s.Token = v
	}
}

func WithLogger(v zerolog.Logger) Setting {
	return func(s *settings) {
		s.Logger = v
	}
}
