package librariesio

import "time"

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
