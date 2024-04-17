package service

import (
	"net/http"

	"github.com/bigmikesolutions/wingman/core/cqrs"
	servicehttp "github.com/bigmikesolutions/wingman/service/http"
)

func NewHttpCqrs() (cqrs.Config, error) {
	cfg := cqrs.NewConfig()

	if err := cfg.AddQueryHandlers(
		&servicehttp.QueryGetResourceHandler{
			Client: http.DefaultClient,
		},
	); err != nil {
		return cfg, err
	}

	return cfg, nil
}
