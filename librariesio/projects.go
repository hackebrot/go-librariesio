package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Project represents a project on libraries.io
type Project struct {
	Description              *string    `json:"description,omitempty"`
	Forks                    *int       `json:"forks,omitempty"`
	Homepage                 *string    `json:"homepage,omitempty"`
	Keywords                 []*string  `json:"keywords,omitempty"`
	Language                 *string    `json:"language,omitempty"`
	LatestReleaseNumber      *string    `json:"latest_release_number,omitempty"`
	LatestReleasePublishedAt *time.Time `json:"latest_release_published_at,omitempty"`
	LatestStableRelease      *Release   `json:"latest_stable_release,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	NormalizedLicenses       []*string  `json:"normalized_licenses,omitempty"`
	PackageManagerURL        *string    `json:"package_manager_url,omitempty"`
	Platform                 *string    `json:"platform,omitempty"`
	Rank                     *int       `json:"rank,omitempty"`
	Stars                    *int       `json:"stars,omitempty"`
	Status                   *string    `json:"status,omitempty"`
	Versions                 []*Release `json:"versions,omitempty"`
}

// Release represents a release of the project
type Release struct {
	Number      string    `json:"number,omitempty"`
	PublishedAt time.Time `json:"published_at,omitempty"`
}

// GetProject returns information about a project and it's versions.
// GET https://libraries.io/api/:platform/:name
func (c *Client) GetProject(ctx context.Context, platform string, name string) (*Project, *http.Response, error) {
	urlStr := fmt.Sprintf("%v/%v", platform, name)

	request, err := c.NewRequest("GET", urlStr, nil)

	if err != nil {
		return nil, nil, err
	}

	project := new(Project)
	response, err := c.Do(ctx, request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}
