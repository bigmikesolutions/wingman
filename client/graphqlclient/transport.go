package graphqlclient

import (
	"fmt"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerEnvToken      = "X-Forwarded-Env-Token"
)

type A10NRoundTrip struct {
	tokenType     *string
	accessToken   *string
	envGrantToken *string
}

func (rt *A10NRoundTrip) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.envGrantToken != nil {
		req.Header.Set(headerEnvToken, *rt.envGrantToken)
	}

	if rt.tokenType != nil && rt.accessToken != nil {
		req.Header.Set(headerAuthorization, fmt.Sprintf("%s %s", *rt.tokenType, *rt.accessToken))
	}

	return http.DefaultTransport.RoundTrip(req)
}

func (rt *A10NRoundTrip) SetEnvToken(t string) {
	rt.envGrantToken = &t
}

func (rt *A10NRoundTrip) SetAccessToken(tokenType string, token string) {
	rt.tokenType = &tokenType
	rt.accessToken = &token
}
