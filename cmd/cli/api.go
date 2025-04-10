package main

import (
	"context"

	"github.com/bigmikesolutions/wingman/client/graphqlclient"
)

func client() *graphqlclient.Client {
	token := getToken()

	a10nRoundTrip := &graphqlclient.A10NRoundTrip{}
	a10nRoundTrip.SetAccessToken(token.TokenType, token.AccessToken)

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
