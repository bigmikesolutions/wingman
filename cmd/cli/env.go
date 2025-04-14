package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/bigmikesolutions/wingman/client/graphqlclient/env"
	"github.com/bigmikesolutions/wingman/graphql/model"
)

type EnvAccessParams struct {
	DatabaseID string
}

var (
	envCmd = &cobra.Command{
		Use:   "env",
		Short: "Environment access",
	}

	envQueryCmd = &cobra.Command{
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

	envAccessParams EnvAccessParams

	envAccessCmd = &cobra.Command{
		Use:   "access",
		Short: "Request access to an environment",
		Long: `Request access to an environment (i.e. in case of incident, emergency or as a regular work-flow).
The procedure looks like following:
1) Access request is raised
2) Resources with required privileges for given request are specified
3) Manager takes a decision about given request.
4) If access is granted then environment session is generated which will be valid only for specified amount of time.
'`,

		PreRun: checkAndAuthenticate,

		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := newCtx()
			defer cancel()

			input := model.EnvGrantInput{
				Resource: []*model.ResourceGrantInput{
					{
						Env: envName,
						Database: []*model.DatabaseResource{
							{
								ID:   envAccessParams.DatabaseID,
								Info: ptr(model.AccessTypeReadOnly),
							},
						},
					},
				},
			}

			resp, err := envClient().EnvGrantMutation(ctx, input)
			if err != nil {
				log.Fatal().
					Str("env", envName).
					Any("input", input).
					Err(err).
					Msg("environment grant server error")
				return
			}

			if resp.Error != nil {
				log.Fatal().
					Any("error", resp.Error.Message).
					Str("error_code", string(resp.Error.Code)).
					Msg("environment grant client error")
				return
			}

			fmt.Printf("Environment access granted!")
		},
	}
)

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envQueryCmd)
	envCmd.AddCommand(envAccessCmd)

	envAccessCmd.Flags().StringVarP(&envAccessParams.DatabaseID, "db", "d", "", "Database name")
}

func envClient() *env.Client {
	return env.New(client())
}
