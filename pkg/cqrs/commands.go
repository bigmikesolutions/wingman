package cqrs

import (
	"context"

	"github.com/pkg/errors"
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

type InMemoryCommandBus struct {
	handlers map[CommandType]CommandHandler
}

func NewInMemoryCommandBus(cfg Config) CommandBus {
	return &InMemoryCommandBus{
		handlers: cfg.commandHandlers,
	}
}

func (b InMemoryCommandBus) ExecuteCommand(ctx context.Context, command Command) error {
	if handler, ok := b.handlers[command.GetType()]; ok {
		if handler.GetType() != command.GetType() {
			return NewHandlerTypeMismatchError(errors.Errorf("command handler type mismatch - expected: %s, got: %s",
				handler.GetType(), command.GetType()))
		}
		return handler.Handle(ctx, command)
	}
	return NewMissingHandlerError(errors.Errorf("handler for command not found: %s", command.GetType()))
}
