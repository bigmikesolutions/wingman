package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database access",

	PreRun: authenticate,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
