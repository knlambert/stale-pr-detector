package models

type Repository struct {
	URL   *string `json:"url"`
	Owner *string `json:"owner,omitempty"`
	Name  *string `json:"name,omitempty"`
}
