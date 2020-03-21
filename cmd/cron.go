package cmd

import (
	"github.com/spf13/cobra"

	"bgm38/pkg/cron"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "run bgm38 cron jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			if args[0] == "run" {
				return cron.Run(args[1])
			}
		}
		return cron.Start()
	},
}
