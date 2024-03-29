package cqrs

import (
	"context"
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
