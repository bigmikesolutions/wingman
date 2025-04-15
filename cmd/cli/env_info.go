package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var envQueryCmd = &cobra.Command{
	Use:   "info",
	Short: "Environment information",

	PreRun: checkAndAuthenticate,

	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := newCtx()
		defer cancel()

		resp, err := envClient().EnvironmentQuery(ctx, envName)
		if err != nil {
			log.Fatal().
				Str("name", dbParams.Name).
				Str("env", envName).
				Err(err).
				Msg("database info query error")
		}

		fmt.Printf("ID: %s\n", resp.ID)
		fmt.Printf("Description: %s\n", resp.Description)
		fmt.Printf("Created at: %s\n", resp.CreatedAt)
		if resp.ModifiedAt != nil {
			fmt.Printf("Modified at: %s\n", resp.ModifiedAt)
		}
	},
}

func init() {
	envCmd.AddCommand(envQueryCmd)
}
