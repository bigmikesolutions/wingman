package cqrs

import "context"

type CommandType = string

type Command interface {
	GetType() CommandType
}

type CommandHandler interface {
	GetType() CommandType
	Handle(ctx context.Context, cmd Command) error
}

type CommandBus interface {
	RegisterCommand(handler CommandHandler) error
	ExecuteCommand(ctx context.Context, cmd Command) error
}

type CommandSerializer interface {
	Serialize(cmd Command) ([]byte, error)
	Deserialize(data []byte, c Command) error
}
