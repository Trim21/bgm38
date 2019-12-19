package cmd

import (
	"bgm38/web/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "run bgm38 server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Serve()
	},
}
