package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/bigmikesolutions/wingman/providers/db"
)

// Resolver holds the state and allows dependency injection.
type Resolver struct {
	DB *db.Service
}
