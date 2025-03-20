package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/rs/zerolog"

	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/server/auth"
)

// Resolver holds the state and allows dependency injection.
type Resolver struct {
	Logger    zerolog.Logger
	Providers *providers.Providers
	Tokens    TokenService
}

type (
	TokenService interface {
		Create(attributes auth.TokenValues) (string, error)
	}
)
