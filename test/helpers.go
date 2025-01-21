package test

import (
	"context"
	"time"

	"github.com/bigmikesolutions/wingman/providers/db"
	"github.com/bigmikesolutions/wingman/test/containers"
)

const (
	clientTimeout = 1 * time.Second
	dbPg1         = "pg-1"
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

func databaseConnectionInfo(id string, name string, cfg containers.PostgresCfg) db.ConnectionInfo {
	info := connectionInfo(id, cfg)
	info.Name = name
	return info
}

func ptr[T any](v T) *T {
	return &v
}
