package providers

import (
	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/providers/db/rbac"

	"github.com/bigmikesolutions/wingman/providers/db"
	"github.com/bigmikesolutions/wingman/providers/db/repo"
	"github.com/bigmikesolutions/wingman/server/vault"
)

type Providers struct {
	DB     *db.Service
	DbRbac *rbac.Service
}

func NewProviders(dbx *sqlx.DB, secrets *vault.Secrets) *Providers {
	dbRbac := rbac.New(
		repo.NewUserRoles(dbx),
	)
	return &Providers{
		DB: db.New(
			dbRbac,
			secrets,
		),
		DbRbac: dbRbac,
	}
}

func (p *Providers) Close() error {
	return p.DB.Close()
}
