package provider

import (
	"github.com/bigmikesolutions/wingman/pkg/cqrs"
)

type Provider struct{}

func NewProvider() (cqrs.Config, error) {
	cfg := cqrs.NewConfig()
	//cfg.AddQueryHandlers(
	//	NewPodsQueryHandler(client),
	//)
	return cfg, nil
}
