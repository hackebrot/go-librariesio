package librariesio

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "1"
	baseURL        = "https://libraries.io/api/"
	userAgent      = "go-librariesio/" + libraryVersion
	contentType    = "application/json"
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

	c := &Client{
		apiKey:    apiKey,
		client:    &http.Client{Timeout: time.Second * 10},
		UserAgent: userAgent,
		BaseURL:   APIBaseURL,
	}

	return c
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

	if data != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		request.Header.Set("User-Agent", c.UserAgent)
	}

	return request, nil
}

// Do sends an HTTP request and returns an HTTP response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
