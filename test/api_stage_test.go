package test

import (
	"net/http/httptest"
	"testing"

	wingman "github.com/bigmikesolutions/wingman/internal/http"
	"github.com/stretchr/testify/require"
)

type ApiStage struct {
	server *httptest.Server
}

func NewApiStage(t *testing.T) *ApiStage {
	handler, err := wingman.NewRouter()
	require.Nil(t, err, "failed to create router")
	svr := httptest.NewServer(handler)
	return &ApiStage{
		server: svr,
	}
}

func (s *ApiStage) Close() {
	s.server.Close()
}
