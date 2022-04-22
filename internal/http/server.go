package http

import (
	"net/http"
	"time"

	"github.com/bigmikesolutions/wingman/providers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() (http.Handler, error) {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	providerCtrl := ProviderCtrl{
		path:      "/providers",
		providers: providers.NewProviders(),
	}
	r.Mount(providerCtrl.path, &providerCtrl)
	return r, nil
}
