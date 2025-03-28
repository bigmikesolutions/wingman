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

	settings struct {
		ReadTimeout   time.Duration
		CompressLevel int
		PprofEnabled  bool
		LogLevel      zerolog.Level
	}

	Setting func(*settings)
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

func NewHTTPRouter(validator tokenValidator, opts ...Setting) *chi.Mux {
	settings := newSettings()
	for _, opt := range opts {
		opt(&settings)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(settings.ReadTimeout))
	r.Use(middleware.Compress(settings.CompressLevel))
	r.Use(httpmiddleware.RedirectProxy)
	r.Use(httpmiddleware.UserOrgAndRoles)
	r.Use(httpmiddleware.UserIdentity)
	r.Use(httpmiddleware.SessionReader(validator))
	r.Use(middleware.Heartbeat(ProbesHealthEndpoint))
	r.Use(httpmiddleware.Logger(settings.LogLevel))

	if settings.PprofEnabled {
		r.Mount(pprofEndpoint, pprofRouter())
	}

	return r
}

func newSettings() settings {
	return settings{
		ReadTimeout:   15 * time.Second,
		CompressLevel: 5,
		PprofEnabled:  false,
		LogLevel:      zerolog.DebugLevel,
	}
}

func WithReadTimeout(v time.Duration) Setting {
	return func(s *settings) {
		s.ReadTimeout = v
	}
}

func WithCompressLevel(v int) Setting {
	return func(s *settings) {
		s.CompressLevel = v
	}
}

func WithPPROF(v bool) Setting {
	return func(s *settings) {
		s.PprofEnabled = v
	}
}

func WithLogLevel(v zerolog.Level) Setting {
	return func(s *settings) {
		s.LogLevel = v
	}
}
