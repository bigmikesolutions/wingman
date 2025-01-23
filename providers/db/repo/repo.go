package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/providers/db"
)

type (
	RBAC struct {
		db *sqlx.DB
	}
)

func NewRBAC(db *sqlx.DB) *RBAC {
	return &RBAC{
		db: db,
	}
}

func (r *RBAC) Close() error {
	return r.db.Close()
}

func (r *RBAC) CreateUserRole(ctx context.Context, role db.UserRole) error {
	_, err := r.db.NamedQueryContext(ctx, createUserRole, role)
	return err
}

func (r *RBAC) FindUserRolesByDatabaseID(ctx context.Context, id db.ID) ([]db.UserRole, error) {
	result := make([]db.UserRole, 0)
	err := r.db.SelectContext(ctx, &result, selectUserRolesByDatabaseID, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
