package http

import (
	"net/http"
	"time"

	mock2 "github.com/bigmikesolutions/wingman/test/mock"

	"github.com/bigmikesolutions/wingman/pkg/cqrs/inmemory"

	"github.com/bigmikesolutions/wingman/pkg/iam/identity"

	"github.com/bigmikesolutions/wingman/pkg/provider"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(providers ...provider.Factory) (http.Handler, error) {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	cqrsCfg := cqrs.NewConfig()
	for _, prov := range providers {
		provCfg, err := prov()
		if err != nil {
			return nil, err
		}

		cqrsCfg, err = cqrs.Merge(cqrsCfg, provCfg)
		if err != nil {
			return nil, err
		}
	}

	mock2.Setup(&cqrsCfg, r)
	authSvc := mock2.InMemoryAuthService{Users: nil}

	cqrs := cqrs.NewCQRS(
		inmemory.NewCommandBus(cqrsCfg),
		identity.NewAuthQueryBus(
			inmemory.NewQueryBus(cqrsCfg),
			authSvc,
		),
		inmemory.NewEventBus(cqrsCfg),
	)

	providerCtrl := ProviderCtrl{
		path: "/providers",
		cqrs: cqrs,
	}
	r.Mount(providerCtrl.path, &providerCtrl)
	return r, nil
}
