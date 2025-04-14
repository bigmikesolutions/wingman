package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/server/a10n"
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

func (r *Resolver) reqLog(ctx context.Context) *zerolog.Logger {
	u, _ := a10n.GetIdentity(ctx)
	l := r.Logger.With().
		Str("user.user_id", u.UserID).
		Str("user.org_id", u.OrgID).
		Any("user.roles", u.Roles).
		Logger()
	return &l
}

type (
	TokenService interface {
		Create(attributes token.Values) (string, error)
	}
)
