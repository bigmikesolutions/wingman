package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/bigmikesolutions/wingman/graphql"
	"github.com/bigmikesolutions/wingman/providers/db"

	gql "github.com/shurcooL/graphql"

	"github.com/bigmikesolutions/wingman/service"
)

type HTTPServer struct {
	server     *httptest.Server
	client     *http.Client
	graphql    *gql.Client
	graphqlURL string
	Resolver   *graphql.Resolver
}

func New() (*HTTPServer, error) {
	cfg, err := service.LoadCfg()
	if err != nil {
		return nil, err
	}

	resolver := &graphql.Resolver{
		DB: db.New(),
	}

	handler, err := service.NewHttpHandler(cfg.HTTP, resolver)
	if err != nil {
		return nil, err
	}

	server := httptest.NewServer(handler)
	client := http.DefaultClient
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
	}, nil
}

func (s *HTTPServer) Close() {
	s.server.Close()
}
