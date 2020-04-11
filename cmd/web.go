package cmd

import (
	"github.com/spf13/cobra"

	"bgm38/pkg/web"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run bgm38 spider",
	RunE: func(cmd *cobra.Command, args []string) error {
		return web.Start()
	},
}
