package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bgm38",
	Short: "bgm38 is a set of api for animation",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	SilenceUsage:  true,
	SilenceErrors: true,
	// RunE: func(cmd *cobra.Command, args []string) error {
	//	return nil
	// },
}

func Execute() {
	rootCmd.AddCommand(
		versionCmd,
		serverCmd,
		spiderCmd,
		cronCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
