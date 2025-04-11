package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/bigmikesolutions/wingman/client/graphqlclient/db"
)

type DBParams struct {
	Name string
}

var (
	dbParams DBParams

	dbCmd = &cobra.Command{
		Use:   "db",
		Short: "Database access",
	}

	dbInfoCmd = &cobra.Command{
		Use:   "info",
		Short: "Database information",

		PreRun: checkAndAuthenticate,

		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := newCtx()
			defer cancel()

			resp, err := dbClient().DatabaseInfoQuery(ctx, envName, dbParams.Name)
			if err != nil {
				log.Fatal().
					Str("name", dbParams.Name).
					Str("env", envName).
					Err(err).
					Msg("database info query error")
			}

			fmt.Printf("ID: %s\n", resp.Info.ID)
			fmt.Printf("Host: %s\n", resp.Info.Host)
			fmt.Printf("Port: %d\n", resp.Info.Port)
			fmt.Printf("Driver: %s\n", resp.Info.Driver)
		},
	}
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbInfoCmd)

	dbCmd.PersistentFlags().StringVarP(&dbParams.Name, "db", "n", "", "ID of the database to connect to")
}

func dbClient() *db.Client {
	return db.New(client())
}
