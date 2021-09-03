package models

//Repository describes a repository extracted from a git vendor.
type Repository struct {
	URL   *string `json:"url"`
	Owner *string `json:"owner,omitempty"`
	Name  *string `json:"name,omitempty"`
}
