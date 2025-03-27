package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/server/env"
)

const errUniqueKeyConstraint = "duplicate key value violates unique constraint"

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

func (r *Environments) Create(ctx context.Context, e env.Environment) error {
	_, err := r.db.NamedQueryContext(ctx, createEnv, e)
	if err != nil && strings.Contains(err.Error(), errUniqueKeyConstraint) {
		return env.ErrAlreadyExists
	}
	return err
}

func (r *Environments) FindByID(ctx context.Context, orgID env.OrganisationID, id env.ID) (*env.Environment, error) {
	result := make([]env.Environment, 0)
	err := r.db.SelectContext(ctx, &result, selectEnvByID, orgID, id)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("expected one,got many environments with the same ID: %v", id)
	}
	return &result[0], nil
}
