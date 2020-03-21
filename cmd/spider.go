package cmd

import (
	"github.com/spf13/cobra"

	"bgm38/pkg/spider"
)

var spiderCmd = &cobra.Command{
	Use:   "crawl",
	Short: "run bgm38 spider",
	RunE: func(cmd *cobra.Command, args []string) error {
		// logrus.SetReportCaller(true)
		return spider.Start()
	},
}
