package server

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"

	"github.com/bigmikesolutions/wingman/graphql"
	"github.com/bigmikesolutions/wingman/graphql/directives"
	"github.com/bigmikesolutions/wingman/graphql/generated"
	"github.com/bigmikesolutions/wingman/server/httpmiddleware"
)

const (
	GraphqlEndpoint      = "/graphql"
	ProbesHealthEndpoint = "/probes/health"
	pprofEndpoint        = "/pprof"
)

type (
	HandlerWrapper = func(next http.Handler) http.Handler

	HTTPConfig struct {
		Address       string        `envconfig:"HTTP_ADDRESS" default:"0.0.0.0:8080"`
		WriteTimeout  time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"15s"`
		ReadTimeout   time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"15s"`
		ShutdownTime  time.Duration `envconfig:"HTTP_SHUTDOWN_TIME" default:"30s"`
		PprofEnabled  bool          `envconfig:"HTTP_PPROF_ENABLED" default:"false"`
		CompressLevel int           `envconfig:"HTTP_COMPRESS_LEVEL" default:"5"`
	}
)

func NewHttpHandler(cfg HTTPConfig, resolver *graphql.Resolver, chain ...HandlerWrapper) (http.Handler, error) {
	graphqlHandler := handler.New(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver,
				Directives: generated.DirectiveRoot{
					EnvSession: directives.EnvSession,
				},
			},
		),
	)

	graphqlHandler.AddTransport(transport.POST{})

	router := newHttpRouter(cfg)

	var mainHandler http.Handler = graphqlHandler
	for _, h := range chain {
		mainHandler = h(mainHandler)
	}

	router.Handle(
		GraphqlEndpoint,
		mainHandler,
	)

	return router, nil
}

func newHttpRouter(cfg HTTPConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(cfg.ReadTimeout))
	r.Use(middleware.Compress(cfg.CompressLevel))
	r.Use(httpmiddleware.RedirectProxy)
	r.Use(httpmiddleware.UserOrgAndRoles)
	r.Use(httpmiddleware.UserIdentity)
	r.Use(middleware.Heartbeat(ProbesHealthEndpoint))
	r.Use(httpmiddleware.Logger(zerolog.DebugLevel))

	if cfg.PprofEnabled {
		r.Mount(pprofEndpoint, pprofRouter())
	}

	return r
}
