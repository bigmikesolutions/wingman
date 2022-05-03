package identity

import (
	"context"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"
)

type AuthQueryBus struct {
	queries cqrs.QueryBus
	authSvc AuthService
}

func NewAuthQueryBus(queries cqrs.QueryBus, authSvc AuthService) cqrs.QueryBus {
	return &AuthQueryBus{
		queries: queries,
		authSvc: authSvc,
	}
}

func (b AuthQueryBus) ExecuteQuery(ctx context.Context, q cqrs.Query) (interface{}, error) {
	userSession := GetUserSession(ctx)
	if userSession == nil {
		return nil, NewUnauthenticatedError("user session missing")
	}
	valid, err := b.authSvc.IsValid(ctx, *userSession)
	if err != nil {
		return nil, NewUnauthenticatedErrorDetails(err, "couldn't validate user session")
	} else if !valid {
		return nil, NewUnauthenticatedError("user session invalid")
	}
	return b.queries.ExecuteQuery(ctx, q)
}
