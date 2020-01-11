package cmd

import (
	"bgm38/pkg/cron"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cronCmd)
}

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "run bgm38 cron jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cron.Start()
	},
}
