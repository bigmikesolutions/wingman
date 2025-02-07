package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/bigmikesolutions/wingman/providers"
)

// Resolver holds the state and allows dependency injection.
type Resolver struct {
	Providers *providers.Providers
	Auth      Auth
}

type (
	Auth interface {
		Create(attributes map[string]any) (string, error)
		Validate(token string) (map[string]any, error)
	}
)
