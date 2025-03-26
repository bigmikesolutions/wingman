package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/server/env"
)

type (
	Environments struct {
		db *sqlx.DB
	}
)

func NewEnvironments(db *sqlx.DB) *Environments {
	return &Environments{
		db: db,
	}
}

func (r *Environments) Close() error {
	return r.db.Close()
}

func (r *Environments) Create(ctx context.Context, env env.Environment) error {
	_, err := r.db.NamedQueryContext(ctx, createEnv, env)
	return err
}

func (r *Environments) FindByID(ctx context.Context, id env.ID) (*env.Environment, error) {
	result := make([]env.Environment, 0)
	err := r.db.SelectContext(ctx, &result, selectEnvByID, id)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, env.ErrNotFound
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("expected one,got many environments with the same ID: %v", id)
	}
	return &result[0], nil
}
