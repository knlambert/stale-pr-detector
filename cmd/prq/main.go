package main

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

func execute() {
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
		"The vendor to request, with only github supported for now",
	)

	rootCmd.PersistentFlags().StringVar(
		&formatType,
		"format",
		"text",
		"The required output format (json|text)",
	)

	rootCmd.PersistentFlags().StringVarP(
		&filePath,
		"file",
		"f",
		"",
		"The file to output the result to",
	)

	_ = rootCmd.MarkPersistentFlagRequired("repositories")

	staleInitialize()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	execute()
}
