package main

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/bigmikesolutions/wingman/client/graphqlclient/env"
	"github.com/bigmikesolutions/wingman/client/vault"
	"github.com/bigmikesolutions/wingman/graphql/model"
)

const (
	secretEnvToken = "wingman/env_token"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Environment access",
}

func init() {
	rootCmd.AddCommand(envCmd)
}

func envClient() *env.Client {
	return env.New(client())
}

func getEnvToken() *model.EnvGrantPayload {
	store := vault.New()
	var token model.EnvGrantPayload
	if err := store.GetValue(secretEnvToken, &token); err != nil {
		if !errors.Is(err, vault.ErrNotFound) {
			logger.Fatal().Err(err).Msg("store: get env token failed")
		}
		return nil
	}
	return &token
}
