package main

import (
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
)

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.PersistentFlags().StringVarP(&dbParams.Name, "db", "n", "", "ID of the database to connect to")
}

func dbClient() *db.Client {
	return db.New(client())
}
