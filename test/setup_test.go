package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/server/vault"
	"github.com/bigmikesolutions/wingman/test/containers"
)

const (
	dockerSetupTimeout = 120 * time.Second
)

var dc *containers.Service

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.RFC3339

	ctx, cancel := context.WithTimeout(context.Background(), dockerSetupTimeout)
	defer cancel()

	var err error
	dc, err = containers.New(ctx)
	if err != nil {
		panic(err)
	}
	defer dc.Close()

	code := m.Run()
	defer os.Exit(code)
}

func newProviders(dbx *sqlx.DB) *providers.Providers {
	vaultCfg := dc.Config().Vault
	v, err := vault.New(
		context.Background(),
		vault.WithAddress(fmt.Sprintf("http://localhost:%d", vaultCfg.Port)),
		vault.WithToken(vaultCfg.RootToken),
	)
	if err != nil {
		panic(fmt.Errorf("could not connect to vault: %w", err))
	}

	return providers.NewProviders(dbx, v)
}

func mustDB() *sqlx.DB {
	dbx, err := dc.DB(context.Background())
	if err != nil {
		panic(fmt.Errorf("could not connect to db: %w", err))
	}
	return dbx
}
