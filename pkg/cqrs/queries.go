package cqrs

import (
	"context"

	"github.com/pkg/errors"
)

type QueryType = string

type Query interface {
	GetType() QueryType
}

type QueryHandler interface {
	GetType() QueryType
	Handle(ctx context.Context, e Query) (interface{}, error)
}

type QueryBus interface {
	ExecuteQuery(ctx context.Context, q Query) (interface{}, error)
}

type InMemoryQueryBus struct {
	handlers map[QueryType]QueryHandler
}

func NewInMemoryQueryBus(cfg Config) QueryBus {
	return &InMemoryQueryBus{
		handlers: cfg.queryHandlers,
	}
}

func (b InMemoryQueryBus) ExecuteQuery(ctx context.Context, q Query) (interface{}, error) {
	if handler, ok := b.handlers[q.GetType()]; ok {
		if handler.GetType() != q.GetType() {
			return nil, NewHandlerTypeMismatchError(errors.Errorf("query handler type mismatch - expected: %s, got: %s",
				handler.GetType(), q.GetType()))
		}
		return handler.Handle(ctx, q)
	}
	return nil, NewMissingHandlerError(errors.Errorf("handler for query not found: %s", q.GetType()))
}
