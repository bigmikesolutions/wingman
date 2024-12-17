package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/shurcooL/graphql"

	"github.com/bigmikesolutions/wingman/service"
)

type HTTPServer struct {
	server     *httptest.Server
	client     *http.Client
	graphql    *graphql.Client
	graphqlURL string
}

func New() (*HTTPServer, error) {
	cfg, err := service.LoadCfg()
	if err != nil {
		return nil, err
	}

	handler, err := service.NewHttpHandler(cfg.HTTP)
	if err != nil {
		return nil, err
	}

	server := httptest.NewServer(handler)
	client := http.DefaultClient
	graphqlURL := fmt.Sprintf("%s%s", server.URL, service.GraphqlEndpoint)

	return &HTTPServer{
		server: server,
		client: client,
		graphql: graphql.NewClient(
			graphqlURL,
			client,
		),
		graphqlURL: graphqlURL,
	}, nil
}

func (s *HTTPServer) Close() {
	s.server.Close()
}
