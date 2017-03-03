package librariesio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "1"
	baseURL        = "https://libraries.io/api/"
	userAgent      = "go-librariesio/" + libraryVersion
	contentType    = "application/json"
	mediaType      = "application/json"
)

// Client for communicating with the libraries.io API
type Client struct {
	apiKey    string
	client    *http.Client
	UserAgent string
	BaseURL   *url.URL
}

// NewClient returns a new libraries.io API client
func NewClient(apiKey string) *Client {
	APIBaseURL, _ := url.Parse(baseURL)

	return &Client{
		apiKey:    apiKey,
		client:    &http.Client{Timeout: time.Second * 10},
		UserAgent: userAgent,
		BaseURL:   APIBaseURL,
	}
}

// NewRequest creates a new API request
func (c *Client) NewRequest(method, urlStr string, data interface{}) (*http.Request, error) {
	relativeURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	absoluteURL := c.BaseURL.ResolveReference(relativeURL)

	var body io.ReadWriter
	if data != nil {
		body = new(bytes.Buffer)

		err := json.NewEncoder(body).Encode(data)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, absoluteURL.String(), body)
	if err != nil {
		return nil, err
	}

	// set api_key for auth
	q := request.URL.Query()
	q.Set("api_key", c.apiKey)
	request.URL.RawQuery = q.Encode()

	request.Header.Set("Accept", mediaType)

	if data != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		request.Header.Set("User-Agent", c.UserAgent)
	}

	return request, nil
}

// ErrorResponse holds information about an unsuccesful API request
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error"`
}

// Error interface implementation for ErrorRespons
func (r *ErrorResponse) Error() string {
	// Make sure to not show api_key
	q := r.Response.Request.URL.Query()
	q.Set("api_key", "REDACTED")
	r.Response.Request.URL.RawQuery = q.Encode()

	return fmt.Sprintf(
		"%v %v: %d %q ",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Message,
	)
}

// Do sends an HTTP request and returns an HTTP response
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errorResponse := &ErrorResponse{Response: response}
		if body != nil {
			json.Unmarshal(body, errorResponse)
		}
		return response, errorResponse
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return response, nil
}
