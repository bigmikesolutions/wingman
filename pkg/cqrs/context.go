package cqrs

import "context"

type ctx string

const (
	contextQueryBus ctx = "cqrs_query_bus"
)

func SetQueryBus(ctx context.Context, bus QueryBus) context.Context {
	return context.WithValue(ctx, contextQueryBus, bus)
}

func GetQueryBus(ctx context.Context) QueryBus {
	if ctx == nil {
		return nil
	}
	if bus, ok := ctx.Value(contextQueryBus).(QueryBus); ok {
		return bus
	}
	return nil
}
