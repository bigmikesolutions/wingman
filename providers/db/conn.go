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
		rbac RBAC
	}
)

func (c *Connection) SelectFromTable(ctx context.Context, name string, first int) (*sqlx.Rows, error) {
	// TODO move authorization part to a separate function
	roles, err := c.rbac.FindUserRolesByDatabaseID(ctx, c.dbID) // TODO use userID from ctx here
	if err != nil {
		return nil, fmt.Errorf("rbac: %w", err)
	}

	canRead := false
	for _, role := range roles {
		if role.CanReadTable(name) {
			canRead = true
			break
		}
	}

	if !canRead {
		return nil, ErrDatabaseAccessDenied
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
