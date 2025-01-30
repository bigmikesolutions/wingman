package repo

import (
	"context"

	"github.com/bigmikesolutions/wingman/providers/db/rbac"

	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/providers/db"
)

type (
	UserRoles struct {
		db *sqlx.DB
	}
)

func NewUserRoles(db *sqlx.DB) *UserRoles {
	return &UserRoles{
		db: db,
	}
}

func (r *UserRoles) Close() error {
	return r.db.Close()
}

func (r *UserRoles) CreateUserRole(ctx context.Context, role rbac.UserRole) error {
	_, err := r.db.NamedQueryContext(ctx, createUserRole, role)
	return err
}

func (r *UserRoles) FindUserRolesByDatabaseID(ctx context.Context, id db.ID) ([]rbac.UserRole, error) {
	result := make([]rbac.UserRole, 0)
	err := r.db.SelectContext(ctx, &result, selectUserRolesByDatabaseID, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
