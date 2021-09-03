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

//StaleList is dedicated to returning pull requests which are without interactions for a certain amount of time.
func (p *PRDetector) StaleList(
	repositories, labels, authors []string,
	lastActivity string,
) error {
	convertedLastActivity, err := p.parseLastActivity(lastActivity)
	states := []string{"open"}

	if err != nil {
		return err
	}

	var allPRs = make([]models.PullRequest, 0)

	for i := range repositories {
		if prs, err := p.gitClient.PullRequestsList(repositories[i], &git.PullRequestsListFilters{
			LastActivity: convertedLastActivity,
			Labels:       &labels,
			States:       &states,
			Authors:      &authors,
		}); err == nil {
			allPRs = append(allPRs, prs...)
		} else {
			return errors.Wrapf(err, "failed to request %s", repositories[i])
		}
	}

	formatted, err := p.formatter.PrettyPrint(allPRs)

	if err != nil {
		return err
	}

	if _, err := p.output.Write(formatted); err != nil {
		return errors.Wrap(err, "failed to output the result")
	}

	return nil
}

//parseLastActivity converts a human friendly duration into a TimeWrapper object.
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

	now := p.time.Now()

	lastActivityDate := map[string]time.Time{
		"d": now.AddDate(0, 0, -count),
		"m": now.AddDate(0, -count, 0),
		"y": now.AddDate(-count, 0, 0),
	}[groups[2]]

	return &lastActivityDate, nil
}
