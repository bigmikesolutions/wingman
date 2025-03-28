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
	"github.com/bigmikesolutions/wingman/server/token"
)

const (
	GraphqlEndpoint      = "/graphql"
	ProbesHealthEndpoint = "/probes/health"
	pprofEndpoint        = "/pprof"
)

type (
	router interface {
		Handle(pattern string, handler http.Handler)
	}

	tokenValidator interface {
		Validate(token string) (*token.Token, error)
	}

	HandlerWrapper = func(next http.Handler) http.Handler

	HTTPSettings struct {
		ReadTimeout   time.Duration
		CompressLevel int
		PprofEnabled  bool
		LogLevel      zerolog.Level
	}

	HTTPSetting func(*HTTPSettings)
)

func SetGraphQLHandler(router router, resolver *graphql.Resolver) {
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

	router.Handle(
		GraphqlEndpoint,
		graphqlHandler,
	)
}

func NewHTTPRouter(validator tokenValidator, settings ...HTTPSetting) *chi.Mux {
	opts := newHTTPConfig()
	for _, opt := range settings {
		opt(&opts)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(opts.ReadTimeout))
	r.Use(middleware.Compress(opts.CompressLevel))
	r.Use(httpmiddleware.RedirectProxy)
	r.Use(httpmiddleware.UserOrgAndRoles)
	r.Use(httpmiddleware.UserIdentity)
	r.Use(httpmiddleware.SessionReader(validator))
	r.Use(middleware.Heartbeat(ProbesHealthEndpoint))
	r.Use(httpmiddleware.Logger(opts.LogLevel))

	if opts.PprofEnabled {
		r.Mount(pprofEndpoint, pprofRouter())
	}

	return r
}

func newHTTPConfig() HTTPSettings {
	return HTTPSettings{
		ReadTimeout:   15 * time.Second,
		CompressLevel: 5,
		PprofEnabled:  false,
		LogLevel:      zerolog.DebugLevel,
	}
}

func WithReadTimeout(v time.Duration) HTTPSetting {
	return func(s *HTTPSettings) {
		s.ReadTimeout = v
	}
}

func WithCompressLevel(v int) HTTPSetting {
	return func(s *HTTPSettings) {
		s.CompressLevel = v
	}
}

func WithPPROF(v bool) HTTPSetting {
	return func(s *HTTPSettings) {
		s.PprofEnabled = v
	}
}

func WithLogLevel(v zerolog.Level) HTTPSetting {
	return func(s *HTTPSettings) {
		s.LogLevel = v
	}
}
