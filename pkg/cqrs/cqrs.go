package cqrs

import (
	"context"
)

type CQRS struct {
	commands CommandBus
	queries  QueryBus
	events   EventBus
}

func NewCQRS(commandBus CommandBus, queryBus QueryBus, eventBus EventBus) *CQRS {
	return &CQRS{
		commands: commandBus,
		queries:  queryBus,
		events:   eventBus,
	}
}

func (c *CQRS) ExecuteCommand(ctx context.Context, command Command) error {
	return c.commands.ExecuteCommand(ctx, command)
}

func (c *CQRS) ExecuteQuery(ctx context.Context, q Query) (interface{}, error) {
	return c.queries.ExecuteQuery(ctx, q)
}

func (c *CQRS) PublishEvent(ctx context.Context, event Event) error {
	return c.events.PublishEvent(ctx, event)
}
