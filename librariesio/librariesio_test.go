package librariesio

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

const APIKey string = "1234"

func checkHeader(t *testing.T, req *http.Request, header string, want string) {
	if got := req.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(APIKey)

	if got, want := c.apiKey, APIKey; got != want {
		t.Errorf("NewClient baseURL is %v, want %v", got, want)
	}

	if got, want := c.BaseURL.String(), "https://libraries.io/api/"; got != want {
		t.Errorf("NewClient baseURL is %v, want %v", got, want)
	}

	if got, want := c.UserAgent, "go-librariesio/1"; got != want {
		t.Errorf("NewClient userAgent is %v, want %v", got, want)
	}

	if got, want := c.client.Timeout, time.Second*10; got != want {
		t.Errorf("NewClient timeout is %v, want %v", got, want)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	client := NewClient(APIKey)
	req, err := client.NewRequest("GET", ":", nil)

	if err == nil {
		t.Fatalf("NewRequest did not return error")
	}
	if req != nil {
		t.Fatalf("did not expect non-nil request, got %+v", req)
	}
}

func TestNewRequest_noPayload(t *testing.T) {
	client := NewClient(APIKey)
	req, err := client.NewRequest("GET", "pypi/cookiecutter", nil)

	if err != nil {
		t.Fatalf("NewRequest returned error: %v", err)
	}

	if req.Body != nil {
		t.Fatalf("request contains a non-nil Body\n%v", req.Body)
	}
}

func TestNewRequest_invalidJSON(t *testing.T) {
	client := NewClient(APIKey)

	foo := make(map[interface{}]interface{})

	_, err := client.NewRequest("GET", "pypi/cookiecutter", foo)

	if err == nil {
		t.Error("Expected error to be returned")
	}
	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error, got %#v", err)
	}
}

func TestNewRequest_auth(t *testing.T) {
	client := NewClient(APIKey)
	req, err := client.NewRequest("GET", "pypi/cookiecutter", nil)

	if err != nil {
		t.Fatalf("NewRequest returned error: %v", err)
	}

	query := req.URL.Query()
	if got, want := query.Get("api_key"), APIKey; got != want {
		t.Fatalf("did not set query param to %v, got %v", want, got)
	}
}

func TestNewRequest_headers(t *testing.T) {
	testCases := []struct {
		name    string
		body    interface{}
		headers map[string]string
	}{
		{
			name: "No Content-Type without body",
			body: nil,
			headers: map[string]string{
				"Accept":     "application/json",
				"User-Agent": "go-librariesio/" + libraryVersion,
			},
		},
		{
			name: "Content-Type with body",
			body: "hello world",
			headers: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
				"User-Agent":   "go-librariesio/" + libraryVersion,
			},
		},
	}

	client := NewClient(APIKey)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := client.NewRequest("GET", "pypi/cookiecutter", testCase.body)
			if err != nil {
				t.Fatal("unexpected error")
			}

			for key, value := range testCase.headers {
				checkHeader(t, req, key, value)
			}
		})
	}
}

func TestCheckResponse(t *testing.T) {
	response := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader(`{"error":"Nope Nope Nope"}`)),
	}
	errResponse, ok := CheckResponse(response).(*ErrorResponse)

	if !ok {
		t.Errorf("Expected ErrorResponse, got %v", errResponse)
	}

	want := &ErrorResponse{
		Response: response,
		Message:  "Nope Nope Nope",
	}
	if !reflect.DeepEqual(errResponse, want) {
		t.Errorf("\nExpected %#v\ngot %#v", want, errResponse)
	}
}
