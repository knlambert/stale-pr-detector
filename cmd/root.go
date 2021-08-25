package cmd

import (
	"github.com/knlambert/stale-pr-detector/pkg"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "prq",
	Short: "PR Query is a command line designed to look for PRs in popular Git vendors",
	Args:  cobra.MinimumNArgs(1),
}

var (
	repositoriesURLs []string
	gitVendor        string
	formatType       string
	filePath         string
)

func createPRDetector() *pkg.PRDetector {
	prDetector, err := pkg.CreatePRDetector(
		pkg.GitClientVendor(gitVendor),
		pkg.OutputFormat(formatType),
		filePath,
	)

	if err != nil {
		log.Fatal(err)
	}

	return prDetector
}

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

	rootCmd.PersistentFlags().StringVar(
		&formatType,
		"format",
		"json",
		"The required output format (json|text)",
	)

	rootCmd.PersistentFlags().StringVarP(
		&filePath,
		"file",
		"f",
		"",
		"The file to output the result to",
	)

	staleInitialize()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

//
//func Execute() {
//	pr, err := pkg.CreatePRDetector(
//		git.GitClientVendorGithub,
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
