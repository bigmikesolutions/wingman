package inmemory

import (
	"context"

	"github.com/bigmikesolutions/wingman/core/cqrs"
	"github.com/pkg/errors"
)

type EventBus struct {
	handlers map[cqrs.EventType][]cqrs.EventHandler
}

func NewEventBus(cfg cqrs.Config) *EventBus {
	return &EventBus{
		handlers: cfg.Events,
	}
}

func (b EventBus) PublishEvent(ctx context.Context, event cqrs.Event) error {
	if handlers, ok := b.handlers[event.GetType()]; ok {
		for _, handler := range handlers {
			if handler.GetType() != event.GetType() {
				return cqrs.NewHandlerTypeMismatchError(errors.Errorf("event handler type mismatch - expected: %s, got: %s",
					handler.GetType(), event.GetType()))
			}
			if err := handler.Handle(ctx, event); err != nil {
				return err
			}
		}
	}
	return nil
}
