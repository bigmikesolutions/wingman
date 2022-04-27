package cqrs

import (
	"context"

	"github.com/pkg/errors"
)

type CQRS struct {
	cfg *Config
}

func NewCQRS(c Config) *CQRS {
	return &CQRS{
		cfg: &c,
	}
}

func (c *CQRS) ExecuteCommand(ctx context.Context, command Command) error {
	if handler, ok := c.cfg.commandHandlers[command.GetType()]; ok {
		if handler.GetType() != command.GetType() {
			return errors.Errorf("command handler type mismatch - expected: %s, got: %s",
				handler.GetType(), command.GetType())
		}
		return handler.Handle(ctx, command)
	}
	return errors.Errorf("handler for command not found: %s", command.GetType())
}

func (c *CQRS) ExecuteQuery(ctx context.Context, q Query) (interface{}, error) {
	if handler, ok := c.cfg.queryHandlers[q.GetType()]; ok {
		if handler.GetType() != q.GetType() {
			return nil, errors.Errorf("query handler type mismatch - expected: %s, got: %s",
				handler.GetType(), q.GetType())
		}
		return handler.Handle(ctx, q)
	}
	return nil, errors.Errorf("handler for query not found: %s", q.GetType())
}

func (c *CQRS) PublishEvent(ctx context.Context, event Event) error {
	if handlers, ok := c.cfg.eventHandlers[event.GetType()]; ok {
		for _, handler := range handlers {
			if handler.GetType() != event.GetType() {
				return errors.Errorf("event handler type mismatch - expected: %s, got: %s",
					handler.GetType(), event.GetType())
			}
			if err := handler.Handle(ctx, event); err != nil {
				return err
			}
		}
	}
	return nil
}
