package api

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bigmikesolutions/wingman/server"
)

func AssertHeartbeat(t *testing.T, s *HTTPServer) {
	resp, err := s.client.Get(fmt.Sprintf("%s%s", s.server.URL, server.ProbesHealthEndpoint))
	require.NoError(t, err, "heartbeat")

	defer func() {
		_ = resp.Body.Close()
	}()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "heartbeat status code")

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "read heartbeat response")
	assert.Equal(t, ".", string(respBody), "heartbeat response")
}
