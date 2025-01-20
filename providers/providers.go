package providers

import (
	"github.com/bigmikesolutions/wingman/providers/db"
	"github.com/bigmikesolutions/wingman/providers/db/repo"
)

type Providers struct {
	DB *db.Service
}

func NewProviders() *Providers {
	return &Providers{
		DB: db.New(repo.NewRBAC(nil)),
	}
}
