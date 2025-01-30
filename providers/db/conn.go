package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	ID = string

	ConnectionInfo struct {
		ID     ID     `json:"id"`
		Driver string `json:"driver"`
		Host   string `json:"host"`
		Name   string `json:"name"`
		Port   int    `json:"port"`
		User   string `json:"user"`
		Pass   string `json:"pass"`
	}

	Connection struct {
		dbID ID
		db   *sqlx.DB
		rbac rbac
	}
)

func (c *Connection) SelectFromTable(ctx context.Context, name string, first int) (*sqlx.Rows, error) {
	if err := c.rbac.ReadTable(ctx, c.dbID, name); err != nil {
		return nil, err
	}

	// TODO use prepared statement here to prevent sql-injection
	return c.db.QueryxContext(ctx, fmt.Sprintf("SELECT * FROM %s LIMIT %d", name, first))
}

func connectionString(db ConnectionInfo) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		db.User, db.Pass,
		db.Host, db.Port,
		db.Name,
	)
}
