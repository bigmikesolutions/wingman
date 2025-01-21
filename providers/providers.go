package providers

import (
	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/providers/db"
	"github.com/bigmikesolutions/wingman/providers/db/repo"
)

type Providers struct {
	DB *db.Service
}

func NewProviders(dbx *sqlx.DB) *Providers {
	return &Providers{
		DB: db.New(repo.NewRBAC(dbx)),
	}
}

func (p *Providers) Close() error {
	return p.DB.Close()
}
