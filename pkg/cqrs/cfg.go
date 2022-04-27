package cqrs

import "github.com/pkg/errors"

type Config struct {
	commandHandlers map[CommandType]CommandHandler
	queryHandlers   map[QueryType]QueryHandler
	eventHandlers   map[EventType][]EventHandler
}

func NewConfig() *Config {
	return &Config{
		commandHandlers: make(map[CommandType]CommandHandler),
		queryHandlers:   make(map[QueryType]QueryHandler),
		eventHandlers:   make(map[EventType][]EventHandler),
	}
}

func (c *Config) AddCommandHandlers(handler ...CommandHandler) error {
	for _, h := range handler {
		if _, ok := c.commandHandlers[h.GetType()]; ok {
			return errors.Errorf("command handler already exists for: %s", h.GetType())
		}
		c.commandHandlers[h.GetType()] = h
	}
	return nil
}

func (c *Config) AddQueryHandlers(handler ...QueryHandler) error {
	for _, h := range handler {
		if _, ok := c.queryHandlers[h.GetType()]; ok {
			return errors.Errorf("query handler already exists for: %s", h.GetType())
		}
		c.queryHandlers[h.GetType()] = h
	}
	return nil
}
