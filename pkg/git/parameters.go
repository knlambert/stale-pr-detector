package git

import "time"

//PullRequestsListFilters can be use as a parameter with methods listing PRs from a Git Vendor.
type PullRequestsListFilters struct {
	LastActivity *time.Time
	States       *[]string
	Labels       *[]string
	Authors      *[]string
}
