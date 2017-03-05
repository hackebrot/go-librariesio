package librariesio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
	transport *http.Transport
	client    *http.Client
	UserAgent string
	BaseURL   *url.URL
}

// NewClient returns a new libraries.io API client
func NewClient(apiKey string) *Client {
	APIBaseURL, _ := url.Parse(baseURL)

	transport := &http.Transport{}
	client := &http.Client{Transport: transport}

	return &Client{
		apiKey:    apiKey,
		client:    client,
		transport: transport,
		UserAgent: userAgent,
		BaseURL:   APIBaseURL,
	}
}

// NewRequest creates a new API request, that can be used for client.Do().
// It creates an absolute URL from the given URL string and serialize the
// given payload, set the according headers and add the api_key query param.
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

	req, err := http.NewRequest(method, absoluteURL.String(), body)
	if err != nil {
		return nil, err
	}

	// set api_key for auth
	q := req.URL.Query()
	q.Set("api_key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// redactAPIKey overwrites the secret api_key query param
func redactAPIKey(url *url.URL) *url.URL {
	q := url.Query()
	q.Set("api_key", "REDACTED")
	url.RawQuery = q.Encode()
	return url
}

// ErrorResponse holds information about an unsuccesful API request.
// The error message from the API response is stored to the Message field.
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error"`
}

// Error returns information about the ErrorResponse
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf(
		"%v %v: %d %q",
		r.Response.Request.Method,
		redactAPIKey(r.Response.Request.URL),
		r.Response.StatusCode,
		r.Message,
	)
}

// CheckResponse checks the API response for errors and returns a ErrorResponse
// Responses are considered unsuccesful for status code other than 2xx.
func CheckResponse(resp *http.Response) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	errResp := &ErrorResponse{Response: resp}

	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errResp)
	}
	return errResp
}

// Do sends an HTTP request, that can be cancelled via the given context.
// It makes sure to redact the API secret key from any URL errors and load
// the body from the HTTP response into the given obj and return the response.
func (c *Client) Do(ctx context.Context, req *http.Request, obj interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			// Cancel the HTTP request if the given context is cancelled
			// is cancelled, when the deadline exceeds or the caller
			// cancels the context explicitly
			c.transport.CancelRequest(req)
			return nil, ctx.Err()
		default:
			// If we have encountered an url.Error make sure
			// to redact the API secret key from the URL
			if urlError, ok := err.(*url.Error); ok {
				if url, err := url.Parse(urlError.URL); err == nil {
					urlError.URL = redactAPIKey(url).String()
					return nil, urlError
				}
			}
			return nil, err
		}
	}
	defer resp.Body.Close()

	// Check that the response's status code is OK
	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	// Load body into the given obj
	if obj != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, obj)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
