package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/bigmikesolutions/wingman/server/a10n"
	"github.com/bigmikesolutions/wingman/server/env"
	"github.com/bigmikesolutions/wingman/server/httpmiddleware"
	"github.com/bigmikesolutions/wingman/server/token"
)

const (
	EnvGrantTokenDuration = 5 * time.Minute
)

type a10nRoundTripper struct {
	envGrantToken *string
	user          *a10n.UserIdentity
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
	if a.user != nil {
		req.Header.Set(httpmiddleware.HeaderUserID, a.user.UserID)
		req.Header.Set(httpmiddleware.HeaderUserName, a.user.UserName)
		req.Header.Set(httpmiddleware.HeaderUserEmail, a.user.Email)
		req.Header.Set(httpmiddleware.HeaderUserRoles, strings.Join(a.user.Roles, httpmiddleware.RolesSeparator))
	}
	return http.DefaultTransport.RoundTrip(req)
}
