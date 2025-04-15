package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bigmikesolutions/wingman/client/vault"
	"github.com/bigmikesolutions/wingman/graphql/model"
)

type EnvAccessParams struct {
	DatabaseID string
}

var (
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
				logger.Fatal().
					Str("env", envName).
					Any("input", input).
					Err(err).
					Msg("environment grant server error")
				return
			}

			if resp.Error != nil {
				logger.Fatal().
					Any("error", resp.Error.Message).
					Str("error_code", string(resp.Error.Code)).
					Msg("environment grant client error")
				return
			}

			store := vault.New()
			if err := store.SetValue(secretEnvToken, resp); err != nil {
				logger.Fatal().Err(err).Msg("store: set env token failed")
			}

			fmt.Printf("Environment access granted!")
		},
	}
)

func init() {
	envCmd.AddCommand(envAccessCmd)

	envAccessCmd.Flags().StringVarP(&envAccessParams.DatabaseID, "db", "d", "", "Database name")
}
