package main

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	noActivitySince string
	labels          []string
	authors      []string
)

var cmdStale = &cobra.Command{
	Use:   "stale --repositories github.com/knlambert/stale-pr-detector --no-activity-since 5d",
	Short: "Command dedicated to find stale PRs",
	Run: func(cmd *cobra.Command, args []string) {
		prDetector := createPRDetector()

		err := prDetector.StaleList(
			repositoriesURLs,
			labels,
			authors,
			noActivitySince,
		)

		if err != nil {
			log.Fatalln(err)
		}
	},
}

func staleInitialize() {
	cmdStale.PersistentFlags().StringVar(
		&noActivitySince,
		"no-activity-since",
		"30d",
		"The minimal amount of time without any activity, ex: 3d, 6m, 1y)",
	)

	cmdStale.PersistentFlags().StringSliceVar(
		&labels,
		"labels",
		[]string{},
		"A list of labels to filter with",
	)

	cmdStale.PersistentFlags().StringSliceVar(
		&authors,
		"authors",
		[]string{},
		"A list of authors to filter with",
	)

	rootCmd.AddCommand(cmdStale)
}
