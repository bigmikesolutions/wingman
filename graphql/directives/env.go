package directives

import (
	"context"

	"github.com/bigmikesolutions/wingman/service/env"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func EnvSession(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	if err := env.ValidateSession(ctx); err != nil {
		return nil, &gqlerror.Error{
			Message: ctx.Err().Error(),
		}
	}

	return next(ctx)
}
