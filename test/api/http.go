package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	gql "github.com/shurcooL/graphql"

	"github.com/bigmikesolutions/wingman/graphql"
	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/service"
	"github.com/bigmikesolutions/wingman/service/env"
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

func New(prov *providers.Providers) (*HTTPServer, error) {
	cfg, err := service.LoadCfg()
	if err != nil {
		return nil, err
	}

	token, err := NewJWT()
	if err != nil {
		return nil, err
	}

	resolver := &graphql.Resolver{
		Providers: prov,
		Tokens:    token,
	}

	handler, err := service.NewHttpHandler(cfg.HTTP, resolver, env.SessionReader(token))
	if err != nil {
		return nil, err
	}

	server := httptest.NewServer(handler)
	rt := &a10nRoundTripper{}
	client := &http.Client{
		Transport: rt,
	}
	graphqlURL := fmt.Sprintf("%s%s", server.URL, service.GraphqlEndpoint)

	return &HTTPServer{
		server: server,
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
	_ = s.providers.Close()
}

func (s *HTTPServer) SetEnvToken(t string) {
	s.rt.envGrantToken = &t
}
