package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hackebrot/go-repr/repr"
)

func TestUser(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}

		if url := r.URL.String(); !strings.Contains(url, "/github/hackebrot") {
			t.Errorf("unexpected URL, got %v", url)
		}

		fmt.Fprintf(w, `{
			"login":"hackebrot",
			"name":"Raphael Pierzina",
			"blog":"https://raphael.codes"
		}`)
	})

	user, _, err := client.User(context.Background(), "hackebrot")
	if err != nil {
		t.Fatalf("User returned unexpected error: %v", err)
	}

	want := &User{
		Name:  String("Raphael Pierzina"),
		Login: String("hackebrot"),
		Blog:  String("https://raphael.codes"),
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(user))
	}
}

func TestUserProjects(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}

		if url := r.URL.String(); !strings.Contains(url, "/github/hackebrot/projects") {
			t.Errorf("unexpected URL, got %v", url)
		}

		fmt.Fprintf(w, `[
			{
				"name":"go-repr",
				"keywords": ["go", "golang", "helper"],
				"repository_url": "https://github.com/hackebrot/go-repr"
			},
			{
				"name":"go-librariesio",
				"keywords": ["api-client", "go", "golang"],
				"repository_url": "https://github.com/hackebrot/go-librariesio"
			}
		]`)
	})

	projects, _, err := client.UserProjects(context.Background(), "hackebrot")

	if err != nil {
		t.Fatalf("UserProjects returned unexpected error: %v", err)
	}

	want := []*Project{
		{
			Name: String("go-repr"),
			Keywords: []*string{
				String("go"),
				String("golang"),
				String("helper"),
			},
			RepositoryURL: String("https://github.com/hackebrot/go-repr"),
		},
		{
			Name: String("go-librariesio"),
			Keywords: []*string{
				String("api-client"),
				String("go"),
				String("golang"),
			},
			RepositoryURL: String("https://github.com/hackebrot/go-librariesio"),
		},
	}

	if !reflect.DeepEqual(projects, want) {
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(projects))
	}
}
