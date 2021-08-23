package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var cmdStale = &cobra.Command{
	Use:   "stale --repositories github.com/knlambert/stale-pr-detector --last-activity 5d",
	Short: "Command dedicated to find stale PRs",
	Run: func(cmd *cobra.Command, args []string) {
		err := prDetector.StaleList(
			repositoriesURLs,
			nil,
			"text",
		)

		if err != nil {
			log.Fatalln(err)
		}
	},
}
