package cqrs

import "context"

type QueryType = string

type Query interface {
	GetType() QueryType
}

type QueryHandler interface {
	GetType() QueryType
	Handle(ctx context.Context, e Query) (interface{}, error)
}

type QueryBus interface {
	RegisterQuery(handler QueryHandler) error
	ExecuteQuery(ctx context.Context, q Query) (interface{}, error)
}
