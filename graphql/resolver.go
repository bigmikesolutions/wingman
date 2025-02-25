package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/service/auth"
)

// Resolver holds the state and allows dependency injection.
type Resolver struct {
	Providers *providers.Providers
	A10N      a10nService
}

type (
	a10nService interface {
		Create(attributes auth.TokenValues) (string, error)
		Validate(token string) (*auth.Token, error)
	}
)
