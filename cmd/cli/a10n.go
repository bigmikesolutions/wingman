package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/bigmikesolutions/wingman/client/a10n"
)

var token *a10n.TokenResponse

func authenticate(cmd *cobra.Command, args []string) {
	if token == nil {
		cfg.A10N.ClientID = "wingman" // TODO: for debug purposes. to be removed!

		a10nDev := a10n.NewDevice(cfg.A10N.Opts()...)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		t, err := a10nDev.Auth(ctx, os.Stdout)
		if err != nil {
			logger.Panic().Err(err).Msg("failed to authenticate")
		}

		log.Info().Any("token", t).Msg("authenticated")
		token = &t
	}
}
