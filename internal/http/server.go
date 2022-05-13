package http

import (
	"net/http"
	"time"

	"github.com/bigmikesolutions/wingman/internal/mock"

	"github.com/bigmikesolutions/wingman/pkg/iam/identity"

	"github.com/bigmikesolutions/wingman/pkg/provider"

	"github.com/bigmikesolutions/wingman/providers/k8s"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var providers = []provider.ProviderFactory{
	k8s.NewProvider,
}

func NewRouter() (http.Handler, error) {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	cqrsCfg := cqrs.NewConfig()
	for _, provider := range providers {
		if err := provider(cqrsCfg); err != nil {
			return nil, err
		}
	}
	mock.Setup(cqrsCfg, r)
	authSvc := mock.InMemoryAuthService{Users: nil}

	cqrs := cqrs.NewCQRS(
		cqrs.NewInMemoryCommandBus(*cqrsCfg),
		identity.NewAuthQueryBus(
			cqrs.NewInMemoryQueryBus(*cqrsCfg),
			authSvc,
		),
		cqrs.NewInMemoryEventBus(*cqrsCfg),
	)

	providerCtrl := ProviderCtrl{
		path: "/providers",
		cqrs: cqrs,
	}
	r.Mount(providerCtrl.path, &providerCtrl)
	return r, nil
}
