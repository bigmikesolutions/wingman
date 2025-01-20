package db

import (
	"context"
	"fmt"

	rbac2 "github.com/bigmikesolutions/wingman/providers/db/rbac"

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
		rbac rbac2.Repository
	}
)

func (c *Connection) SelectFromTable(ctx context.Context, name string, first int) (*sqlx.Rows, error) {
	access, err := c.rbac.Check(c.dbID, "any") // TODO use userID from ctx here
	if err != nil {
		return nil, fmt.Errorf("rbac: %w", err)
	}
	if access == nil {
		return nil, ErrDatabaseAccessDenied
	}

	if !access.CanReadTable(name) {
		return nil, ErrTableAccessDenied
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
