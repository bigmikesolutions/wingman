package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/jmoiron/sqlx"

	"github.com/bigmikesolutions/wingman/server/env/repo"

	"github.com/bigmikesolutions/wingman/server/a10n"

	gql "github.com/shurcooL/graphql"

	"github.com/bigmikesolutions/wingman/graphql"
	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/server"
	"github.com/bigmikesolutions/wingman/server/env"
)

type HTTPServer struct {
	server     *httptest.Server
	client     *http.Client
	graphql    *gql.Client
	graphqlURL string
	Resolver   *graphql.Resolver
	providers  *providers.Providers
	rt         *a10nRoundTripper
}

func New(dbx *sqlx.DB, prov *providers.Providers) (*HTTPServer, error) {
	cfg, err := server.LoadCfg()
	if err != nil {
		return nil, err
	}

	token, err := NewJWT()
	if err != nil {
		return nil, err
	}

	resolver := &graphql.Resolver{
		Providers:    prov,
		Tokens:       token,
		Environments: env.New(repo.NewEnvironments(dbx)),
	}

	handler, err := server.NewHttpHandler(cfg.HTTP, resolver, env.SessionReader(token))
	if err != nil {
		return nil, err
	}

	svc := httptest.NewServer(handler)
	rt := &a10nRoundTripper{}
	client := &http.Client{
		Transport: rt,
	}
	graphqlURL := fmt.Sprintf("%s%s", svc.URL, server.GraphqlEndpoint)

	return &HTTPServer{
		server: svc,
		client: client,
		graphql: gql.NewClient(
			graphqlURL,
			client,
		),
		graphqlURL: graphqlURL,
		Resolver:   resolver,
		providers:  prov,
		rt:         rt,
	}, nil
}

func (s *HTTPServer) Close() {
	s.server.Close()
}

func (s *HTTPServer) SetEnvToken(t string) {
	s.rt.envGrantToken = &t
}

func (s *HTTPServer) SetUser(u a10n.UserIdentity) {
	s.rt.user = &u
}
