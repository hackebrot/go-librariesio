package librariesio

import (
	"fmt"
	"net/http"
)

// Project that holds a name field
type Project struct {
	Name string `json:"name"`
}

// GetProject returns information about a project and it's versions.
// GET https://libraries.io/api/:platform/:name
func (c *Client) GetProject(platform string, name string) (*Project, *http.Response, error) {
	urlStr := fmt.Sprintf("%v/%v", platform, name)

	request, err := c.NewRequest("GET", urlStr, nil)

	if err != nil {
		return nil, nil, err
	}

	project := new(Project)
	response, err := c.Do(request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}
