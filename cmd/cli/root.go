package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	envName string
	logger  zerolog.Logger
	cfg     Config
)

var rootCmd = &cobra.Command{
	Use:   "wingman",
	Short: "Wingman is a CLI tool for managing access to infrastructure & server environments.",
	Long: `Wingman allows visitors (i.e. developers) to access infrastructure and server environments in a secure fashion. 
Before anyone can gain an access, environment grant (token) must be granted and approved by a manager or a supervisor.
Each grant has expiration date and limited access scopes which makes sure that a visitor can see only what's required
and needed to complete a task (i.e. solve production incident).
Each access has audit traces so generating and reviewing proof of concept for each work & task is possible.'
`,
}

func Execute(log zerolog.Logger) {
	logger = log

	c, err := LoadCfg()
	if err != nil {
		logger.Panic().Err(err).Msg("failed to load config")
	}
	cfg = c

	err = rootCmd.Execute()
	if err != nil {
		logger.Panic().Err(err).Msg("failed to execute command")
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&envName, "env", "e", "", "Environment name")
}
