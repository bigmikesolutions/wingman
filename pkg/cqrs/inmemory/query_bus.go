package inmemory

import (
	"context"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"
	"github.com/pkg/errors"
)

type QueryBus struct {
	handlers map[cqrs.QueryType]cqrs.QueryHandler
}

func NewQueryBus(cfg cqrs.Config) *QueryBus {
	return &QueryBus{
		handlers: cfg.Queries,
	}
}

func (b QueryBus) ExecuteQuery(ctx context.Context, q cqrs.Query) (interface{}, error) {
	if handler, ok := b.handlers[q.GetType()]; ok {
		if handler.GetType() != q.GetType() {
			return nil, cqrs.NewHandlerTypeMismatchError(errors.Errorf("query handler type mismatch - expected: %s, got: %s",
				handler.GetType(), q.GetType()))
		}
		return handler.Handle(ctx, q)
	}
	return nil, cqrs.NewMissingHandlerError(errors.Errorf("handler for query not found: %s", q.GetType()))
}
