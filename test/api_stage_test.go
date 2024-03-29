package test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/bigmikesolutions/wingman/core"
	"github.com/bigmikesolutions/wingman/core/cqrs"
	"github.com/bigmikesolutions/wingman/core/cqrs/inmemory"
	"github.com/bigmikesolutions/wingman/core/iam/identity"

	"github.com/bigmikesolutions/wingman/test/mock"

	"github.com/go-chi/chi/v5"

	wingmanHttp "github.com/bigmikesolutions/wingman/service/http"
	"github.com/stretchr/testify/require"
)

const rootPath = "/providers"

type ApiStage struct {
	server *httptest.Server
}

func NewApiStage(t *testing.T) *ApiStage {
	cqrsCfg, err := core.NewCqrsConfig()
	require.Nil(t, err, "failed to create router")

	authSvc := mock.InMemoryAuthService{Users: nil}
	cqrs := cqrs.NewCQRS(
		inmemory.NewCommandBus(cqrsCfg),
		identity.NewAuthQueryBus(
			inmemory.NewQueryBus(cqrsCfg),
			authSvc,
		),
		inmemory.NewEventBus(cqrsCfg),
	)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(mock.MockUser)
	r.Mount(rootPath, wingmanHttp.NewController(
		rootPath,
		cqrs,
	))

	return &ApiStage{
		server: httptest.NewServer(r),
	}
}

func (s *ApiStage) Close() {
	s.server.Close()
}
