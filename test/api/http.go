package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	gql "github.com/shurcooL/graphql"

	"github.com/bigmikesolutions/wingman/graphql"
	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/service"
	"github.com/bigmikesolutions/wingman/service/auth"
)

type HTTPServer struct {
	server     *httptest.Server
	client     *http.Client
	graphql    *gql.Client
	graphqlURL string
	Resolver   *graphql.Resolver
	providers  *providers.Providers
}

func New(prov *providers.Providers) (*HTTPServer, error) {
	cfg, err := service.LoadCfg()
	if err != nil {
		return nil, err
	}

	privateKey, err := os.Open("./api/private.pem")
	if err != nil {
		return nil, fmt.Errorf("could not open private key: %v", err)
	}

	publicKey, err := os.Open("./api/public.pem")
	if err != nil {
		return nil, fmt.Errorf("could not open public key: %v", err)
	}

	token, err := auth.New(privateKey, publicKey, auth.Settings{
		SigningMethod: jwt.SigningMethodRS256,
		ExpTime:       5 * time.Minute,
	})
	if err != nil {
		return nil, err
	}

	resolver := &graphql.Resolver{
		Providers: prov,
		Auth:      token,
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
		providers:  prov,
	}, nil
}

func (s *HTTPServer) Close() {
	s.server.Close()
	_ = s.providers.Close()
}
