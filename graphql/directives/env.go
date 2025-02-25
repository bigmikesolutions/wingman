package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/bigmikesolutions/wingman/service/auth"
)

func EnvSession(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	if err := auth.ValidateEnvSession(ctx); err != nil {
		return nil, &gqlerror.Error{
			Message: ctx.Err().Error(),
		}
	}

	return next(ctx)
}
