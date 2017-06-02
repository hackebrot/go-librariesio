package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// User represents a user on libraries.io
type User struct {
	ID           *int       `json:"id,omitempty"`
	UUID         *int       `json:"uuid,omitempty"`
	Login        *string    `json:"login,omitempty"`
	UserType     *string    `json:"user_type,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Company      *string    `json:"company,omitempty"`
	Blog         *string    `json:"blog,omitempty"`
	Location     *string    `json:"location,omitempty"`
	Hidden       *bool      `json:"hidden,omitempty"`
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
	Email        *string    `json:"email,omitempty"`
	Bio          *string    `json:"bio,omitempty"`
	Followers    *int       `json:"followers,omitempty"`
	Following    *int       `json:"following,omitempty"`
	HostType     *string    `json:"host_type,omitempty"`
	GitHubID     *int       `json:"github_id,omitempty"`
}

// Repository represents a GitHub project
type Repository struct {
	ContributionsCount       *int       `json:"contributions_count,omitempty"`
	CreatedAt                *time.Time `json:"created_at,omitempty"`
	DefaultBranch            *string    `json:"default_branch,omitempty"`
	Description              *string    `json:"description,omitempty"`
	Fork                     *bool      `json:"fork,omitempty"`
	ForkPolicy               *string    `json:"fork_policy,omitempty"`
	ForksCount               *int       `json:"forks_count,omitempty"`
	FullName                 *string    `json:"full_name,omitempty"`
	GithubContributionsCount *int       `json:"github_contributions_count,omitempty"`
	GithubID                 *string    `json:"github_id,omitempty"`
	HasAudit                 *string    `json:"has_audit,omitempty"`
	HasChangelog             *string    `json:"has_changelog,omitempty"`
	HasCoc                   *string    `json:"has_coc,omitempty"`
	HasContributing          *string    `json:"has_contributing,omitempty"`
	HasIssues                *bool      `json:"has_issues,omitempty"`
	HasLicense               *string    `json:"has_license,omitempty"`
	HasPages                 *bool      `json:"has_pages,omitempty"`
	HasReadme                *string    `json:"has_readme,omitempty"`
	HasThreatModel           *string    `json:"has_threat_model,omitempty"`
	HasWiki                  *bool      `json:"has_wiki,omitempty"`
	Homepage                 *string    `json:"homepage,omitempty"`
	HostDomain               *string    `json:"host_domain,omitempty"`
	HostType                 *string    `json:"host_type,omitempty"`
	Keywords                 []*string  `json:"keywords,omitempty"`
	Language                 *string    `json:"language,omitempty"`
	LastSyncedAt             *time.Time `json:"last_synced_at,omitempty"`
	License                  *string    `json:"license,omitempty"`
	LogoURL                  *string    `json:"logo_url,omitempty"`
	MirrorURL                *string    `json:"mirror_url,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	OpenIssuesCount          *int       `json:"open_issues_count,omitempty"`
	Private                  *bool      `json:"private,omitempty"`
	PullRequestsEnabled      *bool      `json:"pull_requests_enabled,omitempty"`
	PushedAt                 *time.Time `json:"pushed_at,omitempty"`
	Rank                     *int       `json:"rank,omitempty"`
	Scm                      *string    `json:"scm,omitempty"`
	Size                     *int       `json:"size,omitempty"`
	SourceName               *string    `json:"source_name,omitempty"`
	StargazersCount          *int       `json:"stargazers_count,omitempty"`
	Status                   *string    `json:"status,omitempty"`
	SubscribersCount         *int       `json:"subscribers_count,omitempty"`
	UUID                     *string    `json:"uuid,omitempty"`
	UpdatedAt                *time.Time `json:"updated_at,omitempty"`
}

// User returns information for a given user or organization
//
// GET https://libraries.io/api/github/:login
//
// login is a user or organization on GitHub
func (c *Client) User(ctx context.Context, login string) (*User, *http.Response, error) {
	urlStr := fmt.Sprintf("github/%v", login)

	request, err := c.NewRequest("GET", urlStr, nil)

	if err != nil {
		return nil, nil, err
	}

	user := new(User)

	response, err := c.Do(ctx, request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

// UserProjects returns projects referencing the given GitHub user
//
// GET https://libraries.io/api/github/:login/projects
//
// login is a user or organization on GitHub
func (c *Client) UserProjects(ctx context.Context, login string) ([]*Project, *http.Response, error) {
	urlStr := fmt.Sprintf("github/%v/projects", login)

	request, err := c.NewRequest("GET", urlStr, nil)

	if err != nil {
		return nil, nil, err
	}

	var projects []*Project

	response, err := c.Do(ctx, request, &projects)
	if err != nil {
		return nil, response, err
	}

	return projects, response, nil
}

// UserRepositories returns repositories owned by the given GitHub user
//
// GET https://libraries.io/api/github/:login/repositories
//
// login is a user or organization on GitHub
func (c *Client) UserRepositories(ctx context.Context, login string) ([]*Repository, *http.Response, error) {
	urlStr := fmt.Sprintf("github/%v/repositories", login)

	request, err := c.NewRequest("GET", urlStr, nil)

	if err != nil {
		return nil, nil, err
	}
	var repos []*Repository

	response, err := c.Do(ctx, request, &repos)
	if err != nil {
		return nil, response, err
	}

	return repos, response, nil
}
