package inmemory

import (
	"context"

	"github.com/bigmikesolutions/wingman/core/cqrs"
	"github.com/pkg/errors"
)

type CommandBus struct {
	handlers map[cqrs.CommandType]cqrs.CommandHandler
}

func NewCommandBus(cfg cqrs.Config) *CommandBus {
	return &CommandBus{
		handlers: cfg.Commands,
	}
}

func (b CommandBus) ExecuteCommand(ctx context.Context, command cqrs.Command) error {
	if handler, ok := b.handlers[command.GetType()]; ok {
		if handler.GetType() != command.GetType() {
			return cqrs.NewHandlerTypeMismatchError(errors.Errorf("command handler type mismatch - expected: %s, got: %s",
				handler.GetType(), command.GetType()))
		}
		return handler.Handle(ctx, command)
	}
	return cqrs.NewMissingHandlerError(errors.Errorf("handler for command not found: %s", command.GetType()))
}
