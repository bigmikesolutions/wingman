package cqrs

import (
	"context"

	"github.com/pkg/errors"
)

type EventType = string

type Event interface {
	GetType() EventType
}

type EventHandler interface {
	GetType() EventType
	Handle(ctx context.Context, e Event) error
}

type EventBus interface {
	PublishEvent(ctx context.Context, event Event) error
}

type InMemoryEventBus struct {
	handlers map[EventType][]EventHandler
}

func NewInMemoryEventBus(cfg Config) EventBus {
	return &InMemoryEventBus{
		handlers: cfg.eventHandlers,
	}
}

func (b InMemoryEventBus) PublishEvent(ctx context.Context, event Event) error {
	if handlers, ok := b.handlers[event.GetType()]; ok {
		for _, handler := range handlers {
			if handler.GetType() != event.GetType() {
				return NewHandlerTypeMismatchError(errors.Errorf("event handler type mismatch - expected: %s, got: %s",
					handler.GetType(), event.GetType()))
			}
			if err := handler.Handle(ctx, event); err != nil {
				return err
			}
		}
	}
	return nil
}
