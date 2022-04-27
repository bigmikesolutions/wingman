package providers

import (
	"context"

	"github.com/bigmikesolutions/wingman/pkg/provider"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"
)

func createCqrs() {
	var bus cqrs.CQRS = cqrs.CQRS{}
	bus.ExecuteCommand(context.Background(), provider.ProviderGetResourceQuery{})
}
