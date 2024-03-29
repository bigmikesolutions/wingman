package cqrs

import (
	"context"
)

type CommandType = string

type Command interface {
	GetType() CommandType
}

type CommandHandler interface {
	GetType() CommandType
	Handle(ctx context.Context, cmd Command) error
}

type CommandBus interface {
	ExecuteCommand(ctx context.Context, cmd Command) error
}
