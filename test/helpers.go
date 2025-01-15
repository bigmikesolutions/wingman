package test

import (
	"context"
	"time"

	"github.com/bigmikesolutions/wingman/test/containers"

	"github.com/bigmikesolutions/wingman/providers/db"
)

const (
	clientTimeout = 1 * time.Second
)

func testContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), clientTimeout)
}

func connectionInfo(id string, cfg containers.PostgresCfg) db.ConnectionInfo {
	return db.ConnectionInfo{
		ID:     id,
		Host:   containers.GetHost(),
		Port:   cfg.Port,
		Driver: containers.PostgresDriverName,
		Name:   cfg.Name,
		User:   cfg.User,
		Pass:   cfg.Pass,
	}
}

func ptr[T any](v T) *T {
	return &v
}
