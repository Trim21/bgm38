package cmd

import (
	"bgm38/pkg/spider"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(spiderCmd)
}

var spiderCmd = &cobra.Command{
	Use:   "crawl",
	Short: "run bgm38 spider",
	RunE: func(cmd *cobra.Command, args []string) error {
		// logrus.SetReportCaller(true)
		return spider.Start()
	},
}
