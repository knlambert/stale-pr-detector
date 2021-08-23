package pkg

import "fmt"

func (p *PRDetector) StaleList(
	repositories []string,
	lastActivity *string,
	outputFormat OutputFormat,
) error {
	fmt.Println("Yoyoyo")
	fmt.Println(repositories)
	return nil
}
