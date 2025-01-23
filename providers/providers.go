package providers

import (
	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/providers/db"
	"github.com/bigmikesolutions/wingman/providers/db/repo"
	"github.com/bigmikesolutions/wingman/service/vault"
)

type Providers struct {
	DB *db.Service
}

func NewProviders(dbx *sqlx.DB, secrets *vault.Secrets) *Providers {
	return &Providers{
		DB: db.New(repo.NewRBAC(dbx), secrets),
	}
}

func (p *Providers) Close() error {
	return p.DB.Close()
}
