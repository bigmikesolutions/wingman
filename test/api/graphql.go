package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

const mimeTypeApplicationJson = "application/json"

func (s *HTTPServer) graphqlExecute(ctx context.Context, query string, vars map[string]any, out any) error {
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

	resp, err := ctxhttp.Post(ctx, s.client, s.graphqlURL, mimeTypeApplicationJson, &inBuff)
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
