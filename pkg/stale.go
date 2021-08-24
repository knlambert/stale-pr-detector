package pkg

import (
	"fmt"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"time"
)


func (p *PRDetector) StaleList(
	repositories []string,
	labels []string,
	lastActivity string,
	outputFormat OutputFormat,
) error {

	convertedLastActivity, err := p.parseLastActivity(lastActivity)
	states := []string{"open"}

	if err != nil {
		return err
	}

	var allPRs []models.PullRequest

	for i := range repositories {
		if prs, err := p.gitClient.PullRequestsList(repositories[i], &git.PullRequestsListFilters{
			LastActivity: convertedLastActivity,
			Labels: &labels,
			States: &states,
		}); err == nil {
			allPRs = append(allPRs, prs...)
		} else {
			return errors.Wrapf(err, "failed to request %s", repositories[i])
		}
	}

	formatted, err := p.formatter.PrettyPrintPullRequests(allPRs)

	if err != nil {
		return err
	}

	fmt.Println(string(formatted))

	return nil
}

func (p *PRDetector) parseLastActivity(lastActivity string) (*time.Time, error) {
	activityRegex := regexp.MustCompile(`(\d+)(d|y|m)`)

	groups := activityRegex.FindStringSubmatch(lastActivity)

	if len(groups) == 0 {
		return nil, errors.New(fmt.Sprintf("wrong date format for %s", lastActivity))
	}

	count, err := strconv.Atoi(groups[1])

	if err != nil {
		return nil, err
	}

	duration := map[string]time.Duration{
		"d": time.Hour * 24,
		"m": time.Hour * 24 * 30,
		"y": time.Hour * 24 * 7 * 365,
	}[groups[2]]

	lastActivityDate := time.Now().Add(time.Duration(count) * duration)
	return &lastActivityDate, nil
}