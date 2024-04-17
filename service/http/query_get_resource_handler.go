package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bigmikesolutions/wingman/core/cqrs"
	"github.com/bigmikesolutions/wingman/core/provider"
)

type QueryGetResourceHandler struct {
	Client *http.Client
}

func (h *QueryGetResourceHandler) GetType() cqrs.QueryType {
	return provider.QueryTypeGetResource
}

func (h *QueryGetResourceHandler) Handle(ctx context.Context, in cqrs.Query) (interface{}, error) {
	query, ok := in.(provider.QueryGetResource)
	if !ok {
		return nil, fmt.Errorf("unexpected query type")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		strings.Join(query.Path, "/"),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("build http request: %w", err)
	}

	httpResp, err := h.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send http: %w", err)
	}

	defer func() {
		_ = httpResp.Body.Close()
	}()

	respBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read http response body: %w", err)
	}

	resp := make(map[string]any)
	if err := json.Unmarshal(respBytes, resp); err != nil {
		return nil, fmt.Errorf("json unmarhsal http response body: %w", err)
	}

	return resp, nil
}
