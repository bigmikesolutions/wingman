package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/rs/zerolog"

	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/server/env"
	"github.com/bigmikesolutions/wingman/server/token"
)

// Resolver holds the state and allows dependency injection.
type Resolver struct {
	Logger       zerolog.Logger
	Providers    *providers.Providers
	Tokens       TokenService
	Environments *env.Service
}

type (
	TokenService interface {
		Create(attributes token.Values) (string, error)
	}
)
