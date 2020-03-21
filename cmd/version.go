package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"bgm38/config"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of app",
	// Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bgm38 server start ", config.Version)
	},
}
