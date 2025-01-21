package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	ConnectionInfo struct {
		ID     ID
		Driver string
		Host   string
		Name   string
		Port   int
		User   string
		Pass   string
	}

	Connection struct {
		dbID ID
		db   *sqlx.DB
		rbac RBAC
	}
)

func (c *Connection) SelectFromTable(ctx context.Context, name string, first int) (*sqlx.Rows, error) {
	roles, err := c.rbac.FindUserRolesByDatabaseID(ctx, c.dbID) // TODO use userID from ctx here
	if err != nil {
		return nil, fmt.Errorf("rbac: %w", err)
	}
	if roles == nil {
		return nil, ErrDatabaseAccessDenied
	}

	for _, role := range roles {
		for _, access := range role.DatabaseAccess {
			if !access.CanReadTable(name) {
				return nil, ErrTableAccessDenied
			}
		}
	}

	// TODO use prepared statement here to prevent sql-injection
	return c.db.QueryxContext(ctx, fmt.Sprintf("SELECT * FROM %s limit %d", name, first))
}

func connectionString(db ConnectionInfo) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		db.User, db.Pass,
		db.Host, db.Port,
		db.Name,
	)
}
