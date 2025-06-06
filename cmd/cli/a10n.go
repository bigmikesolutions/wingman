package main

import (
	"context"
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/bigmikesolutions/wingman/client/a10n"
	"github.com/bigmikesolutions/wingman/client/vault"
)

const (
	secretAccessToken = "wingman/access_token"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate device",

	Run: func(cmd *cobra.Command, args []string) {
		authenticate(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func getToken() a10n.TokenResponse {
	store := vault.New()
	var token a10n.TokenResponse
	if err := store.GetValue(secretAccessToken, &token); err != nil {
		if !errors.Is(err, vault.ErrNotFound) {
			logger.Fatal().Err(err).Msg("store: get access token failed")
		}
	}
	return token
}

func checkAndAuthenticate(cmd *cobra.Command, args []string) {
	token := getToken()
	if token.HasExpired() {
		authenticate(cmd, args)
	}
}

func authenticate(cmd *cobra.Command, args []string) {
	cfg.A10N.ClientID = "wingman" // TODO: for debug purposes. to be removed!

	device := a10n.NewDevice(cfg.A10N.Opts()...)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.A10N.MaxTime)
	defer cancel()

	t, err := device.Auth(ctx, os.Stdout)
	if err != nil {
		logger.Panic().Err(err).Msg("failed to authenticate")
	}

	log.Debug().Msg("authenticated")

	store := vault.New()
	if err := store.SetValue(secretAccessToken, t); err != nil {
		logger.Fatal().Err(err).Msg("store: set access token failed")
	}
}
