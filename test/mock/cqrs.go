package mock

import (
	"github.com/bigmikesolutions/wingman/core/cqrs"
	"github.com/go-chi/chi/v5"
)

func Setup(cfg *cqrs.Config, router chi.Router) error {
	cfg.AddQueryHandlers(
		&QueryGetUserAccessPolicyHandler{},
	)
	router.Use(MockUser)
	return nil
}
