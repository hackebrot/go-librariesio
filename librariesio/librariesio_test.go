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
