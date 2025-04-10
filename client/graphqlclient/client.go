package graphqlclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	gql "github.com/shurcooL/graphql"
	"golang.org/x/net/context/ctxhttp"
)

const (
	graphqlEndpoint         = "/graphql"
	probesEndpoint          = "/probes/health"
	probesExpectedResponse  = "."
	mimeTypeApplicationJson = "application/json"
)

type (
	Client struct {
		serverURL  string
		client     *http.Client
		graphql    *gql.Client
		graphqlURL string
	}

	settings struct {
		client *http.Client
	}

	Setting func(*settings)
)

func newSettings(opt ...Setting) settings {
	s := settings{
		client: &http.Client{},
	}

	for _, o := range opt {
		o(&s)
	}

	return s
}

func WithClient(v *http.Client) Setting {
	return func(s *settings) {
		s.client = v
	}
}

func WithRoundTripper(v http.RoundTripper) Setting {
	return func(s *settings) {
		s.client.Transport = v
	}
}

func New(url string, opt ...Setting) (*Client, error) {
	opts := newSettings(opt...)

	graphqlURL := fmt.Sprintf("%s%s", url, graphqlEndpoint)

	return &Client{
		serverURL: url,
		client:    opts.client,
		graphql: gql.NewClient(
			graphqlURL,
			opts.client,
		),
		graphqlURL: graphqlURL,
	}, nil
}

func (c *Client) Query(ctx context.Context, q any, variables map[string]any) error {
	return c.graphql.Query(ctx, q, variables)
}

func (c *Client) Mutate(ctx context.Context, m any, variables map[string]any) error {
	return c.graphql.Mutate(ctx, m, variables)
}

func (c *Client) Execute(ctx context.Context, query string, vars map[string]any, out any) error {
	in := struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables,omitempty"`
	}{
		Query:     query,
		Variables: vars,
	}

	var inBuff bytes.Buffer
	err := json.NewEncoder(&inBuff).Encode(in)
	if err != nil {
		return err
	}

	resp, err := ctxhttp.Post(ctx, c.client, c.graphqlURL, mimeTypeApplicationJson, &inBuff)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(
			"non-200 OK status code: %d %s",
			resp.StatusCode,
			string(body),
		)
	}

	return json.NewDecoder(resp.Body).Decode(&out)
}

func (c *Client) Healthcheck(ctx context.Context) error {
	endpoint := fmt.Sprintf("%s%s", c.serverURL, probesEndpoint)
	resp, err := ctxhttp.Get(ctx, c.client, endpoint)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 OK status code: %d %s", resp.StatusCode, string(body))
	}

	if string(body) != probesExpectedResponse {
		return fmt.Errorf("unexpected response: %s", string(body))
	}

	return nil
}
