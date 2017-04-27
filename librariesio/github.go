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

// User returns information for a given user or organization
//
// GET https://libraries.io/api/github/hackebrot
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
