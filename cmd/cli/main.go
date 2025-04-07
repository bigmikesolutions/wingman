package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/bigmikesolutions/wingman/client"
	"github.com/bigmikesolutions/wingman/client/a10n"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = time.RFC3339
	logger := log.Logger

	cfg, err := LoadCfg()
	if err != nil {
		logger.Panic().Err(err).Msg("failed to load config")
	}

	cfg.A10N.ClientID = "wingman" // TODO: for debug purposes. to be removed!

	a10nDev := a10n.NewDevice(cfg.A10N.Opts()...)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token, err := a10nDev.Auth(ctx, os.Stdout)
	if err != nil {
		logger.Panic().Err(err).Msg("failed to authenticate")
	}

	log.Info().Any("token", token).Msg("authenticated")

	client.Execute()
}
