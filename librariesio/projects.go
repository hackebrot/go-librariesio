package librariesio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Project that holds a name field
type Project struct {
	Name string `json:"name,omitempty"`
}

// GetProject returns information about a project and it's versions.
// GET https://libraries.io/api/:platform/:name
func (c *Client) GetProject(platform string, name string) (*Project, error) {
	urlStr := fmt.Sprintf("%v/%v", platform, name)

	request, err := c.NewRequest("GET", urlStr, nil)

	if err != nil {
		return nil, err
	}

	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request unsuccesful for %v/%v: %v", platform, name, response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read body %v", err)
	}

	var p Project
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, fmt.Errorf("unable to deserialize project %v", err)
	}
	return &p, nil
}
