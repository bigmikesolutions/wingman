package cqrs

import "github.com/pkg/errors"

type Config struct {
	Commands map[CommandType]CommandHandler
	Queries  map[QueryType]QueryHandler
	Events   map[EventType][]EventHandler
}

func NewConfig() Config {
	return Config{
		Commands: make(map[CommandType]CommandHandler),
		Queries:  make(map[QueryType]QueryHandler),
		Events:   make(map[EventType][]EventHandler),
	}
}

func (c *Config) GetCommands() []CommandHandler {
	handlers := make([]CommandHandler, len(c.Commands))
	idx := 0
	for _, h := range c.Commands {
		handlers[idx] = h
		idx++
	}
	return handlers
}

func (c *Config) GetQueries() []QueryHandler {
	handlers := make([]QueryHandler, len(c.Queries))
	idx := 0
	for _, h := range c.Queries {
		handlers[idx] = h
		idx++
	}
	return handlers
}

func (c *Config) GetEvents() []EventHandler {
	handlers := make([]EventHandler, 0)
	for _, h := range c.Events {
		handlers = append(handlers, h...)
	}
	return handlers
}

func (c *Config) AddCommandHandlers(handler ...CommandHandler) error {
	for _, h := range handler {
		if _, ok := c.Commands[h.GetType()]; ok {
			return errors.Errorf("command handler already exists for: %s", h.GetType())
		}
		c.Commands[h.GetType()] = h
	}
	return nil
}

func (c *Config) AddQueryHandlers(handler ...QueryHandler) error {
	for _, h := range handler {
		if _, ok := c.Queries[h.GetType()]; ok {
			return errors.Errorf("query handler already exists for: %s", h.GetType())
		}
		c.Queries[h.GetType()] = h
	}
	return nil
}

func (c *Config) AddEventHandlers(handler ...EventHandler) error {
	for _, h := range handler {
		if _, ok := c.Events[h.GetType()]; ok {
			return errors.Errorf("event handler already exists for: %s", h.GetType())
		}
		c.Events[h.GetType()] = append(c.Events[h.GetType()], h)
	}
	return nil
}

func Merge(source Config, other Config) (Config, error) {
	if err := source.AddCommandHandlers(other.GetCommands()...); err != nil {
		return Config{}, err
	}
	if err := source.AddQueryHandlers(other.GetQueries()...); err != nil {
		return Config{}, err
	}
	if err := source.AddEventHandlers(other.GetEvents()...); err != nil {
		return Config{}, err
	}
	return source, nil
}
