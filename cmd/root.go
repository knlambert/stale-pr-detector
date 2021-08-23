package cmd

import (
	"github.com/knlambert/stale-pr-detector/pkg"
	"github.com/knlambert/stale-pr-detector/pkg/crawler"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "prq",
	Short: "PR Query is a command line designed to look for PRs in popular Git vendors",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		prDetector, err = pkg.CreatePRDetector(
			git.Vendor(gitVendor),
			crawler.Parallel,
		)

		if err != nil {
			log.Fatal(err)
		}
	},
}

var (
	repositoriesURLs []string
	gitVendor string
	prDetector *pkg.PRDetector
)

func Execute() {


	rootCmd.PersistentFlags().StringSliceVar(
		&repositoriesURLs,
		"repositories",
		[]string{},
		"A list of repositories to scan",
	)

	rootCmd.PersistentFlags().StringVar(
		&gitVendor,
		"vendor",
		"github",
		"The vendor to request (only github supported for now)",
	)

	rootCmd.AddCommand(cmdStale)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

//
//func Execute() {
//	pr, err := pkg.CreatePRDetector(
//		git.VendorGithub,
//		crawler.Parallel,
//	)
//
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	err = pr.StaleList(
//		[]string{"github.com:knlambert/stale-pr-detector.git"},
//		nil,
//		"text",
//	)
//
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
