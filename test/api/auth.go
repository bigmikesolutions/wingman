package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/bigmikesolutions/wingman/server/env"
	"github.com/bigmikesolutions/wingman/server/token"
)

const (
	EnvGrantTokenDuration = 5 * time.Minute
)

type a10nRoundTripper struct {
	envGrantToken *string
}

func NewJWT() (*token.JWT, error) {
	privateKey, err := os.Open("./api/private.pem")
	if err != nil {
		return nil, fmt.Errorf("could not open private key: %v", err)
	}

	publicKey, err := os.Open("./api/public.pem")
	if err != nil {
		return nil, fmt.Errorf("could not open public key: %v", err)
	}

	return token.New(privateKey, publicKey, token.Settings{
		SigningMethod: jwt.SigningMethodRS256,
		ExpTime:       EnvGrantTokenDuration,
	})
}

func (a a10nRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if a.envGrantToken != nil {
		req.Header.Set(env.A10NHeaderEnvToken, *a.envGrantToken)
	}
	return http.DefaultTransport.RoundTrip(req)
}
