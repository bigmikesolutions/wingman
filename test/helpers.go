package test

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/test/containers"
)

const (
	clientTimeout = 5 * time.Second

	testUserID = "admin"
	testOrg    = "bms"
)

func testContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), clientTimeout)
}

func cleanTables(dbx *sqlx.DB, tables []string) {
	for _, table := range tables {
		ctx, cancel := testContext()
		if _, err := dbx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s", table)); err != nil {
			panic(fmt.Errorf("clean table %s: %w", table, err))
		}
		cancel()
	}
}

func newTestDatabase(dc *containers.Service) containers.PostgresCfg {
	ctx, cancel := testContext()
	defer cancel()

	dbx, err := dc.DB(ctx)
	if err != nil {
		panic(fmt.Errorf("db conn: %w", err))
	}
	defer func() {
		_ = dbx.Close()
	}()

	dbCfg := dc.Config().Postgres
	dbCfg.Name = uuid.New().String()

	_, err = dbx.ExecContext(ctx, fmt.Sprintf(`CREATE DATABASE "%s"`, dbCfg.Name))
	if err != nil {
		panic(fmt.Errorf("create db: %w", err))
	}
	_, err = dbx.ExecContext(ctx, fmt.Sprintf(`GRANT ALL PRIVILEGES ON DATABASE "%s" TO %s;`, dbCfg.Name, dbCfg.User))
	if err != nil {
		panic(fmt.Errorf("create db: %w", err))
	}

	return dbCfg
}

func ptr[T any](v T) *T {
	return &v
}
