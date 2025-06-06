package main

import (
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Send wingman healthcheck",

	PreRun: checkAndAuthenticate,

	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := newCtx()
		defer cancel()

		if err := client().Healthcheck(ctx); err != nil {
			logger.Fatal().
				Str("env", envName).
				Err(err).
				Msg("ping failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
