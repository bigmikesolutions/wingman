package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"

	"github.com/bigmikesolutions/wingman/service/auth"
)

const (
	A10NHeaderEnvToken = "X-A10N-Env-Token"
)

type a10nRoundTripper struct {
	envGrantToken *string
}

func NewJWT() (*auth.JWT, error) {
	privateKey, err := os.Open("./api/private.pem")
	if err != nil {
		return nil, fmt.Errorf("could not open private key: %v", err)
	}

	publicKey, err := os.Open("./api/public.pem")
	if err != nil {
		return nil, fmt.Errorf("could not open public key: %v", err)
	}

	return auth.New(privateKey, publicKey, auth.Settings{
		SigningMethod: jwt.SigningMethodRS256,
		ExpTime:       EnvGrantTokenDuration,
	})
}

func (a a10nRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if a.envGrantToken != nil {
		req.Header.Set(A10NHeaderEnvToken, *a.envGrantToken)
	}
	return http.DefaultTransport.RoundTrip(req)
}
