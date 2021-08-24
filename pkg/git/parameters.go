package git

import "time"

type PullRequestsListFilters struct {
	LastActivity *time.Time
	States       *[]string
	Labels       *[]string
}

