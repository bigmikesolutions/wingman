package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/bigmikesolutions/wingman/graphql"
	"github.com/bigmikesolutions/wingman/providers"
	"github.com/bigmikesolutions/wingman/server"
	"github.com/bigmikesolutions/wingman/service/env"
	"github.com/bigmikesolutions/wingman/service/vault"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = time.RFC3339
	logger := log.Logger

	cfg, err := LoadCfg()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load config")
	}

	handler := mustHTTPHandler(logger, cfg)
	srv := http.Server{
		Addr:         cfg.HTTP.Address,
		Handler:      handler,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	logger.Info().Str("address", cfg.HTTP.Address).Msg("listening...")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal().Err(err).Msg("failed to serve HTTP")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTime)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("failed to shutdown HTTP gracefully")
	}
	logger.Info().Msg("server is down!")
}

func mustHTTPHandler(logger zerolog.Logger, cfg Config) http.Handler {
	a10n, err := newA10N(cfg.A10N)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create A10N service")
	}

	dbx, err := newDB(cfg.Database)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create DB connection")
	}

	secrets, err := vault.New(context.Background(), cfg.Vault)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create vault secrets service")
	}

	handler, err := server.NewHttpHandler(
		cfg.HTTP,
		&graphql.Resolver{
			Providers: providers.NewProviders(dbx, secrets),
			Tokens:    a10n,
		},
		env.SessionReader(a10n),
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to set-up HTTP handler")
	}
	return handler
}
