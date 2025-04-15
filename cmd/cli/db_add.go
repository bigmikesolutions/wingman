package main

import (
	"github.com/spf13/cobra"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

var dbAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add database information",

	PreRun: checkAndAuthenticate,

	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := newCtx()
		defer cancel()

		input := model.AddDatabaseInput{
			Env:      envName,
			ID:       "test-db-1",
			Name:     "test-db-1",
			Host:     "db-host",
			Port:     3306,
			User:     "db-user",
			Password: "db-password",
			Driver:   model.DriverTypePostgres,
		}

		resp, err := dbClient().AddDatabaseMutation(ctx, input)
		if err != nil {
			logger.Fatal().
				Str("env", envName).
				Any("input", input).
				Err(err).
				Msg("add database mutation server error")
		}

		if resp.Error != nil {
			logger.Fatal().
				Str("env", envName).
				Any("input", input).
				Any("resp", resp).
				Err(err).
				Msg("add database mutation client error")
		}
	},
}

func init() {
	dbCmd.AddCommand(dbAddCmd)
}
