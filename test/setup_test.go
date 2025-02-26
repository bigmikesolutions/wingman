package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/service/vault"
	"github.com/bigmikesolutions/wingman/test/containers"
)

const (
	dockerSetupTimeout = 120 * time.Second
)

var dc *containers.Service

func TestMain(m *testing.M) {
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

func newProviders() *providers.Providers {
	db, err := dc.DB(context.Background())
	if err != nil {
		panic(fmt.Errorf("could not connect to db: %w", err))
	}

	vaultCfg := dc.Config().Vault
	v, err := vault.New(context.Background(), vault.Config{
		Address: fmt.Sprintf("http://localhost:%d", vaultCfg.Port),
		Token:   vaultCfg.RootToken,
	})
	if err != nil {
		panic(fmt.Errorf("could not connect to vault: %w", err))
	}

	return providers.NewProviders(db, v)
}

func mustDB() *sqlx.DB {
	dbx, err := dc.DB(context.Background())
	if err != nil {
		panic(fmt.Errorf("could not connect to db: %w", err))
	}
	return dbx
}
