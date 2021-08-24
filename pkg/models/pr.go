package models

import "time"

type PullRequest struct {
	Number     *string     `json:"number"`
	State      *string     `json:"state"`
	Title      *string     `json:"title"`
	Author     *string     `json:"author"`
	Labels     []string    `json:"labels"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
	Repository *Repository `json:"repository" `
}
