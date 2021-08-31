package main

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	lastActivity string
	labels       []string
)

var cmdStale = &cobra.Command{
	Use:   "stale --repositories github.com/knlambert/stale-pr-detector --last-activity 5d",
	Short: "Command dedicated to find stale PRs",
	Run: func(cmd *cobra.Command, args []string) {
		prDetector := createPRDetector()

		err := prDetector.StaleList(
			repositoriesURLs,
			labels,
			lastActivity,
		)

		if err != nil {
			log.Fatalln(err)
		}
	},
}

func staleInitialize() {
	cmdStale.PersistentFlags().StringVar(
		&lastActivity,
		"last-activity",
		"30d",
		"The last activity limit, ex: 3d, 6m, 1y",
	)

	cmdStale.PersistentFlags().StringSliceVar(
		&labels,
		"labels",
		[]string{},
		"A list of labels to filter on",
	)

	rootCmd.AddCommand(cmdStale)
}
