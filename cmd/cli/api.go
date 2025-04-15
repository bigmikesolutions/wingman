package main

import (
	"context"
	"time"

	"github.com/bigmikesolutions/wingman/client/a10n"
	"github.com/bigmikesolutions/wingman/client/graphqlclient"
)

func client() *graphqlclient.Client {
	accessToken := getToken()
	envToken := getEnvToken()

	a10nRoundTrip := &graphqlclient.A10NRoundTrip{}
	a10nRoundTrip.SetAccessToken(accessToken.TokenType, accessToken.AccessToken)

	if envToken != nil {
		if expTime, err := a10n.GetExpirationDate(*envToken.Token); err != nil {
			logger.Debug().Err(err).Msg("GraphQL client - unable to get expiration date for env token")
		} else if expTime.After(time.Now()) {
			a10nRoundTrip.SetEnvToken(*envToken.Token)
		}
	}

	c, err := graphqlclient.New(
		cfg.HTTP.Endpoint,
		graphqlclient.WithRoundTripper(a10nRoundTrip),
	)
	if err != nil {
		logger.Panic().Err(err).Msg("GraphQL client error")
	}

	return c
}

func newCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), cfg.HTTP.Timeout)
}
