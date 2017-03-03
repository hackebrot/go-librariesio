package librariesio

import (
	"testing"
	"time"
)

const APIKey string = "1234"

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
	client := NewClient("")
	req, err := client.NewRequest("GET", ":", nil)

	if err == nil {
		t.Fatalf("NewRequest did not return error")
	}
	if req != nil {
		t.Fatalf("did not expect non-nil request, got %+v", req)
	}
}

func TestNewRequest_noPayload(t *testing.T) {
	client := NewClient("")
	req, err := client.NewRequest("GET", "pypi/cookiecutter", nil)

	if err != nil {
		t.Fatalf("NewRequest returned error: %v", err)
	}

	if req.Body != nil {
		t.Fatalf("request contains a non-nil Body\n%v", req.Body)
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
