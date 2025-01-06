package test

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/bigmikesolutions/wingman/providers/db"
	"github.com/bigmikesolutions/wingman/test/containers/pg"
)

const (
	clientTimeout = 1 * time.Second
)

func testContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), clientTimeout)
}

func connectionInfo(id string, c *postgres.PostgresContainer) db.ConnectionInfo {
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
	defer cancel()

	host, err := c.Host(ctx)
	if err != nil {
		panic(err)
	}

	containerPort, err := c.MappedPort(ctx, pg.DbPort)
	if err != nil {
		panic(err)
	}

	return db.ConnectionInfo{
		ID:     id,
		Host:   host,
		Port:   containerPort.Int(),
		Driver: pg.DriverName,
		Name:   pg.DbName,
		User:   pg.DbUser,
		Pass:   pg.DbPassword,
	}
}

func ptr[T any](v T) *T {
	return &v
}
