package provider

import (
	"github.com/bigmikesolutions/wingman/core/cqrs"
)

type Provider struct{}

func NewProvider() (cqrs.Config, error) {
	cfg := cqrs.NewConfig()
	//cfg.AddQueryHandlers(
	//	NewPodsQueryHandler(client),
	//)
	return cfg, nil
}
